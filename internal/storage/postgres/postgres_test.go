//+build integration

package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"

	m "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/Decentr-net/theseus/internal/entities"
	"github.com/Decentr-net/theseus/internal/storage"
)

var (
	db  *sql.DB
	ctx = context.Background()
	s   storage.Storage
)

func TestMain(m *testing.M) {
	shutdown := setup()

	s = New(db)

	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() func() {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		Env:          map[string]string{"POSTGRES_PASSWORD": "root"},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}
	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
	})
	if err != nil {
		logrus.WithError(err).Fatalf("failed to create container")
	}

	if err := c.Start(ctx); err != nil {
		logrus.WithError(err).Fatal("failed to start container")
	}

	host, err := c.Host(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to get host")
	}

	port, err := c.MappedPort(ctx, "5432")
	if err != nil {
		logrus.WithError(err).Fatal("failed to map port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=postgres password=root sslmode=disable", host, port.Int())

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		logrus.WithError(err).Fatal("failed to open connection")
	}

	if err := db.Ping(); err != nil {
		logrus.WithError(err).Fatal("failed to ping postgres")
	}

	shutdownFn := func() {
		if c != nil {
			c.Terminate(ctx)
		}
	}

	migrate("postgres", "root", host, "postgres", port.Int())

	return shutdownFn
}

func migrate(username, password, hostname, dbname string, port int) {
	_, currFile, _, ok := runtime.Caller(0)
	if !ok {
		logrus.Fatal("failed to get current file location")
	}

	migrations := filepath.Join(currFile, "../../../../scripts/migrations/postgres/")

	migrator, err := m.New(
		fmt.Sprintf("file://%s", migrations),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			username, password, hostname, port, dbname),
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create migrator")
	}
	defer migrator.Close()

	if err := migrator.Up(); err != nil {
		logrus.WithError(err).Fatal("failed to migrate")
	}
}

func cleanup(t *testing.T) {
	_, err := db.ExecContext(ctx, `UPDATE height SET height=0`)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, `DELETE FROM "like"`)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, `DELETE FROM post`)
	require.NoError(t, err)
}

func TestPg_GetHeight(t *testing.T) {
	defer cleanup(t)

	h, err := s.GetHeight(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 0, h)
}

func TestPg_OnLockedHeight_Errors(t *testing.T) {
	defer cleanup(t)
	require.True(t, errors.Is(s.WithLockedHeight(context.Background(), 0, func(locked storage.Storage) error { return nil }), storage.ErrRequestedHeightIsTooLow))
	require.True(t, errors.Is(s.WithLockedHeight(context.Background(), 2, func(locked storage.Storage) error { return nil }), storage.ErrRequestedHeightIsTooHigh))
}

func TestPg_WithLockedHeight(t *testing.T) {
	defer cleanup(t)

	mu := sync.Mutex{}

	// Lock mutex to be sure if routine is started
	mu.Lock()
	go require.NoError(t, s.WithLockedHeight(context.Background(), 1, func(locked storage.Storage) error {
		mu.Unlock()                        // allow main routine execution
		time.Sleep(time.Millisecond * 500) // next OnLockHeight or GetHeight should wait

		h, err := locked.GetHeight(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, 0, h)

		return nil
	}))

	mu.Lock() // there we lock to prevent execution continuing

	go func() {
		mu.Lock()         // wait until second WithLockedHeight will start
		defer mu.Unlock() // allow test to finish

		h, err := s.GetHeight(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, 2, h)
	}()

	require.NoError(t, s.WithLockedHeight(context.Background(), 2, func(locked storage.Storage) error {
		mu.Unlock()                        // allow second routine to start
		time.Sleep(time.Millisecond * 500) // to be sure that second routine is started and GetHeight is called

		h, err := locked.GetHeight(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, 1, h)

		return nil
	}))

	mu.Lock() // do not finish until second routine will finish
}

func TestPg_CreatePost(t *testing.T) {
	defer cleanup(t)

	expected := entities.Post{
		UUID:         "1",
		Owner:        "2",
		Title:        "3",
		Category:     4,
		PreviewImage: "5",
		Text:         "6",
		CreatedAt:    time.Now(),
	}

	require.NoError(t, s.CreatePost(ctx, &expected))

	p, err := s.GetPost(ctx, expected.Owner, expected.UUID)
	require.NoError(t, err)
	require.Equal(t, expected.Owner, p.Owner)
	require.Equal(t, expected.UUID, p.UUID)
	require.Equal(t, expected.Title, p.Title)
	require.Equal(t, expected.Category, p.Category)
	require.Equal(t, expected.PreviewImage, p.PreviewImage)
	require.Equal(t, expected.Text, p.Text)
	require.Equal(t, expected.CreatedAt.UTC().Unix(), p.CreatedAt.Unix())
}

func TestPg_GetPost(t *testing.T) {
	defer cleanup(t)

	// GetPost tested in other tests

	_, err := s.GetPost(ctx, "1", "2")
	require.Equal(t, storage.ErrNotFound, err)
}

func TestPg_DeletePost(t *testing.T) {
	defer cleanup(t)

	p := entities.Post{
		UUID:         "1",
		Owner:        "2",
		Title:        "3",
		Category:     4,
		PreviewImage: "5",
		Text:         "6",
		CreatedAt:    time.Now(),
	}

	require.NoError(t, s.CreatePost(ctx, &p))

	require.NoError(t, s.DeletePost(ctx, p.Owner, p.UUID, p.CreatedAt, "moderator"))

	_, err := s.GetPost(ctx, p.Owner, p.UUID)
	require.Equal(t, storage.ErrNotFound, err)

	var info struct {
		DeletedAt time.Time `db:"deleted_at"`
		DeletedBy string    `db:"deleted_by"`
	}

	require.NoError(t, sqlx.Get(sqlx.NewDb(db, "postgres"), &info,
		`SELECT deleted_at, deleted_by FROM post WHERE owner=$1 AND uuid=$2`,
		p.Owner, p.UUID,
	))
	require.Equal(t, p.CreatedAt.UTC().Unix(), info.DeletedAt.Unix())
	require.Equal(t, "moderator", info.DeletedBy)
}

func TestPg_SetLike(t *testing.T) {
	defer cleanup(t)

	require.Equal(t, storage.ErrNotFound, s.SetLike(ctx, "1", "2", 1, time.Now(), "liker"))

	p := entities.Post{
		UUID:         "1",
		Owner:        "2",
		Title:        "3",
		Category:     4,
		PreviewImage: "5",
		Text:         "6",
		CreatedAt:    time.Now().UTC(),
	}

	require.NoError(t, s.CreatePost(ctx, &p))
	require.NoError(t, s.SetLike(ctx, p.Owner, p.UUID, 1, p.CreatedAt, "liker"))
}