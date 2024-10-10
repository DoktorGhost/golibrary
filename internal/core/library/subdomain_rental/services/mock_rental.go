// Code generated by MockGen. DO NOT EDIT.
// Source: rental_service.go

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	dao "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"
	gomock "github.com/golang/mock/gomock"
)

// MockRentalsRepository is a mock of RentalsRepository interface.
type MockRentalsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRentalsRepositoryMockRecorder
}

// MockRentalsRepositoryMockRecorder is the mock recorder for MockRentalsRepository.
type MockRentalsRepositoryMockRecorder struct {
	mock *MockRentalsRepository
}

// NewMockRentalsRepository creates a new mock instance.
func NewMockRentalsRepository(ctrl *gomock.Controller) *MockRentalsRepository {
	mock := &MockRentalsRepository{ctrl: ctrl}
	mock.recorder = &MockRentalsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRentalsRepository) EXPECT() *MockRentalsRepositoryMockRecorder {
	return m.recorder
}

// CreateRentals mocks base method.
func (m *MockRentalsRepository) CreateRentals(bookID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRentals", bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRentals indicates an expected call of CreateRentals.
func (mr *MockRentalsRepositoryMockRecorder) CreateRentals(bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRentals", reflect.TypeOf((*MockRentalsRepository)(nil).CreateRentals), bookID)
}

// CreateRentalsInfo mocks base method.
func (m *MockRentalsRepository) CreateRentalsInfo(userID, bookID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRentalsInfo", userID, bookID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRentalsInfo indicates an expected call of CreateRentalsInfo.
func (mr *MockRentalsRepositoryMockRecorder) CreateRentalsInfo(userID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRentalsInfo", reflect.TypeOf((*MockRentalsRepository)(nil).CreateRentalsInfo), userID, bookID)
}

// DeleteRentals mocks base method.
func (m *MockRentalsRepository) DeleteRentals(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRentals", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRentals indicates an expected call of DeleteRentals.
func (mr *MockRentalsRepositoryMockRecorder) DeleteRentals(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRentals", reflect.TypeOf((*MockRentalsRepository)(nil).DeleteRentals), id)
}

// DeleteRentalsInfo mocks base method.
func (m *MockRentalsRepository) DeleteRentalsInfo(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRentalsInfo", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRentalsInfo indicates an expected call of DeleteRentalsInfo.
func (mr *MockRentalsRepositoryMockRecorder) DeleteRentalsInfo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRentalsInfo", reflect.TypeOf((*MockRentalsRepository)(nil).DeleteRentalsInfo), id)
}

// GetActiveRentals mocks base method.
func (m *MockRentalsRepository) GetActiveRentals() (map[int][]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveRentals")
	ret0, _ := ret[0].(map[int][]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveRentals indicates an expected call of GetActiveRentals.
func (mr *MockRentalsRepositoryMockRecorder) GetActiveRentals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveRentals", reflect.TypeOf((*MockRentalsRepository)(nil).GetActiveRentals))
}

// GetRentalsByID mocks base method.
func (m *MockRentalsRepository) GetRentalsByID(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRentalsByID", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRentalsByID indicates an expected call of GetRentalsByID.
func (mr *MockRentalsRepositoryMockRecorder) GetRentalsByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRentalsByID", reflect.TypeOf((*MockRentalsRepository)(nil).GetRentalsByID), id)
}

// GetRentalsInfoByID mocks base method.
func (m *MockRentalsRepository) GetRentalsInfoByID(id int) (dao.RentalsTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRentalsInfoByID", id)
	ret0, _ := ret[0].(dao.RentalsTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRentalsInfoByID indicates an expected call of GetRentalsInfoByID.
func (mr *MockRentalsRepositoryMockRecorder) GetRentalsInfoByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRentalsInfoByID", reflect.TypeOf((*MockRentalsRepository)(nil).GetRentalsInfoByID), id)
}

// GetTopAuthors mocks base method.
func (m *MockRentalsRepository) GetTopAuthors(days, limit int) ([]dao.TopAuthor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopAuthors", days, limit)
	ret0, _ := ret[0].([]dao.TopAuthor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopAuthors indicates an expected call of GetTopAuthors.
func (mr *MockRentalsRepositoryMockRecorder) GetTopAuthors(days, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopAuthors", reflect.TypeOf((*MockRentalsRepository)(nil).GetTopAuthors), days, limit)
}

// UpdateRentals mocks base method.
func (m *MockRentalsRepository) UpdateRentals(id, rentals_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRentals", id, rentals_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRentals indicates an expected call of UpdateRentals.
func (mr *MockRentalsRepositoryMockRecorder) UpdateRentals(id, rentals_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRentals", reflect.TypeOf((*MockRentalsRepository)(nil).UpdateRentals), id, rentals_id)
}

// UpdateRentalsInfo mocks base method.
func (m *MockRentalsRepository) UpdateRentalsInfo(rentals dao.RentalsTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRentalsInfo", rentals)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRentalsInfo indicates an expected call of UpdateRentalsInfo.
func (mr *MockRentalsRepositoryMockRecorder) UpdateRentalsInfo(rentals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRentalsInfo", reflect.TypeOf((*MockRentalsRepository)(nil).UpdateRentalsInfo), rentals)
}