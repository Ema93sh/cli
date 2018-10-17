// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeUnsetEnvActor struct {
	UnsetEnvironmentVariableByApplicationNameAndSpaceStub        func(string, string, string) (v7action.Warnings, error)
	unsetEnvironmentVariableByApplicationNameAndSpaceMutex       sync.RWMutex
	unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	unsetEnvironmentVariableByApplicationNameAndSpaceReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpace(arg1 string, arg2 string, arg3 string) (v7action.Warnings, error) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall[len(fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall)]
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall = append(fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("UnsetEnvironmentVariableByApplicationNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	if fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub != nil {
		return fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceCallCount() int {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RLock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeUnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceCalls(stub func(string, string, string) (v7action.Warnings, error)) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub = stub
}

func (fake *FakeUnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall(i int) (string, string, string) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RLock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RUnlock()
	argsForCall := fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceReturns(result1 v7action.Warnings, result2 error) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub = nil
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub = nil
	if fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetEnvActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RLock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnsetEnvActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v7.UnsetEnvActor = new(FakeUnsetEnvActor)