// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v6"
)

type FakeTargetActor struct {
	GetOrganizationByNameStub        func(orgName string) (v2action.Organization, v2action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		orgName string
	}
	getOrganizationByNameReturns struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationSpacesStub        func(orgGUID string) ([]v2action.Space, v2action.Warnings, error)
	getOrganizationSpacesMutex       sync.RWMutex
	getOrganizationSpacesArgsForCall []struct {
		orgGUID string
	}
	getOrganizationSpacesReturns struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationSpacesReturnsOnCall map[int]struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	GetSpaceByOrganizationAndNameStub        func(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error)
	getSpaceByOrganizationAndNameMutex       sync.RWMutex
	getSpaceByOrganizationAndNameArgsForCall []struct {
		orgGUID   string
		spaceName string
	}
	getSpaceByOrganizationAndNameReturns struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	getSpaceByOrganizationAndNameReturnsOnCall map[int]struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTargetActor) GetOrganizationByName(orgName string) (v2action.Organization, v2action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrganizationByName", []interface{}{orgName})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationByNameReturns.result1, fake.getOrganizationByNameReturns.result2, fake.getOrganizationByNameReturns.result3
}

func (fake *FakeTargetActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeTargetActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return fake.getOrganizationByNameArgsForCall[i].orgName
}

func (fake *FakeTargetActor) GetOrganizationByNameReturns(result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTargetActor) GetOrganizationByNameReturnsOnCall(i int, result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Organization
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTargetActor) GetOrganizationSpaces(orgGUID string) ([]v2action.Space, v2action.Warnings, error) {
	fake.getOrganizationSpacesMutex.Lock()
	ret, specificReturn := fake.getOrganizationSpacesReturnsOnCall[len(fake.getOrganizationSpacesArgsForCall)]
	fake.getOrganizationSpacesArgsForCall = append(fake.getOrganizationSpacesArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("GetOrganizationSpaces", []interface{}{orgGUID})
	fake.getOrganizationSpacesMutex.Unlock()
	if fake.GetOrganizationSpacesStub != nil {
		return fake.GetOrganizationSpacesStub(orgGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationSpacesReturns.result1, fake.getOrganizationSpacesReturns.result2, fake.getOrganizationSpacesReturns.result3
}

func (fake *FakeTargetActor) GetOrganizationSpacesCallCount() int {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return len(fake.getOrganizationSpacesArgsForCall)
}

func (fake *FakeTargetActor) GetOrganizationSpacesArgsForCall(i int) string {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return fake.getOrganizationSpacesArgsForCall[i].orgGUID
}

func (fake *FakeTargetActor) GetOrganizationSpacesReturns(result1 []v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationSpacesStub = nil
	fake.getOrganizationSpacesReturns = struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTargetActor) GetOrganizationSpacesReturnsOnCall(i int, result1 []v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationSpacesStub = nil
	if fake.getOrganizationSpacesReturnsOnCall == nil {
		fake.getOrganizationSpacesReturnsOnCall = make(map[int]struct {
			result1 []v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationSpacesReturnsOnCall[i] = struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTargetActor) GetSpaceByOrganizationAndName(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error) {
	fake.getSpaceByOrganizationAndNameMutex.Lock()
	ret, specificReturn := fake.getSpaceByOrganizationAndNameReturnsOnCall[len(fake.getSpaceByOrganizationAndNameArgsForCall)]
	fake.getSpaceByOrganizationAndNameArgsForCall = append(fake.getSpaceByOrganizationAndNameArgsForCall, struct {
		orgGUID   string
		spaceName string
	}{orgGUID, spaceName})
	fake.recordInvocation("GetSpaceByOrganizationAndName", []interface{}{orgGUID, spaceName})
	fake.getSpaceByOrganizationAndNameMutex.Unlock()
	if fake.GetSpaceByOrganizationAndNameStub != nil {
		return fake.GetSpaceByOrganizationAndNameStub(orgGUID, spaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSpaceByOrganizationAndNameReturns.result1, fake.getSpaceByOrganizationAndNameReturns.result2, fake.getSpaceByOrganizationAndNameReturns.result3
}

func (fake *FakeTargetActor) GetSpaceByOrganizationAndNameCallCount() int {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return len(fake.getSpaceByOrganizationAndNameArgsForCall)
}

func (fake *FakeTargetActor) GetSpaceByOrganizationAndNameArgsForCall(i int) (string, string) {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return fake.getSpaceByOrganizationAndNameArgsForCall[i].orgGUID, fake.getSpaceByOrganizationAndNameArgsForCall[i].spaceName
}

func (fake *FakeTargetActor) GetSpaceByOrganizationAndNameReturns(result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	fake.getSpaceByOrganizationAndNameReturns = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTargetActor) GetSpaceByOrganizationAndNameReturnsOnCall(i int, result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	if fake.getSpaceByOrganizationAndNameReturnsOnCall == nil {
		fake.getSpaceByOrganizationAndNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getSpaceByOrganizationAndNameReturnsOnCall[i] = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTargetActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTargetActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.TargetActor = new(FakeTargetActor)