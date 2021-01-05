// Code generated by MockGen. DO NOT EDIT.
// Source: crm.go

// Package itf is a generated GoMock package.
package itf

import (
	entity "github.com/Linus-Boehm/go-serverless-suite/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNewsWriter is a mock of NewsWriter interface
type MockNewsWriter struct {
	ctrl     *gomock.Controller
	recorder *MockNewsWriterMockRecorder
}

// MockNewsWriterMockRecorder is the mock recorder for MockNewsWriter
type MockNewsWriterMockRecorder struct {
	mock *MockNewsWriter
}

// NewMockNewsWriter creates a new mock instance
func NewMockNewsWriter(ctrl *gomock.Controller) *MockNewsWriter {
	mock := &MockNewsWriter{ctrl: ctrl}
	mock.recorder = &MockNewsWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNewsWriter) EXPECT() *MockNewsWriterMockRecorder {
	return m.recorder
}

// MockCRMServicer is a mock of CRMServicer interface
type MockCRMServicer struct {
	ctrl     *gomock.Controller
	recorder *MockCRMServicerMockRecorder
}

// MockCRMServicerMockRecorder is the mock recorder for MockCRMServicer
type MockCRMServicerMockRecorder struct {
	mock *MockCRMServicer
}

// NewMockCRMServicer creates a new mock instance
func NewMockCRMServicer(ctrl *gomock.Controller) *MockCRMServicer {
	mock := &MockCRMServicer{ctrl: ctrl}
	mock.recorder = &MockCRMServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCRMServicer) EXPECT() *MockCRMServicerMockRecorder {
	return m.recorder
}

// GetMailer mocks base method
func (m *MockCRMServicer) GetMailer() Mailer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMailer")
	ret0, _ := ret[0].(Mailer)
	return ret0
}

// GetMailer indicates an expected call of GetMailer
func (mr *MockCRMServicerMockRecorder) GetMailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMailer", reflect.TypeOf((*MockCRMServicer)(nil).GetMailer))
}

// CreateSubscription mocks base method
func (m *MockCRMServicer) CreateSubscription(subs []entity.CRMEmailListSubscription, confirmationTPL entity.HTMLTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", subs, confirmationTPL)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscription indicates an expected call of CreateSubscription
func (mr *MockCRMServicerMockRecorder) CreateSubscription(subs, confirmationTPL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockCRMServicer)(nil).CreateSubscription), subs, confirmationTPL)
}

// ValidateEmail mocks base method
func (m *MockCRMServicer) ValidateEmail(email entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateEmail", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateEmail indicates an expected call of ValidateEmail
func (mr *MockCRMServicerMockRecorder) ValidateEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateEmail", reflect.TypeOf((*MockCRMServicer)(nil).ValidateEmail), email)
}

// MockCRMProvider is a mock of CRMProvider interface
type MockCRMProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCRMProviderMockRecorder
}

// MockCRMProviderMockRecorder is the mock recorder for MockCRMProvider
type MockCRMProviderMockRecorder struct {
	mock *MockCRMProvider
}

// NewMockCRMProvider creates a new mock instance
func NewMockCRMProvider(ctrl *gomock.Controller) *MockCRMProvider {
	mock := &MockCRMProvider{ctrl: ctrl}
	mock.recorder = &MockCRMProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCRMProvider) EXPECT() *MockCRMProviderMockRecorder {
	return m.recorder
}

// GetSubscriptionsOfEmail mocks base method
func (m *MockCRMProvider) GetSubscriptionsOfEmail(email entity.ID) ([]entity.CRMEmailListSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionsOfEmail", email)
	ret0, _ := ret[0].([]entity.CRMEmailListSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionsOfEmail indicates an expected call of GetSubscriptionsOfEmail
func (mr *MockCRMProviderMockRecorder) GetSubscriptionsOfEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionsOfEmail", reflect.TypeOf((*MockCRMProvider)(nil).GetSubscriptionsOfEmail), email)
}

// PutSubscription mocks base method
func (m *MockCRMProvider) PutSubscription(arg0 entity.CRMEmailListSubscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSubscription", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSubscription indicates an expected call of PutSubscription
func (mr *MockCRMProviderMockRecorder) PutSubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSubscription", reflect.TypeOf((*MockCRMProvider)(nil).PutSubscription), arg0)
}

// PutSubscriptions mocks base method
func (m *MockCRMProvider) PutSubscriptions(arg0 []entity.CRMEmailListSubscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSubscriptions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSubscriptions indicates an expected call of PutSubscriptions
func (mr *MockCRMProviderMockRecorder) PutSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSubscriptions", reflect.TypeOf((*MockCRMProvider)(nil).PutSubscriptions), arg0)
}

// GetSubscriptionsOfList mocks base method
func (m *MockCRMProvider) GetSubscriptionsOfList(listID entity.ID) ([]entity.CRMEmailListSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionsOfList", listID)
	ret0, _ := ret[0].([]entity.CRMEmailListSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionsOfList indicates an expected call of GetSubscriptionsOfList
func (mr *MockCRMProviderMockRecorder) GetSubscriptionsOfList(listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionsOfList", reflect.TypeOf((*MockCRMProvider)(nil).GetSubscriptionsOfList), listID)
}