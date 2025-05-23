// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dcbltech/go-kit/email (interfaces: Emailer)
//
// Generated by this command:
//
//	mockgen -destination email/mock/email.go -package mock github.com/dcbltech/go-kit/email Emailer
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEmailer is a mock of Emailer interface.
type MockEmailer struct {
	ctrl     *gomock.Controller
	recorder *MockEmailerMockRecorder
	isgomock struct{}
}

// MockEmailerMockRecorder is the mock recorder for MockEmailer.
type MockEmailerMockRecorder struct {
	mock *MockEmailer
}

// NewMockEmailer creates a new mock instance.
func NewMockEmailer(ctrl *gomock.Controller) *MockEmailer {
	mock := &MockEmailer{ctrl: ctrl}
	mock.recorder = &MockEmailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailer) EXPECT() *MockEmailerMockRecorder {
	return m.recorder
}

// SendTemplateEmail mocks base method.
func (m *MockEmailer) SendTemplateEmail(ctx context.Context, template int, name, email string, data map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTemplateEmail", ctx, template, name, email, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTemplateEmail indicates an expected call of SendTemplateEmail.
func (mr *MockEmailerMockRecorder) SendTemplateEmail(ctx, template, name, email, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTemplateEmail", reflect.TypeOf((*MockEmailer)(nil).SendTemplateEmail), ctx, template, name, email, data)
}
