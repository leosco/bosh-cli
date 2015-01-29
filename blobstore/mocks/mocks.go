// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-micro-cli/blobstore (interfaces: Factory,Blobstore)

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
	blobstore "github.com/cloudfoundry/bosh-micro-cli/blobstore"
)

// Mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *_MockFactoryRecorder
}

// Recorder for MockFactory (not exported)
type _MockFactoryRecorder struct {
	mock *MockFactory
}

func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &_MockFactoryRecorder{mock}
	return mock
}

func (_m *MockFactory) EXPECT() *_MockFactoryRecorder {
	return _m.recorder
}

func (_m *MockFactory) Create(_param0 string) (blobstore.Blobstore, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(blobstore.Blobstore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFactoryRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

// Mock of Blobstore interface
type MockBlobstore struct {
	ctrl     *gomock.Controller
	recorder *_MockBlobstoreRecorder
}

// Recorder for MockBlobstore (not exported)
type _MockBlobstoreRecorder struct {
	mock *MockBlobstore
}

func NewMockBlobstore(ctrl *gomock.Controller) *MockBlobstore {
	mock := &MockBlobstore{ctrl: ctrl}
	mock.recorder = &_MockBlobstoreRecorder{mock}
	return mock
}

func (_m *MockBlobstore) EXPECT() *_MockBlobstoreRecorder {
	return _m.recorder
}

func (_m *MockBlobstore) Add(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "Add", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBlobstoreRecorder) Add(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0)
}

func (_m *MockBlobstore) Get(_param0 string) (blobstore.LocalBlob, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(blobstore.LocalBlob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBlobstoreRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}
