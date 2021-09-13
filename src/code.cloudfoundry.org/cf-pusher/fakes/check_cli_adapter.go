// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type CheckCLIAdapter struct {
	AppCountStub        func(string) (int, error)
	appCountMutex       sync.RWMutex
	appCountArgsForCall []struct {
		arg1 string
	}
	appCountReturns struct {
		result1 int
		result2 error
	}
	appCountReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	AppGuidStub        func(string) (string, error)
	appGuidMutex       sync.RWMutex
	appGuidArgsForCall []struct {
		arg1 string
	}
	appGuidReturns struct {
		result1 string
		result2 error
	}
	appGuidReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CheckAppStub        func(string) ([]byte, error)
	checkAppMutex       sync.RWMutex
	checkAppArgsForCall []struct {
		arg1 string
	}
	checkAppReturns struct {
		result1 []byte
		result2 error
	}
	checkAppReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	OrgGuidStub        func(string) (string, error)
	orgGuidMutex       sync.RWMutex
	orgGuidArgsForCall []struct {
		arg1 string
	}
	orgGuidReturns struct {
		result1 string
		result2 error
	}
	orgGuidReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CheckCLIAdapter) AppCount(arg1 string) (int, error) {
	fake.appCountMutex.Lock()
	ret, specificReturn := fake.appCountReturnsOnCall[len(fake.appCountArgsForCall)]
	fake.appCountArgsForCall = append(fake.appCountArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AppCountStub
	fakeReturns := fake.appCountReturns
	fake.recordInvocation("AppCount", []interface{}{arg1})
	fake.appCountMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CheckCLIAdapter) AppCountCallCount() int {
	fake.appCountMutex.RLock()
	defer fake.appCountMutex.RUnlock()
	return len(fake.appCountArgsForCall)
}

func (fake *CheckCLIAdapter) AppCountCalls(stub func(string) (int, error)) {
	fake.appCountMutex.Lock()
	defer fake.appCountMutex.Unlock()
	fake.AppCountStub = stub
}

func (fake *CheckCLIAdapter) AppCountArgsForCall(i int) string {
	fake.appCountMutex.RLock()
	defer fake.appCountMutex.RUnlock()
	argsForCall := fake.appCountArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CheckCLIAdapter) AppCountReturns(result1 int, result2 error) {
	fake.appCountMutex.Lock()
	defer fake.appCountMutex.Unlock()
	fake.AppCountStub = nil
	fake.appCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) AppCountReturnsOnCall(i int, result1 int, result2 error) {
	fake.appCountMutex.Lock()
	defer fake.appCountMutex.Unlock()
	fake.AppCountStub = nil
	if fake.appCountReturnsOnCall == nil {
		fake.appCountReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.appCountReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) AppGuid(arg1 string) (string, error) {
	fake.appGuidMutex.Lock()
	ret, specificReturn := fake.appGuidReturnsOnCall[len(fake.appGuidArgsForCall)]
	fake.appGuidArgsForCall = append(fake.appGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AppGuidStub
	fakeReturns := fake.appGuidReturns
	fake.recordInvocation("AppGuid", []interface{}{arg1})
	fake.appGuidMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CheckCLIAdapter) AppGuidCallCount() int {
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	return len(fake.appGuidArgsForCall)
}

func (fake *CheckCLIAdapter) AppGuidCalls(stub func(string) (string, error)) {
	fake.appGuidMutex.Lock()
	defer fake.appGuidMutex.Unlock()
	fake.AppGuidStub = stub
}

func (fake *CheckCLIAdapter) AppGuidArgsForCall(i int) string {
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	argsForCall := fake.appGuidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CheckCLIAdapter) AppGuidReturns(result1 string, result2 error) {
	fake.appGuidMutex.Lock()
	defer fake.appGuidMutex.Unlock()
	fake.AppGuidStub = nil
	fake.appGuidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) AppGuidReturnsOnCall(i int, result1 string, result2 error) {
	fake.appGuidMutex.Lock()
	defer fake.appGuidMutex.Unlock()
	fake.AppGuidStub = nil
	if fake.appGuidReturnsOnCall == nil {
		fake.appGuidReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.appGuidReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) CheckApp(arg1 string) ([]byte, error) {
	fake.checkAppMutex.Lock()
	ret, specificReturn := fake.checkAppReturnsOnCall[len(fake.checkAppArgsForCall)]
	fake.checkAppArgsForCall = append(fake.checkAppArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CheckAppStub
	fakeReturns := fake.checkAppReturns
	fake.recordInvocation("CheckApp", []interface{}{arg1})
	fake.checkAppMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CheckCLIAdapter) CheckAppCallCount() int {
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	return len(fake.checkAppArgsForCall)
}

func (fake *CheckCLIAdapter) CheckAppCalls(stub func(string) ([]byte, error)) {
	fake.checkAppMutex.Lock()
	defer fake.checkAppMutex.Unlock()
	fake.CheckAppStub = stub
}

func (fake *CheckCLIAdapter) CheckAppArgsForCall(i int) string {
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	argsForCall := fake.checkAppArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CheckCLIAdapter) CheckAppReturns(result1 []byte, result2 error) {
	fake.checkAppMutex.Lock()
	defer fake.checkAppMutex.Unlock()
	fake.CheckAppStub = nil
	fake.checkAppReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) CheckAppReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.checkAppMutex.Lock()
	defer fake.checkAppMutex.Unlock()
	fake.CheckAppStub = nil
	if fake.checkAppReturnsOnCall == nil {
		fake.checkAppReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.checkAppReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) OrgGuid(arg1 string) (string, error) {
	fake.orgGuidMutex.Lock()
	ret, specificReturn := fake.orgGuidReturnsOnCall[len(fake.orgGuidArgsForCall)]
	fake.orgGuidArgsForCall = append(fake.orgGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.OrgGuidStub
	fakeReturns := fake.orgGuidReturns
	fake.recordInvocation("OrgGuid", []interface{}{arg1})
	fake.orgGuidMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CheckCLIAdapter) OrgGuidCallCount() int {
	fake.orgGuidMutex.RLock()
	defer fake.orgGuidMutex.RUnlock()
	return len(fake.orgGuidArgsForCall)
}

func (fake *CheckCLIAdapter) OrgGuidCalls(stub func(string) (string, error)) {
	fake.orgGuidMutex.Lock()
	defer fake.orgGuidMutex.Unlock()
	fake.OrgGuidStub = stub
}

func (fake *CheckCLIAdapter) OrgGuidArgsForCall(i int) string {
	fake.orgGuidMutex.RLock()
	defer fake.orgGuidMutex.RUnlock()
	argsForCall := fake.orgGuidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CheckCLIAdapter) OrgGuidReturns(result1 string, result2 error) {
	fake.orgGuidMutex.Lock()
	defer fake.orgGuidMutex.Unlock()
	fake.OrgGuidStub = nil
	fake.orgGuidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) OrgGuidReturnsOnCall(i int, result1 string, result2 error) {
	fake.orgGuidMutex.Lock()
	defer fake.orgGuidMutex.Unlock()
	fake.OrgGuidStub = nil
	if fake.orgGuidReturnsOnCall == nil {
		fake.orgGuidReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.orgGuidReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CheckCLIAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appCountMutex.RLock()
	defer fake.appCountMutex.RUnlock()
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	fake.orgGuidMutex.RLock()
	defer fake.orgGuidMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CheckCLIAdapter) recordInvocation(key string, args []interface{}) {
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