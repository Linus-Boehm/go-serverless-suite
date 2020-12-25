// Code generated by MockGen. DO NOT EDIT.
// Source: mailings.go

// Package itf is a generated GoMock package.
package itf

import (
	entity "github.com/Linus-Boehm/go-serverless-suite/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTplRenderer is a mock of TplRenderer interface
type MockTplRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockTplRendererMockRecorder
}

// MockTplRendererMockRecorder is the mock recorder for MockTplRenderer
type MockTplRendererMockRecorder struct {
	mock *MockTplRenderer
}

// NewMockTplRenderer creates a new mock instance
func NewMockTplRenderer(ctrl *gomock.Controller) *MockTplRenderer {
	mock := &MockTplRenderer{ctrl: ctrl}
	mock.recorder = &MockTplRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTplRenderer) EXPECT() *MockTplRendererMockRecorder {
	return m.recorder
}

// GetRaw mocks base method
func (m *MockTplRenderer) GetRaw() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRaw")
	ret0, _ := ret[0].(*string)
	return ret0
}

// GetRaw indicates an expected call of GetRaw
func (mr *MockTplRendererMockRecorder) GetRaw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRaw", reflect.TypeOf((*MockTplRenderer)(nil).GetRaw))
}

// Render mocks base method
func (m *MockTplRenderer) Render(data interface{}) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", data)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render
func (mr *MockTplRendererMockRecorder) Render(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockTplRenderer)(nil).Render), data)
}

// RenderWithHTML mocks base method
func (m *MockTplRenderer) RenderWithHTML(data interface{}) (*entity.HTMLTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderWithHTML", data)
	ret0, _ := ret[0].(*entity.HTMLTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderWithHTML indicates an expected call of RenderWithHTML
func (mr *MockTplRendererMockRecorder) RenderWithHTML(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderWithHTML", reflect.TypeOf((*MockTplRenderer)(nil).RenderWithHTML), data)
}

// MockMailer is a mock of Mailer interface
type MockMailer struct {
	ctrl     *gomock.Controller
	recorder *MockMailerMockRecorder
}

// MockMailerMockRecorder is the mock recorder for MockMailer
type MockMailerMockRecorder struct {
	mock *MockMailer
}

// NewMockMailer creates a new mock instance
func NewMockMailer(ctrl *gomock.Controller) *MockMailer {
	mock := &MockMailer{ctrl: ctrl}
	mock.recorder = &MockMailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMailer) EXPECT() *MockMailerMockRecorder {
	return m.recorder
}

// SendSimpleContactForm mocks base method
func (m *MockMailer) SendSimpleContactForm(input entity.ContactForm, renderer TplRenderer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSimpleContactForm", input, renderer)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSimpleContactForm indicates an expected call of SendSimpleContactForm
func (mr *MockMailerMockRecorder) SendSimpleContactForm(input, renderer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSimpleContactForm", reflect.TypeOf((*MockMailer)(nil).SendSimpleContactForm), input, renderer)
}

// GetContactLists mocks base method
func (m *MockMailer) GetContactLists() ([]entity.ContactList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactLists")
	ret0, _ := ret[0].([]entity.ContactList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactLists indicates an expected call of GetContactLists
func (mr *MockMailerMockRecorder) GetContactLists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactLists", reflect.TypeOf((*MockMailer)(nil).GetContactLists))
}

// GetProvider mocks base method
func (m *MockMailer) GetProvider() MailerProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProvider")
	ret0, _ := ret[0].(MailerProvider)
	return ret0
}

// GetProvider indicates an expected call of GetProvider
func (mr *MockMailerMockRecorder) GetProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProvider", reflect.TypeOf((*MockMailer)(nil).GetProvider))
}

// MockMailerProvider is a mock of MailerProvider interface
type MockMailerProvider struct {
	ctrl     *gomock.Controller
	recorder *MockMailerProviderMockRecorder
}

// MockMailerProviderMockRecorder is the mock recorder for MockMailerProvider
type MockMailerProviderMockRecorder struct {
	mock *MockMailerProvider
}

// NewMockMailerProvider creates a new mock instance
func NewMockMailerProvider(ctrl *gomock.Controller) *MockMailerProvider {
	mock := &MockMailerProvider{ctrl: ctrl}
	mock.recorder = &MockMailerProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMailerProvider) EXPECT() *MockMailerProviderMockRecorder {
	return m.recorder
}

// SendSingleMail mocks base method
func (m *MockMailerProvider) SendSingleMail(input entity.MinimalMail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSingleMail", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSingleMail indicates an expected call of SendSingleMail
func (mr *MockMailerProviderMockRecorder) SendSingleMail(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSingleMail", reflect.TypeOf((*MockMailerProvider)(nil).SendSingleMail), input)
}

// GetContactLists mocks base method
func (m *MockMailerProvider) GetContactLists() ([]entity.ContactList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactLists")
	ret0, _ := ret[0].([]entity.ContactList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactLists indicates an expected call of GetContactLists
func (mr *MockMailerProviderMockRecorder) GetContactLists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactLists", reflect.TypeOf((*MockMailerProvider)(nil).GetContactLists))
}
