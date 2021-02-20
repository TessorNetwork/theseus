// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	types "github.com/Decentr-net/decentr/x/community/types"
	entities "github.com/Decentr-net/theseus/internal/entities"
	service "github.com/Decentr-net/theseus/internal/service"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// OnHeight mocks base method
func (m *MockService) OnHeight(ctx context.Context, height uint64, f func(service.Service) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnHeight", ctx, height, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnHeight indicates an expected call of OnHeight
func (mr *MockServiceMockRecorder) OnHeight(ctx, height, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnHeight", reflect.TypeOf((*MockService)(nil).OnHeight), ctx, height, f)
}

// GetHeight mocks base method
func (m *MockService) GetHeight(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeight", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeight indicates an expected call of GetHeight
func (mr *MockServiceMockRecorder) GetHeight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeight", reflect.TypeOf((*MockService)(nil).GetHeight), ctx)
}

// GetProfiles mocks base method
func (m *MockService) GetProfiles(ctx context.Context, addr []string) ([]entities.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfiles", ctx, addr)
	ret0, _ := ret[0].([]entities.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfiles indicates an expected call of GetProfiles
func (mr *MockServiceMockRecorder) GetProfiles(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfiles", reflect.TypeOf((*MockService)(nil).GetProfiles), ctx, addr)
}

// SetProfile mocks base method
func (m *MockService) SetProfile(ctx context.Context, p *entities.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProfile", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProfile indicates an expected call of SetProfile
func (mr *MockServiceMockRecorder) SetProfile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProfile", reflect.TypeOf((*MockService)(nil).SetProfile), ctx, p)
}

// Follow mocks base method
func (m *MockService) Follow(ctx context.Context, follower, followee string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", ctx, follower, followee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Follow indicates an expected call of Follow
func (mr *MockServiceMockRecorder) Follow(ctx, follower, followee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockService)(nil).Follow), ctx, follower, followee)
}

// Unfollow mocks base method
func (m *MockService) Unfollow(ctx context.Context, follower, followee string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", ctx, follower, followee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unfollow indicates an expected call of Unfollow
func (mr *MockServiceMockRecorder) Unfollow(ctx, follower, followee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockService)(nil).Unfollow), ctx, follower, followee)
}

// ListPosts mocks base method
func (m *MockService) ListPosts(ctx context.Context, filter service.ListPostsParams) ([]entities.CalculatedPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", ctx, filter)
	ret0, _ := ret[0].([]entities.CalculatedPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts
func (mr *MockServiceMockRecorder) ListPosts(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockService)(nil).ListPosts), ctx, filter)
}

// GetPost mocks base method
func (m *MockService) GetPost(ctx context.Context, id service.PostID) (entities.CalculatedPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", ctx, id)
	ret0, _ := ret[0].(entities.CalculatedPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost
func (mr *MockServiceMockRecorder) GetPost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockService)(nil).GetPost), ctx, id)
}

// CreatePost mocks base method
func (m *MockService) CreatePost(ctx context.Context, p *entities.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockServiceMockRecorder) CreatePost(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockService)(nil).CreatePost), ctx, p)
}

// DeletePost mocks base method
func (m *MockService) DeletePost(ctx context.Context, postOwner, postUUID string, timestamp time.Time, deletedByUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, postOwner, postUUID, timestamp, deletedByUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost
func (mr *MockServiceMockRecorder) DeletePost(ctx, postOwner, postUUID, timestamp, deletedByUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockService)(nil).DeletePost), ctx, postOwner, postUUID, timestamp, deletedByUUID)
}

// SetLike mocks base method
func (m *MockService) SetLike(ctx context.Context, postOwner, postUUID string, weight types.LikeWeight, timestamp time.Time, likedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLike", ctx, postOwner, postUUID, weight, timestamp, likedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLike indicates an expected call of SetLike
func (mr *MockServiceMockRecorder) SetLike(ctx, postOwner, postUUID, weight, timestamp, likedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLike", reflect.TypeOf((*MockService)(nil).SetLike), ctx, postOwner, postUUID, weight, timestamp, likedBy)
}

// GetStats mocks base method
func (m *MockService) GetStats(ctx context.Context, id []service.PostID) (map[service.PostID]entities.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", ctx, id)
	ret0, _ := ret[0].(map[service.PostID]entities.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats
func (mr *MockServiceMockRecorder) GetStats(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockService)(nil).GetStats), ctx, id)
}
