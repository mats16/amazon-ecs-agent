// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/taskresource (interfaces: TaskResource)

// Package mock_taskresource is a generated GoMock package.
package mock_taskresource

import (
	reflect "reflect"
	time "time"

	taskresource "github.com/aws/amazon-ecs-agent/agent/taskresource"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskResource is a mock of TaskResource interface
type MockTaskResource struct {
	ctrl     *gomock.Controller
	recorder *MockTaskResourceMockRecorder
}

// MockTaskResourceMockRecorder is the mock recorder for MockTaskResource
type MockTaskResourceMockRecorder struct {
	mock *MockTaskResource
}

// NewMockTaskResource creates a new mock instance
func NewMockTaskResource(ctrl *gomock.Controller) *MockTaskResource {
	mock := &MockTaskResource{ctrl: ctrl}
	mock.recorder = &MockTaskResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskResource) EXPECT() *MockTaskResourceMockRecorder {
	return m.recorder
}

// Cleanup mocks base method
func (m *MockTaskResource) Cleanup() error {
	ret := m.ctrl.Call(m, "Cleanup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup
func (mr *MockTaskResourceMockRecorder) Cleanup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockTaskResource)(nil).Cleanup))
}

// Create mocks base method
func (m *MockTaskResource) Create() error {
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTaskResourceMockRecorder) Create() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskResource)(nil).Create))
}

// GetCreatedAt mocks base method
func (m *MockTaskResource) GetCreatedAt() time.Time {
	ret := m.ctrl.Call(m, "GetCreatedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCreatedAt indicates an expected call of GetCreatedAt
func (mr *MockTaskResourceMockRecorder) GetCreatedAt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedAt", reflect.TypeOf((*MockTaskResource)(nil).GetCreatedAt))
}

// GetDesiredStatus mocks base method
func (m *MockTaskResource) GetDesiredStatus() taskresource.ResourceStatus {
	ret := m.ctrl.Call(m, "GetDesiredStatus")
	ret0, _ := ret[0].(taskresource.ResourceStatus)
	return ret0
}

// GetDesiredStatus indicates an expected call of GetDesiredStatus
func (mr *MockTaskResourceMockRecorder) GetDesiredStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDesiredStatus", reflect.TypeOf((*MockTaskResource)(nil).GetDesiredStatus))
}

// GetKnownStatus mocks base method
func (m *MockTaskResource) GetKnownStatus() taskresource.ResourceStatus {
	ret := m.ctrl.Call(m, "GetKnownStatus")
	ret0, _ := ret[0].(taskresource.ResourceStatus)
	return ret0
}

// GetKnownStatus indicates an expected call of GetKnownStatus
func (mr *MockTaskResourceMockRecorder) GetKnownStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnownStatus", reflect.TypeOf((*MockTaskResource)(nil).GetKnownStatus))
}

// GetName mocks base method
func (m *MockTaskResource) GetName() string {
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockTaskResourceMockRecorder) GetName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockTaskResource)(nil).GetName))
}

// MarshalJSON mocks base method
func (m *MockTaskResource) MarshalJSON() ([]byte, error) {
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockTaskResourceMockRecorder) MarshalJSON() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockTaskResource)(nil).MarshalJSON))
}

// SetCreatedAt mocks base method
func (m *MockTaskResource) SetCreatedAt(arg0 time.Time) {
	m.ctrl.Call(m, "SetCreatedAt", arg0)
}

// SetCreatedAt indicates an expected call of SetCreatedAt
func (mr *MockTaskResourceMockRecorder) SetCreatedAt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreatedAt", reflect.TypeOf((*MockTaskResource)(nil).SetCreatedAt), arg0)
}

// SetDesiredStatus mocks base method
func (m *MockTaskResource) SetDesiredStatus(arg0 taskresource.ResourceStatus) {
	m.ctrl.Call(m, "SetDesiredStatus", arg0)
}

// SetDesiredStatus indicates an expected call of SetDesiredStatus
func (mr *MockTaskResourceMockRecorder) SetDesiredStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDesiredStatus", reflect.TypeOf((*MockTaskResource)(nil).SetDesiredStatus), arg0)
}

// SetKnownStatus mocks base method
func (m *MockTaskResource) SetKnownStatus(arg0 taskresource.ResourceStatus) {
	m.ctrl.Call(m, "SetKnownStatus", arg0)
}

// SetKnownStatus indicates an expected call of SetKnownStatus
func (mr *MockTaskResourceMockRecorder) SetKnownStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKnownStatus", reflect.TypeOf((*MockTaskResource)(nil).SetKnownStatus), arg0)
}

// UnmarshalJSON mocks base method
func (m *MockTaskResource) UnmarshalJSON(arg0 []byte) error {
	ret := m.ctrl.Call(m, "UnmarshalJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalJSON indicates an expected call of UnmarshalJSON
func (mr *MockTaskResourceMockRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalJSON", reflect.TypeOf((*MockTaskResource)(nil).UnmarshalJSON), arg0)
}
