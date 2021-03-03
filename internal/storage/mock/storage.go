// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	types "github.com/Decentr-net/decentr/x/community/types"
	storage "github.com/Decentr-net/theseus/internal/storage"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// WithLockedHeight mocks base method
func (m *MockStorage) WithLockedHeight(ctx context.Context, height uint64, f func(storage.Storage) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithLockedHeight", ctx, height, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithLockedHeight indicates an expected call of WithLockedHeight
func (mr *MockStorageMockRecorder) WithLockedHeight(ctx, height, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLockedHeight", reflect.TypeOf((*MockStorage)(nil).WithLockedHeight), ctx, height, f)
}

// GetHeight mocks base method
func (m *MockStorage) GetHeight(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeight", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeight indicates an expected call of GetHeight
func (mr *MockStorageMockRecorder) GetHeight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeight", reflect.TypeOf((*MockStorage)(nil).GetHeight), ctx)
}

// GetProfiles mocks base method
func (m *MockStorage) GetProfiles(ctx context.Context, addr ...string) ([]*storage.Profile, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range addr {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProfiles", varargs...)
	ret0, _ := ret[0].([]*storage.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfiles indicates an expected call of GetProfiles
func (mr *MockStorageMockRecorder) GetProfiles(ctx interface{}, addr ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, addr...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfiles", reflect.TypeOf((*MockStorage)(nil).GetProfiles), varargs...)
}

// SetProfile mocks base method
func (m *MockStorage) SetProfile(ctx context.Context, p *storage.SetProfileParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProfile", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProfile indicates an expected call of SetProfile
func (mr *MockStorageMockRecorder) SetProfile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProfile", reflect.TypeOf((*MockStorage)(nil).SetProfile), ctx, p)
}

// Follow mocks base method
func (m *MockStorage) Follow(ctx context.Context, follower, followee string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", ctx, follower, followee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Follow indicates an expected call of Follow
func (mr *MockStorageMockRecorder) Follow(ctx, follower, followee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockStorage)(nil).Follow), ctx, follower, followee)
}

// Unfollow mocks base method
func (m *MockStorage) Unfollow(ctx context.Context, follower, followee string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", ctx, follower, followee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unfollow indicates an expected call of Unfollow
func (mr *MockStorageMockRecorder) Unfollow(ctx, follower, followee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockStorage)(nil).Unfollow), ctx, follower, followee)
}

// ListPosts mocks base method
func (m *MockStorage) ListPosts(ctx context.Context, p *storage.ListPostsParams) ([]*storage.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", ctx, p)
	ret0, _ := ret[0].([]*storage.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts
func (mr *MockStorageMockRecorder) ListPosts(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockStorage)(nil).ListPosts), ctx, p)
}

// CreatePost mocks base method
func (m *MockStorage) CreatePost(ctx context.Context, p *storage.CreatePostParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockStorageMockRecorder) CreatePost(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStorage)(nil).CreatePost), ctx, p)
}

// GetPost mocks base method
func (m *MockStorage) GetPost(ctx context.Context, id storage.PostID) (*storage.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", ctx, id)
	ret0, _ := ret[0].(*storage.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost
func (mr *MockStorageMockRecorder) GetPost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockStorage)(nil).GetPost), ctx, id)
}

// DeletePost mocks base method
func (m *MockStorage) DeletePost(ctx context.Context, id storage.PostID, timestamp time.Time, deletedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, id, timestamp, deletedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost
func (mr *MockStorageMockRecorder) DeletePost(ctx, id, timestamp, deletedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockStorage)(nil).DeletePost), ctx, id, timestamp, deletedBy)
}

// GetLikes mocks base method
func (m *MockStorage) GetLikes(ctx context.Context, likedBy string, id ...storage.PostID) (map[storage.PostID]types.LikeWeight, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, likedBy}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLikes", varargs...)
	ret0, _ := ret[0].(map[storage.PostID]types.LikeWeight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikes indicates an expected call of GetLikes
func (mr *MockStorageMockRecorder) GetLikes(ctx, likedBy interface{}, id ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, likedBy}, id...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikes", reflect.TypeOf((*MockStorage)(nil).GetLikes), varargs...)
}

// SetLike mocks base method
func (m *MockStorage) SetLike(ctx context.Context, id storage.PostID, weight types.LikeWeight, timestamp time.Time, likeOwner string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLike", ctx, id, weight, timestamp, likeOwner)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLike indicates an expected call of SetLike
func (mr *MockStorageMockRecorder) SetLike(ctx, id, weight, timestamp, likeOwner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLike", reflect.TypeOf((*MockStorage)(nil).SetLike), ctx, id, weight, timestamp, likeOwner)
}

// GetStats mocks base method
func (m *MockStorage) GetStats(ctx context.Context, id ...storage.PostID) (map[storage.PostID]storage.Stats, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStats", varargs...)
	ret0, _ := ret[0].(map[storage.PostID]storage.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats
func (mr *MockStorageMockRecorder) GetStats(ctx interface{}, id ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, id...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockStorage)(nil).GetStats), varargs...)
}
