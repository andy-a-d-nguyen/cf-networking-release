// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/store"
)

type EgressPolicyStore struct {
	AllStub        func() ([]store.EgressPolicy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
	}
	allReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	CreateStub        func([]store.EgressPolicy) ([]store.EgressPolicy, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []store.EgressPolicy
	}
	createReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	DeleteStub        func(...string) ([]store.EgressPolicy, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []string
	}
	deleteReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByFilterStub        func([]string, []string, []string, []string, []string) ([]store.EgressPolicy, error)
	getByFilterMutex       sync.RWMutex
	getByFilterArgsForCall []struct {
		arg1 []string
		arg2 []string
		arg3 []string
		arg4 []string
		arg5 []string
	}
	getByFilterReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByFilterReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetBySourceGuidsAndDefaultsStub        func([]string) ([]store.EgressPolicy, error)
	getBySourceGuidsAndDefaultsMutex       sync.RWMutex
	getBySourceGuidsAndDefaultsArgsForCall []struct {
		arg1 []string
	}
	getBySourceGuidsAndDefaultsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getBySourceGuidsAndDefaultsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyStore) All() ([]store.EgressPolicy, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct {
	}{})
	stub := fake.AllStub
	fakeReturns := fake.allReturns
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyStore) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressPolicyStore) AllCalls(stub func() ([]store.EgressPolicy, error)) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = stub
}

func (fake *EgressPolicyStore) AllReturns(result1 []store.EgressPolicy, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) AllReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Create(arg1 []store.EgressPolicy) ([]store.EgressPolicy, error) {
	var arg1Copy []store.EgressPolicy
	if arg1 != nil {
		arg1Copy = make([]store.EgressPolicy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []store.EgressPolicy
	}{arg1Copy})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *EgressPolicyStore) CreateCalls(stub func([]store.EgressPolicy) ([]store.EgressPolicy, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *EgressPolicyStore) CreateArgsForCall(i int) []store.EgressPolicy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressPolicyStore) CreateReturns(result1 []store.EgressPolicy, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) CreateReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Delete(arg1 ...string) ([]store.EgressPolicy, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *EgressPolicyStore) DeleteCalls(stub func(...string) ([]store.EgressPolicy, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *EgressPolicyStore) DeleteArgsForCall(i int) []string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressPolicyStore) DeleteReturns(result1 []store.EgressPolicy, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) DeleteReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetByFilter(arg1 []string, arg2 []string, arg3 []string, arg4 []string, arg5 []string) ([]store.EgressPolicy, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	var arg4Copy []string
	if arg4 != nil {
		arg4Copy = make([]string, len(arg4))
		copy(arg4Copy, arg4)
	}
	var arg5Copy []string
	if arg5 != nil {
		arg5Copy = make([]string, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.getByFilterMutex.Lock()
	ret, specificReturn := fake.getByFilterReturnsOnCall[len(fake.getByFilterArgsForCall)]
	fake.getByFilterArgsForCall = append(fake.getByFilterArgsForCall, struct {
		arg1 []string
		arg2 []string
		arg3 []string
		arg4 []string
		arg5 []string
	}{arg1Copy, arg2Copy, arg3Copy, arg4Copy, arg5Copy})
	stub := fake.GetByFilterStub
	fakeReturns := fake.getByFilterReturns
	fake.recordInvocation("GetByFilter", []interface{}{arg1Copy, arg2Copy, arg3Copy, arg4Copy, arg5Copy})
	fake.getByFilterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyStore) GetByFilterCallCount() int {
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	return len(fake.getByFilterArgsForCall)
}

func (fake *EgressPolicyStore) GetByFilterCalls(stub func([]string, []string, []string, []string, []string) ([]store.EgressPolicy, error)) {
	fake.getByFilterMutex.Lock()
	defer fake.getByFilterMutex.Unlock()
	fake.GetByFilterStub = stub
}

func (fake *EgressPolicyStore) GetByFilterArgsForCall(i int) ([]string, []string, []string, []string, []string) {
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	argsForCall := fake.getByFilterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *EgressPolicyStore) GetByFilterReturns(result1 []store.EgressPolicy, result2 error) {
	fake.getByFilterMutex.Lock()
	defer fake.getByFilterMutex.Unlock()
	fake.GetByFilterStub = nil
	fake.getByFilterReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetByFilterReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.getByFilterMutex.Lock()
	defer fake.getByFilterMutex.Unlock()
	fake.GetByFilterStub = nil
	if fake.getByFilterReturnsOnCall == nil {
		fake.getByFilterReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByFilterReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetBySourceGuidsAndDefaults(arg1 []string) ([]store.EgressPolicy, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	ret, specificReturn := fake.getBySourceGuidsAndDefaultsReturnsOnCall[len(fake.getBySourceGuidsAndDefaultsArgsForCall)]
	fake.getBySourceGuidsAndDefaultsArgsForCall = append(fake.getBySourceGuidsAndDefaultsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.GetBySourceGuidsAndDefaultsStub
	fakeReturns := fake.getBySourceGuidsAndDefaultsReturns
	fake.recordInvocation("GetBySourceGuidsAndDefaults", []interface{}{arg1Copy})
	fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyStore) GetBySourceGuidsAndDefaultsCallCount() int {
	fake.getBySourceGuidsAndDefaultsMutex.RLock()
	defer fake.getBySourceGuidsAndDefaultsMutex.RUnlock()
	return len(fake.getBySourceGuidsAndDefaultsArgsForCall)
}

func (fake *EgressPolicyStore) GetBySourceGuidsAndDefaultsCalls(stub func([]string) ([]store.EgressPolicy, error)) {
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	defer fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	fake.GetBySourceGuidsAndDefaultsStub = stub
}

func (fake *EgressPolicyStore) GetBySourceGuidsAndDefaultsArgsForCall(i int) []string {
	fake.getBySourceGuidsAndDefaultsMutex.RLock()
	defer fake.getBySourceGuidsAndDefaultsMutex.RUnlock()
	argsForCall := fake.getBySourceGuidsAndDefaultsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressPolicyStore) GetBySourceGuidsAndDefaultsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	defer fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	fake.GetBySourceGuidsAndDefaultsStub = nil
	fake.getBySourceGuidsAndDefaultsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetBySourceGuidsAndDefaultsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	defer fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	fake.GetBySourceGuidsAndDefaultsStub = nil
	if fake.getBySourceGuidsAndDefaultsReturnsOnCall == nil {
		fake.getBySourceGuidsAndDefaultsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getBySourceGuidsAndDefaultsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	fake.getBySourceGuidsAndDefaultsMutex.RLock()
	defer fake.getBySourceGuidsAndDefaultsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyStore) recordInvocation(key string, args []interface{}) {
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