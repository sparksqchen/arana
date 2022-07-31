/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package pkg is a generated GoMock package.
package testdata

import (
	reflect "reflect"
)

import (
	gomock "github.com/golang/mock/gomock"
)

import (
	"github.com/arana-db/arana/pkg/config"
)

// MockStoreOperate is a mock of StoreOperate interface.
type MockStoreOperate struct {
	ctrl     *gomock.Controller
	recorder *MockStoreOperateMockRecorder
}

// MockStoreOperateMockRecorder is the mock recorder for MockStoreOperate.
type MockStoreOperateMockRecorder struct {
	mock *MockStoreOperate
}

// NewMockStoreOperate creates a new mock instance.
func NewMockStoreOperate(ctrl *gomock.Controller) *MockStoreOperate {
	mock := &MockStoreOperate{ctrl: ctrl}
	mock.recorder = &MockStoreOperateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreOperate) EXPECT() *MockStoreOperateMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStoreOperate) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStoreOperateMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStoreOperate)(nil).Close))
}

// Get mocks base method.
func (m *MockStoreOperate) Get(key config.PathKey) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreOperateMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStoreOperate)(nil).Get), key)
}

// Init mocks base method.
func (m *MockStoreOperate) Init(options map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockStoreOperateMockRecorder) Init(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockStoreOperate)(nil).Init), options)
}

// Name mocks base method.
func (m *MockStoreOperate) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockStoreOperateMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockStoreOperate)(nil).Name))
}

// Save mocks base method.
func (m *MockStoreOperate) Save(key config.PathKey, val []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStoreOperateMockRecorder) Save(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStoreOperate)(nil).Save), key, val)
}

// Watch mocks base method.
func (m *MockStoreOperate) Watch(key config.PathKey) (<-chan []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", key)
	ret0, _ := ret[0].(<-chan []byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockStoreOperateMockRecorder) Watch(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockStoreOperate)(nil).Watch), key)
}
