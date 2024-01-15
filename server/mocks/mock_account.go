// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/models/Account.go
//
// Generated by this command:
//
//	mockgen -source=pkg/models/Account.go -destination=mocks/mock_account.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "server/pkg/models"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountRepository) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, account)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountRepositoryMockRecorder) CreateAccount(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountRepository)(nil).CreateAccount), ctx, account)
}

// GetAccountById mocks base method.
func (m *MockAccountRepository) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", ctx, id)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockAccountRepositoryMockRecorder) GetAccountById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountById), ctx, id)
}

// GetAccountByUserId mocks base method.
func (m *MockAccountRepository) GetAccountByUserId(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByUserId", ctx, id)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByUserId indicates an expected call of GetAccountByUserId.
func (mr *MockAccountRepositoryMockRecorder) GetAccountByUserId(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUserId", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountByUserId), ctx, id)
}

// GetAllAccounts mocks base method.
func (m *MockAccountRepository) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccounts", ctx)
	ret0, _ := ret[0].([]*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAccounts indicates an expected call of GetAllAccounts.
func (mr *MockAccountRepositoryMockRecorder) GetAllAccounts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccounts", reflect.TypeOf((*MockAccountRepository)(nil).GetAllAccounts), ctx)
}

// UpdateAccount mocks base method.
func (m *MockAccountRepository) UpdateAccount(ctx context.Context, id uuid.UUID, sum float64) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, id, sum)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountRepositoryMockRecorder) UpdateAccount(ctx, id, sum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountRepository)(nil).UpdateAccount), ctx, id, sum)
}

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountService) CreateAccount(ctx context.Context, account *models.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountServiceMockRecorder) CreateAccount(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountService)(nil).CreateAccount), ctx, account)
}

// GetAccount mocks base method.
func (m *MockAccountService) GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, id)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountServiceMockRecorder) GetAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountService)(nil).GetAccount), ctx, id)
}

// GetAccountByUserId mocks base method.
func (m *MockAccountService) GetAccountByUserId(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByUserId", ctx, id)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByUserId indicates an expected call of GetAccountByUserId.
func (mr *MockAccountServiceMockRecorder) GetAccountByUserId(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUserId", reflect.TypeOf((*MockAccountService)(nil).GetAccountByUserId), ctx, id)
}

// GetAllAccounts mocks base method.
func (m *MockAccountService) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccounts", ctx)
	ret0, _ := ret[0].([]*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAccounts indicates an expected call of GetAllAccounts.
func (mr *MockAccountServiceMockRecorder) GetAllAccounts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccounts", reflect.TypeOf((*MockAccountService)(nil).GetAllAccounts), ctx)
}

// UpdateAccount mocks base method.
func (m *MockAccountService) UpdateAccount(ctx context.Context, id uuid.UUID, sum float64, transactionType models.TransactionType) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, id, sum, transactionType)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountServiceMockRecorder) UpdateAccount(ctx, id, sum, transactionType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountService)(nil).UpdateAccount), ctx, id, sum, transactionType)
}