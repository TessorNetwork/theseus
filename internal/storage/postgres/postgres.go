// Package postgres is implementation of storage interface.
package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"

	community "github.com/TessorNetwork/furya/x/community/types"

	"github.com/TessorNetwork/theseus/internal/storage"
)

var log = logrus.WithField("layer", "storage").WithField("package", "postgres")
var errBeginCalledWithinTx = errors.New("can not run WithLockedHeight in tx")

const foreignKeyViolation = "23503"

type pg struct {
	ext sqlx.ExtContext
}

type postDTO struct {
	UUID         string    `db:"uuid"`
	Owner        string    `db:"owner"`
	Title        string    `db:"title"`
	Category     uint8     `db:"category"`
	PreviewImage string    `db:"preview_image"`
	Text         string    `db:"text"`
	CreatedAt    time.Time `db:"created_at"`
	Likes        uint32    `db:"likes"`
	Dislikes     uint32    `db:"dislikes"`
	UPDV         int64     `db:"updv"`
	Slug         string    `db:"slug"`
}

func (p *postDTO) toStorage() *storage.Post {
	o := storage.Post{
		UUID:         p.UUID,
		Owner:        p.Owner,
		Title:        p.Title,
		Category:     community.Category(p.Category),
		PreviewImage: p.PreviewImage,
		Text:         p.Text,
		Likes:        p.Likes,
		Dislikes:     p.Dislikes,
		UPDV:         p.UPDV,
		CreatedAt:    p.CreatedAt,
		Slug:         p.Slug,
	}

	// return post consistent with blockchain
	if strings.HasPrefix(o.Owner, "deleted") {
		o.Owner = ""
	}

	return &o
}

func (s pg) InTx(ctx context.Context, f func(s storage.Storage) error) (err error) {
	db, ok := s.ext.(*sqlx.DB)
	if !ok {
		return errBeginCalledWithinTx
	}

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to create tx: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.WithError(err).Error("failed to rollback tx")
			}
		}
	}()

	if err := f(pg{ext: tx}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}

	return nil
}

func (s pg) GetHeight(ctx context.Context) (uint64, error) {
	var h uint64
	if err := sqlx.GetContext(ctx, s.ext, &h, `SELECT height FROM height`); err != nil {
		return 0, fmt.Errorf("failed to query: %w", err)
	}

	return h, nil
}

func (s pg) SetHeight(ctx context.Context, height uint64) error {
	if _, err := s.ext.ExecContext(ctx, `UPDATE height SET height = $1`, height); err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	return nil
}

func (s pg) RefreshViews(ctx context.Context, postView, statsView bool) error {
	if postView {
		if _, err := s.ext.ExecContext(ctx, `REFRESH MATERIALIZED VIEW calculated_post`); err != nil {
			return fmt.Errorf("failed to refresh calculated_post view: failed to exec: %w", err)
		}

		if _, err := s.ext.ExecContext(ctx, `REFRESH MATERIALIZED VIEW stats`); err != nil {
			return fmt.Errorf("failed to refresh stats view: failed to exec: %w", err)
		}
	}

	if statsView {
		if _, err := s.ext.ExecContext(ctx, `REFRESH MATERIALIZED VIEW pdv_stats`); err != nil {
			return fmt.Errorf("failed to refresh pdv_stats view: failed to exec: %w", err)
		}
	}

	return nil
}

func (s pg) GetProfileStats(ctx context.Context, addr ...string) ([]*storage.ProfileStats, error) {
	if len(addr) == 0 {
		return []*storage.ProfileStats{}, nil
	}

	addr = stringsUnique(addr)

	query, args, err := sqlx.In(`
			WITH
			pc AS (
			    SELECT owner AS address, COUNT(*) as posts_count
			    FROM post
			    WHERE deleted_at IS NULL
			    GROUP BY address
			),
			r AS (
			    SELECT UNNEST(ARRAY[?]::TEXT[]) AS address
			)
			SELECT
			    r.address, COALESCE(posts_count, 0) AS posts_count, stats
			FROM r
		         LEFT JOIN pc USING (address)
		         LEFT JOIN pdv_stats USING (address)
			ORDER BY address
		`, addr)

	if err != nil {
		return nil, fmt.Errorf("failed to construct IN clause: %w", err)
	}

	var p []*struct {
		Address    string `db:"address"`
		PostsCount uint16 `db:"posts_count"`
		Stats      []byte `db:"stats"`
	}

	if err := sqlx.SelectContext(ctx, s.ext, &p, s.ext.Rebind(query), args...); err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	out := make([]*storage.ProfileStats, len(p))
	for i, v := range p {
		out[i] = &storage.ProfileStats{
			Address:    v.Address,
			PostsCount: v.PostsCount,
			Stats:      storage.PostStats{},
		}

		if v.Stats != nil {
			if err := json.Unmarshal(v.Stats, &out[i].Stats); err != nil {
				return nil, fmt.Errorf("failed to unmarshal stats: %w", err)
			}
		}
	}

	return out, nil
}

