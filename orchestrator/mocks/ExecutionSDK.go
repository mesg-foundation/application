// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import execution "github.com/mesg-foundation/engine/execution"
import executionsdk "github.com/mesg-foundation/engine/sdk/execution"
import hash "github.com/mesg-foundation/engine/hash"
import mock "github.com/stretchr/testify/mock"

import types "github.com/mesg-foundation/engine/protobuf/types"

// ExecutionSDK is an autogenerated mock type for the ExecutionSDK type
type ExecutionSDK struct {
	mock.Mock
}

// Execute provides a mock function with given fields: processHash, instanceHash, eventHash, parentHash, stepID, taskKey, inputData, tags
func (_m *ExecutionSDK) Execute(processHash hash.Hash, instanceHash hash.Hash, eventHash hash.Hash, parentHash hash.Hash, stepID string, taskKey string, inputData *types.Struct, tags []string) (hash.Hash, error) {
	ret := _m.Called(processHash, instanceHash, eventHash, parentHash, stepID, taskKey, inputData, tags)

	var r0 hash.Hash
	if rf, ok := ret.Get(0).(func(hash.Hash, hash.Hash, hash.Hash, hash.Hash, string, string, *types.Struct, []string) hash.Hash); ok {
		r0 = rf(processHash, instanceHash, eventHash, parentHash, stepID, taskKey, inputData, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hash.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(hash.Hash, hash.Hash, hash.Hash, hash.Hash, string, string, *types.Struct, []string) error); ok {
		r1 = rf(processHash, instanceHash, eventHash, parentHash, stepID, taskKey, inputData, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0
func (_m *ExecutionSDK) Get(_a0 hash.Hash) (*execution.Execution, error) {
	ret := _m.Called(_a0)

	var r0 *execution.Execution
	if rf, ok := ret.Get(0).(func(hash.Hash) *execution.Execution); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(hash.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStream provides a mock function with given fields: f
func (_m *ExecutionSDK) GetStream(f *executionsdk.Filter) *executionsdk.Listener {
	ret := _m.Called(f)

	var r0 *executionsdk.Listener
	if rf, ok := ret.Get(0).(func(*executionsdk.Filter) *executionsdk.Listener); ok {
		r0 = rf(f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executionsdk.Listener)
		}
	}

	return r0
}
