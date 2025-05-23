// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dcbltech/go-kit/tts (interfaces: TextToSpeech)
//
// Generated by this command:
//
//	mockgen -destination tts/mock/tts.go -package mock github.com/dcbltech/go-kit/tts TextToSpeech
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTextToSpeech is a mock of TextToSpeech interface.
type MockTextToSpeech struct {
	ctrl     *gomock.Controller
	recorder *MockTextToSpeechMockRecorder
	isgomock struct{}
}

// MockTextToSpeechMockRecorder is the mock recorder for MockTextToSpeech.
type MockTextToSpeechMockRecorder struct {
	mock *MockTextToSpeech
}

// NewMockTextToSpeech creates a new mock instance.
func NewMockTextToSpeech(ctrl *gomock.Controller) *MockTextToSpeech {
	mock := &MockTextToSpeech{ctrl: ctrl}
	mock.recorder = &MockTextToSpeechMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTextToSpeech) EXPECT() *MockTextToSpeechMockRecorder {
	return m.recorder
}

// Synthesize mocks base method.
func (m *MockTextToSpeech) Synthesize(text, voice string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Synthesize", text, voice)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Synthesize indicates an expected call of Synthesize.
func (mr *MockTextToSpeechMockRecorder) Synthesize(text, voice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Synthesize", reflect.TypeOf((*MockTextToSpeech)(nil).Synthesize), text, voice)
}