func (s pg) CreatePost(ctx context.Context, p *storage.CreatePostParams) error {
	post := postDTO{
		UUID:         p.UUID,
		Owner:        p.Owner,
		Title:        p.Title,
		Category:     uint8(p.Category),
		PreviewImage: p.PreviewImage,
		Text:         p.Text,
		CreatedAt:    p.CreatedAt.UTC(),
	}

	if _, err := sqlx.NamedExecContext(ctx, s.ext,
		`
			INSERT INTO post(owner, uuid, title, category, preview_image, text, created_at)
			VALUES(:owner, :uuid, :title, :category, :preview_image, :text, :created_at)
		`, post,
	); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (s pg) GetPost(ctx context.Context, id storage.PostID) (*storage.Post, error) {
	var p postDTO

	if err := sqlx.GetContext(ctx, s.ext, &p, `
			SELECT owner, uuid, title, category, preview_image, text, created_at, likes, dislikes, updv, slug
			FROM calculated_post
			WHERE owner = $1 AND uuid = $2
		`,
		id.Owner, id.UUID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed to query: %w", err)
	}

	return p.toStorage(), nil
}

func (s pg) GetPostBySlug(ctx context.Context, slug string) (*storage.Post, error) {
	var p postDTO

	if err := sqlx.GetContext(ctx, s.ext, &p, `
			SELECT owner, uuid, title, category, preview_image, text, created_at, likes, dislikes, updv, slug
			FROM calculated_post
			WHERE slug = $1
		`,
		slug,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed to query: %w", err)
	}

	return p.toStorage(), nil
}

func (s pg) DeletePost(ctx context.Context, id storage.PostID, timestamp time.Time, deletedBy string) error {
	res, err := s.ext.ExecContext(ctx,
		`UPDATE post SET deleted_at=$3, deleted_by=$4 WHERE owner=$1 AND uuid=$2 AND deleted_at IS NULL`,
		id.Owner, id.UUID, timestamp.UTC(), deletedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	if c, _ := res.RowsAffected(); c == 0 {
		return storage.ErrNotFound
	}

	return nil
}

func (s pg) GetLikes(ctx context.Context, likedBy string, id ...storage.PostID) (map[storage.PostID]community.LikeWeight, error) {
	if len(id) == 0 {
		return map[storage.PostID]community.LikeWeight{}, nil
	}

	owners, uuids := make([]string, len(id)), make([]string, len(id))
	for i := range id {
		owners[i] = id[i].Owner
		uuids[i] = id[i].UUID
	}

	type likesDTO struct {
		Owner  string `db:"post_owner"`
		UUID   string `db:"post_uuid"`
		Weight int8   `db:"weight"`
	}

	var res []*likesDTO

	if err := sqlx.SelectContext(ctx, s.ext, &res, `
		WITH clause AS ( SELECT UNNEST($1::TEXT[]) AS post_owner, UNNEST($2::TEXT[]) AS post_uuid )
		SELECT post_owner, post_uuid, weight FROM clause
			LEFT JOIN "like" USING(post_owner, post_uuid)
		WHERE liked_by = $3
	`, pq.StringArray(owners), pq.StringArray(uuids), likedBy); err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	out := make(map[storage.PostID]community.LikeWeight, len(res))

	for _, v := range res {
		out[storage.PostID{Owner: v.Owner, UUID: v.UUID}] = community.LikeWeight(v.Weight)
	}

	return out, nil
}

func (s pg) SetLike(ctx context.Context, id storage.PostID, weight community.LikeWeight,
	timestamp time.Time, likeOwner string) error {
	if _, err := s.ext.ExecContext(ctx, `
			INSERT INTO "like"(post_owner, post_uuid, liked_by, weight, liked_at)
				VALUES($1, $2, $3, $4, $5)
			ON CONFLICT(post_owner, post_uuid, liked_by) DO UPDATE SET
				weight=excluded.weight, liked_at=excluded.liked_at`,
		id.Owner, id.UUID, likeOwner, weight, timestamp.UTC(),
	); err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code == foreignKeyViolation {
			return storage.ErrNotFound
		}

		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (s pg) Follow(ctx context.Context, follower, followee string) error {
	if _, err := s.ext.ExecContext(ctx,
		`
			INSERT INTO follow(follower, followee) VALUES($1, $2) ON CONFLICT DO NOTHING
		`, follower, followee,
	); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (s pg) Unfollow(ctx context.Context, follower, followee string) error {
	if _, err := s.ext.ExecContext(ctx,
		`
			DELETE FROM follow WHERE follower=$1 AND followee=$2
		`, follower, followee,
	); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (s pg) ListPosts(ctx context.Context, p *storage.ListPostsParams) ([]*storage.Post, error) {
	var b strings.Builder
	var args []interface{}

	b.WriteString(`
		SELECT
			owner, uuid, title, category, preview_image, text, created_at, likes, dislikes, updv, slug
		FROM calculated_post
	`)

	if p.FollowedBy != nil {
		b.WriteString(`
			INNER JOIN follow ON calculated_post.owner = follow.followee AND follow.follower = ?
		`)
		args = append(args, *p.FollowedBy)
	}

	if p.LikedBy != nil {
		b.WriteString(`
			INNER JOIN "like"
			ON
				calculated_post.owner = "like".post_owner AND calculated_post.uuid = "like".post_uuid AND
				"like".liked_by = ? AND "like".weight != 0
		`)
		args = append(args, *p.LikedBy)
	}

	if wheres, whereArgs := whereClausesFromListPostsParams(p); len(wheres) > 0 {
		toJoin := make([]string, len(wheres))
		for i, v := range wheres {
			toJoin[i] = fmt.Sprintf("(%s)", v)
		}
		b.WriteString(` WHERE ` + strings.Join(toJoin, " AND ")) // nolint: gosec
		args = append(args, whereArgs...)
	}

	b.WriteString(fmt.Sprintf(`
		ORDER BY %s %s, owner %s, uuid %s LIMIT ?
	`, p.SortBy, p.OrderBy, p.OrderBy, p.OrderBy))
	args = append(args, p.Limit)

	query := s.ext.Rebind(b.String())

	var res []*postDTO
	if err := sqlx.SelectContext(ctx, s.ext, &res, query, args...); err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	out := make([]*storage.Post, len(res))
	for i, v := range res {
		out[i] = v.toStorage()
	}

	return out, nil
}

func (s pg) GetPostStats(ctx context.Context, id ...storage.PostID) (map[storage.PostID]storage.PostStats, error) {
	if len(id) == 0 {
		return map[storage.PostID]storage.PostStats{}, nil
	}

	owners, uuids := make([]string, len(id)), make([]string, len(id))
	for i := range id {
		owners[i] = id[i].Owner
		uuids[i] = id[i].UUID
	}

	type statsDTO struct {
		Owner string          `db:"owner"`
		UUID  string          `db:"uuid"`
		Stats json.RawMessage `db:"stats"`
	}

	var res []*statsDTO

	if err := sqlx.SelectContext(ctx, s.ext, &res, `
		WITH clause AS ( SELECT UNNEST($1::TEXT[]) AS owner, UNNEST($2::TEXT[]) AS uuid )
		SELECT owner, uuid, stats FROM stats
			INNER JOIN clause USING(owner, uuid)
	`, pq.StringArray(owners), pq.StringArray(uuids)); err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	out := make(map[storage.PostID]storage.PostStats, len(res))

	for _, v := range res {
		var s storage.PostStats
		if err := json.Unmarshal(v.Stats, &s); err != nil {
			return nil, fmt.Errorf("failed to unmarshal stats: %w", err)
		}

		out[storage.PostID{Owner: v.Owner, UUID: v.UUID}] = s
	}

	return out, nil
}

func (s pg) AddPDV(ctx context.Context, address string, updv int64, timestamp time.Time) error {
	_, err := s.ext.ExecContext(ctx, `
		INSERT INTO updv(address, updv, timestamp) VALUES($1, $2, $3)
	`, address, updv, timestamp)

	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}

func (s pg) GetFuryaStats(ctx context.Context) (*storage.FuryaStats, error) {
	var statsDTO struct {
		DDV int64   `db:"ddv"`
		ADV float64 `db:"adv"`
	}

	if err := sqlx.GetContext(ctx, s.ext, &statsDTO, `
		WITH pdv AS (
			SELECT address, SUM(updv) AS earned_pdv, SUM(updv) + $1 AS current_pdv from updv
			group by address
		)
		SELECT SUM(earned_pdv) AS ddv, AVG(current_pdv) AS adv FROM pdv;
	`, storage.PDVDenominator); err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	return &storage.FuryaStats{
		ADV: statsDTO.ADV,
		DDV: statsDTO.DDV,
	}, nil
}

func (s pg) GetDDVStats(ctx context.Context) ([]*storage.DDVStatsItem, error) {
	var stats []*storage.DDVStatsItem
	err := sqlx.SelectContext(ctx, s.ext, &stats, `
				SELECT timestamp::DATE as date, SUM(updv) as value
				FROM updv
				WHERE timestamp > NOW() -'90 day'::INTERVAL
				GROUP BY date
				ORDER BY date DESC
	`)
	return stats, err
}

func (s pg) ResetAccount(ctx context.Context, owner string) error {
	if _, ok := s.ext.(*sqlx.Tx); !ok {
		return errors.New("ResetAccount can be run only in tx mode") //nolint:goerr113
	}

	if _, err := s.ext.ExecContext(ctx, `
		DELETE FROM follow WHERE followee = $1 OR follower = $1
	`, owner); err != nil {
		return fmt.Errorf("failed to delete follows: %w", err)
	}

	if _, err := s.ext.ExecContext(ctx, `
		DELETE FROM "like" WHERE liked_by = $1 OR post_owner = $1
	`, owner); err != nil {
		return fmt.Errorf("failed to delete likes: %w", err)
	}

	if _, err := s.ext.ExecContext(ctx, `
		DELETE FROM post WHERE owner = $1
	`, owner); err != nil {
		return fmt.Errorf("failed to delete user from posts: %w", err)
	}

	if _, err := s.ext.ExecContext(ctx, `
		DELETE FROM updv WHERE address = $1
	`, owner); err != nil {
		return fmt.Errorf("failed to delete updv: %w", err)
	}

	return s.RefreshViews(ctx, true, true)
}

// New creates new instance of pg.
func New(db *sql.DB) storage.Storage {
	return pg{
		ext: sqlx.NewDb(db, "postgres"),
	}
}

func stringsUnique(s []string) []string {
	m := make(map[string]struct{}, len(s))
	out := make([]string, 0, len(s))

	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			out = append(out, v)
		}
	}

	return out
}

func whereClausesFromListPostsParams(p *storage.ListPostsParams) ([]string, []interface{}) {
	var (
		where []string
		args  []interface{}
	)

	if p.Category != nil {
		where = append(where, `category = ?`)
		args = append(args, *p.Category)
	}

	if p.Owner != nil {
		where = append(where, `owner = ?`)
		args = append(args, *p.Owner)
	}

	if p.From != nil {
		where = append(where, `created_at > ?`)
		args = append(args, time.Unix(int64(*p.From), 0).UTC())
	}

	if p.To != nil {
		where = append(where, `created_at < ?`)
		args = append(args, time.Unix(int64(*p.To), 0).UTC())
	}

	if p.ExcludeNegative {
		where = append(where, `updv >= 0`)
	}

	if p.ExcludeNeutral {
		where = append(where, `updv <> 0`)
	}

	if p.After != nil {
		comp := "<"
		if p.OrderBy == storage.AscendingOrder {
			comp = ">"
		}

		// nolint: gosec
		where = append(where, fmt.Sprintf(`
			%s %s (SELECT %s FROM calculated_post WHERE owner = ? AND uuid = ? FETCH FIRST ROW ONLY) OR (
				%s = (SELECT %s FROM calculated_post WHERE owner = ? AND uuid = ? FETCH FIRST ROW ONLY) AND
				CONCAT(owner,uuid) %s CONCAT(?::TEXT,?::TEXT)
			)
		`, p.SortBy, comp, p.SortBy, p.SortBy, p.SortBy, comp))

		args = append(args, p.After.Owner, p.After.UUID, p.After.Owner, p.After.UUID, p.After.Owner, p.After.UUID)
	}

	return where, args
}
