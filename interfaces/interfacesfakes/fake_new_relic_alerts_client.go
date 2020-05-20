// Code generated by counterfeiter. DO NOT EDIT.
package interfacesfakes

import (
	"sync"

	"github.com/newrelic/newrelic-client-go/pkg/alerts"

	"github.com/newrelic/newrelic-kubernetes-operator/interfaces"
)

type FakeNewRelicAlertsClient struct {
	CreateConditionStub        func(int, alerts.Condition) (*alerts.Condition, error)
	createConditionMutex       sync.RWMutex
	createConditionArgsForCall []struct {
		arg1 int
		arg2 alerts.Condition
	}
	createConditionReturns struct {
		result1 *alerts.Condition
		result2 error
	}
	createConditionReturnsOnCall map[int]struct {
		result1 *alerts.Condition
		result2 error
	}
	CreateNrqlConditionStub        func(int, alerts.NrqlCondition) (*alerts.NrqlCondition, error)
	createNrqlConditionMutex       sync.RWMutex
	createNrqlConditionArgsForCall []struct {
		arg1 int
		arg2 alerts.NrqlCondition
	}
	createNrqlConditionReturns struct {
		result1 *alerts.NrqlCondition
		result2 error
	}
	createNrqlConditionReturnsOnCall map[int]struct {
		result1 *alerts.NrqlCondition
		result2 error
	}
	CreatePolicyStub        func(alerts.Policy) (*alerts.Policy, error)
	createPolicyMutex       sync.RWMutex
	createPolicyArgsForCall []struct {
		arg1 alerts.Policy
	}
	createPolicyReturns struct {
		result1 *alerts.Policy
		result2 error
	}
	createPolicyReturnsOnCall map[int]struct {
		result1 *alerts.Policy
		result2 error
	}
	DeleteConditionStub        func(int) (*alerts.Condition, error)
	deleteConditionMutex       sync.RWMutex
	deleteConditionArgsForCall []struct {
		arg1 int
	}
	deleteConditionReturns struct {
		result1 *alerts.Condition
		result2 error
	}
	deleteConditionReturnsOnCall map[int]struct {
		result1 *alerts.Condition
		result2 error
	}
	DeleteNrqlConditionStub        func(int) (*alerts.NrqlCondition, error)
	deleteNrqlConditionMutex       sync.RWMutex
	deleteNrqlConditionArgsForCall []struct {
		arg1 int
	}
	deleteNrqlConditionReturns struct {
		result1 *alerts.NrqlCondition
		result2 error
	}
	deleteNrqlConditionReturnsOnCall map[int]struct {
		result1 *alerts.NrqlCondition
		result2 error
	}
	DeletePolicyStub        func(int) (*alerts.Policy, error)
	deletePolicyMutex       sync.RWMutex
	deletePolicyArgsForCall []struct {
		arg1 int
	}
	deletePolicyReturns struct {
		result1 *alerts.Policy
		result2 error
	}
	deletePolicyReturnsOnCall map[int]struct {
		result1 *alerts.Policy
		result2 error
	}
	GetPolicyStub        func(int) (*alerts.Policy, error)
	getPolicyMutex       sync.RWMutex
	getPolicyArgsForCall []struct {
		arg1 int
	}
	getPolicyReturns struct {
		result1 *alerts.Policy
		result2 error
	}
	getPolicyReturnsOnCall map[int]struct {
		result1 *alerts.Policy
		result2 error
	}
	ListConditionsStub        func(int) ([]*alerts.Condition, error)
	listConditionsMutex       sync.RWMutex
	listConditionsArgsForCall []struct {
		arg1 int
	}
	listConditionsReturns struct {
		result1 []*alerts.Condition
		result2 error
	}
	listConditionsReturnsOnCall map[int]struct {
		result1 []*alerts.Condition
		result2 error
	}
	ListNrqlConditionsStub        func(int) ([]*alerts.NrqlCondition, error)
	listNrqlConditionsMutex       sync.RWMutex
	listNrqlConditionsArgsForCall []struct {
		arg1 int
	}
	listNrqlConditionsReturns struct {
		result1 []*alerts.NrqlCondition
		result2 error
	}
	listNrqlConditionsReturnsOnCall map[int]struct {
		result1 []*alerts.NrqlCondition
		result2 error
	}
	ListPoliciesStub        func(*alerts.ListPoliciesParams) ([]alerts.Policy, error)
	listPoliciesMutex       sync.RWMutex
	listPoliciesArgsForCall []struct {
		arg1 *alerts.ListPoliciesParams
	}
	listPoliciesReturns struct {
		result1 []alerts.Policy
		result2 error
	}
	listPoliciesReturnsOnCall map[int]struct {
		result1 []alerts.Policy
		result2 error
	}
	UpdateConditionStub        func(alerts.Condition) (*alerts.Condition, error)
	updateConditionMutex       sync.RWMutex
	updateConditionArgsForCall []struct {
		arg1 alerts.Condition
	}
	updateConditionReturns struct {
		result1 *alerts.Condition
		result2 error
	}
	updateConditionReturnsOnCall map[int]struct {
		result1 *alerts.Condition
		result2 error
	}
	UpdateNrqlConditionStub        func(alerts.NrqlCondition) (*alerts.NrqlCondition, error)
	updateNrqlConditionMutex       sync.RWMutex
	updateNrqlConditionArgsForCall []struct {
		arg1 alerts.NrqlCondition
	}
	updateNrqlConditionReturns struct {
		result1 *alerts.NrqlCondition
		result2 error
	}
	updateNrqlConditionReturnsOnCall map[int]struct {
		result1 *alerts.NrqlCondition
		result2 error
	}
	UpdatePolicyStub        func(alerts.Policy) (*alerts.Policy, error)
	updatePolicyMutex       sync.RWMutex
	updatePolicyArgsForCall []struct {
		arg1 alerts.Policy
	}
	updatePolicyReturns struct {
		result1 *alerts.Policy
		result2 error
	}
	updatePolicyReturnsOnCall map[int]struct {
		result1 *alerts.Policy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNewRelicAlertsClient) CreateCondition(arg1 int, arg2 alerts.Condition) (*alerts.Condition, error) {
	fake.createConditionMutex.Lock()
	ret, specificReturn := fake.createConditionReturnsOnCall[len(fake.createConditionArgsForCall)]
	fake.createConditionArgsForCall = append(fake.createConditionArgsForCall, struct {
		arg1 int
		arg2 alerts.Condition
	}{arg1, arg2})
	fake.recordInvocation("CreateCondition", []interface{}{arg1, arg2})
	fake.createConditionMutex.Unlock()
	if fake.CreateConditionStub != nil {
		return fake.CreateConditionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createConditionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) CreateConditionCallCount() int {
	fake.createConditionMutex.RLock()
	defer fake.createConditionMutex.RUnlock()
	return len(fake.createConditionArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) CreateConditionCalls(stub func(int, alerts.Condition) (*alerts.Condition, error)) {
	fake.createConditionMutex.Lock()
	defer fake.createConditionMutex.Unlock()
	fake.CreateConditionStub = stub
}

func (fake *FakeNewRelicAlertsClient) CreateConditionArgsForCall(i int) (int, alerts.Condition) {
	fake.createConditionMutex.RLock()
	defer fake.createConditionMutex.RUnlock()
	argsForCall := fake.createConditionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNewRelicAlertsClient) CreateConditionReturns(result1 *alerts.Condition, result2 error) {
	fake.createConditionMutex.Lock()
	defer fake.createConditionMutex.Unlock()
	fake.CreateConditionStub = nil
	fake.createConditionReturns = struct {
		result1 *alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) CreateConditionReturnsOnCall(i int, result1 *alerts.Condition, result2 error) {
	fake.createConditionMutex.Lock()
	defer fake.createConditionMutex.Unlock()
	fake.CreateConditionStub = nil
	if fake.createConditionReturnsOnCall == nil {
		fake.createConditionReturnsOnCall = make(map[int]struct {
			result1 *alerts.Condition
			result2 error
		})
	}
	fake.createConditionReturnsOnCall[i] = struct {
		result1 *alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) CreateNrqlCondition(arg1 int, arg2 alerts.NrqlCondition) (*alerts.NrqlCondition, error) {
	fake.createNrqlConditionMutex.Lock()
	ret, specificReturn := fake.createNrqlConditionReturnsOnCall[len(fake.createNrqlConditionArgsForCall)]
	fake.createNrqlConditionArgsForCall = append(fake.createNrqlConditionArgsForCall, struct {
		arg1 int
		arg2 alerts.NrqlCondition
	}{arg1, arg2})
	fake.recordInvocation("CreateNrqlCondition", []interface{}{arg1, arg2})
	fake.createNrqlConditionMutex.Unlock()
	if fake.CreateNrqlConditionStub != nil {
		return fake.CreateNrqlConditionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createNrqlConditionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) CreateNrqlConditionCallCount() int {
	fake.createNrqlConditionMutex.RLock()
	defer fake.createNrqlConditionMutex.RUnlock()
	return len(fake.createNrqlConditionArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) CreateNrqlConditionCalls(stub func(int, alerts.NrqlCondition) (*alerts.NrqlCondition, error)) {
	fake.createNrqlConditionMutex.Lock()
	defer fake.createNrqlConditionMutex.Unlock()
	fake.CreateNrqlConditionStub = stub
}

func (fake *FakeNewRelicAlertsClient) CreateNrqlConditionArgsForCall(i int) (int, alerts.NrqlCondition) {
	fake.createNrqlConditionMutex.RLock()
	defer fake.createNrqlConditionMutex.RUnlock()
	argsForCall := fake.createNrqlConditionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNewRelicAlertsClient) CreateNrqlConditionReturns(result1 *alerts.NrqlCondition, result2 error) {
	fake.createNrqlConditionMutex.Lock()
	defer fake.createNrqlConditionMutex.Unlock()
	fake.CreateNrqlConditionStub = nil
	fake.createNrqlConditionReturns = struct {
		result1 *alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) CreateNrqlConditionReturnsOnCall(i int, result1 *alerts.NrqlCondition, result2 error) {
	fake.createNrqlConditionMutex.Lock()
	defer fake.createNrqlConditionMutex.Unlock()
	fake.CreateNrqlConditionStub = nil
	if fake.createNrqlConditionReturnsOnCall == nil {
		fake.createNrqlConditionReturnsOnCall = make(map[int]struct {
			result1 *alerts.NrqlCondition
			result2 error
		})
	}
	fake.createNrqlConditionReturnsOnCall[i] = struct {
		result1 *alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) CreatePolicy(arg1 alerts.Policy) (*alerts.Policy, error) {
	fake.createPolicyMutex.Lock()
	ret, specificReturn := fake.createPolicyReturnsOnCall[len(fake.createPolicyArgsForCall)]
	fake.createPolicyArgsForCall = append(fake.createPolicyArgsForCall, struct {
		arg1 alerts.Policy
	}{arg1})
	fake.recordInvocation("CreatePolicy", []interface{}{arg1})
	fake.createPolicyMutex.Unlock()
	if fake.CreatePolicyStub != nil {
		return fake.CreatePolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) CreatePolicyCallCount() int {
	fake.createPolicyMutex.RLock()
	defer fake.createPolicyMutex.RUnlock()
	return len(fake.createPolicyArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) CreatePolicyCalls(stub func(alerts.Policy) (*alerts.Policy, error)) {
	fake.createPolicyMutex.Lock()
	defer fake.createPolicyMutex.Unlock()
	fake.CreatePolicyStub = stub
}

func (fake *FakeNewRelicAlertsClient) CreatePolicyArgsForCall(i int) alerts.Policy {
	fake.createPolicyMutex.RLock()
	defer fake.createPolicyMutex.RUnlock()
	argsForCall := fake.createPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) CreatePolicyReturns(result1 *alerts.Policy, result2 error) {
	fake.createPolicyMutex.Lock()
	defer fake.createPolicyMutex.Unlock()
	fake.CreatePolicyStub = nil
	fake.createPolicyReturns = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) CreatePolicyReturnsOnCall(i int, result1 *alerts.Policy, result2 error) {
	fake.createPolicyMutex.Lock()
	defer fake.createPolicyMutex.Unlock()
	fake.CreatePolicyStub = nil
	if fake.createPolicyReturnsOnCall == nil {
		fake.createPolicyReturnsOnCall = make(map[int]struct {
			result1 *alerts.Policy
			result2 error
		})
	}
	fake.createPolicyReturnsOnCall[i] = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) DeleteCondition(arg1 int) (*alerts.Condition, error) {
	fake.deleteConditionMutex.Lock()
	ret, specificReturn := fake.deleteConditionReturnsOnCall[len(fake.deleteConditionArgsForCall)]
	fake.deleteConditionArgsForCall = append(fake.deleteConditionArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("DeleteCondition", []interface{}{arg1})
	fake.deleteConditionMutex.Unlock()
	if fake.DeleteConditionStub != nil {
		return fake.DeleteConditionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteConditionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) DeleteConditionCallCount() int {
	fake.deleteConditionMutex.RLock()
	defer fake.deleteConditionMutex.RUnlock()
	return len(fake.deleteConditionArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) DeleteConditionCalls(stub func(int) (*alerts.Condition, error)) {
	fake.deleteConditionMutex.Lock()
	defer fake.deleteConditionMutex.Unlock()
	fake.DeleteConditionStub = stub
}

func (fake *FakeNewRelicAlertsClient) DeleteConditionArgsForCall(i int) int {
	fake.deleteConditionMutex.RLock()
	defer fake.deleteConditionMutex.RUnlock()
	argsForCall := fake.deleteConditionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) DeleteConditionReturns(result1 *alerts.Condition, result2 error) {
	fake.deleteConditionMutex.Lock()
	defer fake.deleteConditionMutex.Unlock()
	fake.DeleteConditionStub = nil
	fake.deleteConditionReturns = struct {
		result1 *alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) DeleteConditionReturnsOnCall(i int, result1 *alerts.Condition, result2 error) {
	fake.deleteConditionMutex.Lock()
	defer fake.deleteConditionMutex.Unlock()
	fake.DeleteConditionStub = nil
	if fake.deleteConditionReturnsOnCall == nil {
		fake.deleteConditionReturnsOnCall = make(map[int]struct {
			result1 *alerts.Condition
			result2 error
		})
	}
	fake.deleteConditionReturnsOnCall[i] = struct {
		result1 *alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) DeleteNrqlCondition(arg1 int) (*alerts.NrqlCondition, error) {
	fake.deleteNrqlConditionMutex.Lock()
	ret, specificReturn := fake.deleteNrqlConditionReturnsOnCall[len(fake.deleteNrqlConditionArgsForCall)]
	fake.deleteNrqlConditionArgsForCall = append(fake.deleteNrqlConditionArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("DeleteNrqlCondition", []interface{}{arg1})
	fake.deleteNrqlConditionMutex.Unlock()
	if fake.DeleteNrqlConditionStub != nil {
		return fake.DeleteNrqlConditionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteNrqlConditionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) DeleteNrqlConditionCallCount() int {
	fake.deleteNrqlConditionMutex.RLock()
	defer fake.deleteNrqlConditionMutex.RUnlock()
	return len(fake.deleteNrqlConditionArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) DeleteNrqlConditionCalls(stub func(int) (*alerts.NrqlCondition, error)) {
	fake.deleteNrqlConditionMutex.Lock()
	defer fake.deleteNrqlConditionMutex.Unlock()
	fake.DeleteNrqlConditionStub = stub
}

func (fake *FakeNewRelicAlertsClient) DeleteNrqlConditionArgsForCall(i int) int {
	fake.deleteNrqlConditionMutex.RLock()
	defer fake.deleteNrqlConditionMutex.RUnlock()
	argsForCall := fake.deleteNrqlConditionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) DeleteNrqlConditionReturns(result1 *alerts.NrqlCondition, result2 error) {
	fake.deleteNrqlConditionMutex.Lock()
	defer fake.deleteNrqlConditionMutex.Unlock()
	fake.DeleteNrqlConditionStub = nil
	fake.deleteNrqlConditionReturns = struct {
		result1 *alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) DeleteNrqlConditionReturnsOnCall(i int, result1 *alerts.NrqlCondition, result2 error) {
	fake.deleteNrqlConditionMutex.Lock()
	defer fake.deleteNrqlConditionMutex.Unlock()
	fake.DeleteNrqlConditionStub = nil
	if fake.deleteNrqlConditionReturnsOnCall == nil {
		fake.deleteNrqlConditionReturnsOnCall = make(map[int]struct {
			result1 *alerts.NrqlCondition
			result2 error
		})
	}
	fake.deleteNrqlConditionReturnsOnCall[i] = struct {
		result1 *alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) DeletePolicy(arg1 int) (*alerts.Policy, error) {
	fake.deletePolicyMutex.Lock()
	ret, specificReturn := fake.deletePolicyReturnsOnCall[len(fake.deletePolicyArgsForCall)]
	fake.deletePolicyArgsForCall = append(fake.deletePolicyArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("DeletePolicy", []interface{}{arg1})
	fake.deletePolicyMutex.Unlock()
	if fake.DeletePolicyStub != nil {
		return fake.DeletePolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deletePolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) DeletePolicyCallCount() int {
	fake.deletePolicyMutex.RLock()
	defer fake.deletePolicyMutex.RUnlock()
	return len(fake.deletePolicyArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) DeletePolicyCalls(stub func(int) (*alerts.Policy, error)) {
	fake.deletePolicyMutex.Lock()
	defer fake.deletePolicyMutex.Unlock()
	fake.DeletePolicyStub = stub
}

func (fake *FakeNewRelicAlertsClient) DeletePolicyArgsForCall(i int) int {
	fake.deletePolicyMutex.RLock()
	defer fake.deletePolicyMutex.RUnlock()
	argsForCall := fake.deletePolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) DeletePolicyReturns(result1 *alerts.Policy, result2 error) {
	fake.deletePolicyMutex.Lock()
	defer fake.deletePolicyMutex.Unlock()
	fake.DeletePolicyStub = nil
	fake.deletePolicyReturns = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) DeletePolicyReturnsOnCall(i int, result1 *alerts.Policy, result2 error) {
	fake.deletePolicyMutex.Lock()
	defer fake.deletePolicyMutex.Unlock()
	fake.DeletePolicyStub = nil
	if fake.deletePolicyReturnsOnCall == nil {
		fake.deletePolicyReturnsOnCall = make(map[int]struct {
			result1 *alerts.Policy
			result2 error
		})
	}
	fake.deletePolicyReturnsOnCall[i] = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) GetPolicy(arg1 int) (*alerts.Policy, error) {
	fake.getPolicyMutex.Lock()
	ret, specificReturn := fake.getPolicyReturnsOnCall[len(fake.getPolicyArgsForCall)]
	fake.getPolicyArgsForCall = append(fake.getPolicyArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("GetPolicy", []interface{}{arg1})
	fake.getPolicyMutex.Unlock()
	if fake.GetPolicyStub != nil {
		return fake.GetPolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) GetPolicyCallCount() int {
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	return len(fake.getPolicyArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) GetPolicyCalls(stub func(int) (*alerts.Policy, error)) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = stub
}

func (fake *FakeNewRelicAlertsClient) GetPolicyArgsForCall(i int) int {
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	argsForCall := fake.getPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) GetPolicyReturns(result1 *alerts.Policy, result2 error) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = nil
	fake.getPolicyReturns = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) GetPolicyReturnsOnCall(i int, result1 *alerts.Policy, result2 error) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = nil
	if fake.getPolicyReturnsOnCall == nil {
		fake.getPolicyReturnsOnCall = make(map[int]struct {
			result1 *alerts.Policy
			result2 error
		})
	}
	fake.getPolicyReturnsOnCall[i] = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) ListConditions(arg1 int) ([]*alerts.Condition, error) {
	fake.listConditionsMutex.Lock()
	ret, specificReturn := fake.listConditionsReturnsOnCall[len(fake.listConditionsArgsForCall)]
	fake.listConditionsArgsForCall = append(fake.listConditionsArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ListConditions", []interface{}{arg1})
	fake.listConditionsMutex.Unlock()
	if fake.ListConditionsStub != nil {
		return fake.ListConditionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listConditionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) ListConditionsCallCount() int {
	fake.listConditionsMutex.RLock()
	defer fake.listConditionsMutex.RUnlock()
	return len(fake.listConditionsArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) ListConditionsCalls(stub func(int) ([]*alerts.Condition, error)) {
	fake.listConditionsMutex.Lock()
	defer fake.listConditionsMutex.Unlock()
	fake.ListConditionsStub = stub
}

func (fake *FakeNewRelicAlertsClient) ListConditionsArgsForCall(i int) int {
	fake.listConditionsMutex.RLock()
	defer fake.listConditionsMutex.RUnlock()
	argsForCall := fake.listConditionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) ListConditionsReturns(result1 []*alerts.Condition, result2 error) {
	fake.listConditionsMutex.Lock()
	defer fake.listConditionsMutex.Unlock()
	fake.ListConditionsStub = nil
	fake.listConditionsReturns = struct {
		result1 []*alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) ListConditionsReturnsOnCall(i int, result1 []*alerts.Condition, result2 error) {
	fake.listConditionsMutex.Lock()
	defer fake.listConditionsMutex.Unlock()
	fake.ListConditionsStub = nil
	if fake.listConditionsReturnsOnCall == nil {
		fake.listConditionsReturnsOnCall = make(map[int]struct {
			result1 []*alerts.Condition
			result2 error
		})
	}
	fake.listConditionsReturnsOnCall[i] = struct {
		result1 []*alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) ListNrqlConditions(arg1 int) ([]*alerts.NrqlCondition, error) {
	fake.listNrqlConditionsMutex.Lock()
	ret, specificReturn := fake.listNrqlConditionsReturnsOnCall[len(fake.listNrqlConditionsArgsForCall)]
	fake.listNrqlConditionsArgsForCall = append(fake.listNrqlConditionsArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ListNrqlConditions", []interface{}{arg1})
	fake.listNrqlConditionsMutex.Unlock()
	if fake.ListNrqlConditionsStub != nil {
		return fake.ListNrqlConditionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listNrqlConditionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) ListNrqlConditionsCallCount() int {
	fake.listNrqlConditionsMutex.RLock()
	defer fake.listNrqlConditionsMutex.RUnlock()
	return len(fake.listNrqlConditionsArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) ListNrqlConditionsCalls(stub func(int) ([]*alerts.NrqlCondition, error)) {
	fake.listNrqlConditionsMutex.Lock()
	defer fake.listNrqlConditionsMutex.Unlock()
	fake.ListNrqlConditionsStub = stub
}

func (fake *FakeNewRelicAlertsClient) ListNrqlConditionsArgsForCall(i int) int {
	fake.listNrqlConditionsMutex.RLock()
	defer fake.listNrqlConditionsMutex.RUnlock()
	argsForCall := fake.listNrqlConditionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) ListNrqlConditionsReturns(result1 []*alerts.NrqlCondition, result2 error) {
	fake.listNrqlConditionsMutex.Lock()
	defer fake.listNrqlConditionsMutex.Unlock()
	fake.ListNrqlConditionsStub = nil
	fake.listNrqlConditionsReturns = struct {
		result1 []*alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) ListNrqlConditionsReturnsOnCall(i int, result1 []*alerts.NrqlCondition, result2 error) {
	fake.listNrqlConditionsMutex.Lock()
	defer fake.listNrqlConditionsMutex.Unlock()
	fake.ListNrqlConditionsStub = nil
	if fake.listNrqlConditionsReturnsOnCall == nil {
		fake.listNrqlConditionsReturnsOnCall = make(map[int]struct {
			result1 []*alerts.NrqlCondition
			result2 error
		})
	}
	fake.listNrqlConditionsReturnsOnCall[i] = struct {
		result1 []*alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) ListPolicies(arg1 *alerts.ListPoliciesParams) ([]alerts.Policy, error) {
	fake.listPoliciesMutex.Lock()
	ret, specificReturn := fake.listPoliciesReturnsOnCall[len(fake.listPoliciesArgsForCall)]
	fake.listPoliciesArgsForCall = append(fake.listPoliciesArgsForCall, struct {
		arg1 *alerts.ListPoliciesParams
	}{arg1})
	fake.recordInvocation("ListPolicies", []interface{}{arg1})
	fake.listPoliciesMutex.Unlock()
	if fake.ListPoliciesStub != nil {
		return fake.ListPoliciesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listPoliciesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) ListPoliciesCallCount() int {
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	return len(fake.listPoliciesArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) ListPoliciesCalls(stub func(*alerts.ListPoliciesParams) ([]alerts.Policy, error)) {
	fake.listPoliciesMutex.Lock()
	defer fake.listPoliciesMutex.Unlock()
	fake.ListPoliciesStub = stub
}

func (fake *FakeNewRelicAlertsClient) ListPoliciesArgsForCall(i int) *alerts.ListPoliciesParams {
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	argsForCall := fake.listPoliciesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) ListPoliciesReturns(result1 []alerts.Policy, result2 error) {
	fake.listPoliciesMutex.Lock()
	defer fake.listPoliciesMutex.Unlock()
	fake.ListPoliciesStub = nil
	fake.listPoliciesReturns = struct {
		result1 []alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) ListPoliciesReturnsOnCall(i int, result1 []alerts.Policy, result2 error) {
	fake.listPoliciesMutex.Lock()
	defer fake.listPoliciesMutex.Unlock()
	fake.ListPoliciesStub = nil
	if fake.listPoliciesReturnsOnCall == nil {
		fake.listPoliciesReturnsOnCall = make(map[int]struct {
			result1 []alerts.Policy
			result2 error
		})
	}
	fake.listPoliciesReturnsOnCall[i] = struct {
		result1 []alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) UpdateCondition(arg1 alerts.Condition) (*alerts.Condition, error) {
	fake.updateConditionMutex.Lock()
	ret, specificReturn := fake.updateConditionReturnsOnCall[len(fake.updateConditionArgsForCall)]
	fake.updateConditionArgsForCall = append(fake.updateConditionArgsForCall, struct {
		arg1 alerts.Condition
	}{arg1})
	fake.recordInvocation("UpdateCondition", []interface{}{arg1})
	fake.updateConditionMutex.Unlock()
	if fake.UpdateConditionStub != nil {
		return fake.UpdateConditionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateConditionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) UpdateConditionCallCount() int {
	fake.updateConditionMutex.RLock()
	defer fake.updateConditionMutex.RUnlock()
	return len(fake.updateConditionArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) UpdateConditionCalls(stub func(alerts.Condition) (*alerts.Condition, error)) {
	fake.updateConditionMutex.Lock()
	defer fake.updateConditionMutex.Unlock()
	fake.UpdateConditionStub = stub
}

func (fake *FakeNewRelicAlertsClient) UpdateConditionArgsForCall(i int) alerts.Condition {
	fake.updateConditionMutex.RLock()
	defer fake.updateConditionMutex.RUnlock()
	argsForCall := fake.updateConditionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) UpdateConditionReturns(result1 *alerts.Condition, result2 error) {
	fake.updateConditionMutex.Lock()
	defer fake.updateConditionMutex.Unlock()
	fake.UpdateConditionStub = nil
	fake.updateConditionReturns = struct {
		result1 *alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) UpdateConditionReturnsOnCall(i int, result1 *alerts.Condition, result2 error) {
	fake.updateConditionMutex.Lock()
	defer fake.updateConditionMutex.Unlock()
	fake.UpdateConditionStub = nil
	if fake.updateConditionReturnsOnCall == nil {
		fake.updateConditionReturnsOnCall = make(map[int]struct {
			result1 *alerts.Condition
			result2 error
		})
	}
	fake.updateConditionReturnsOnCall[i] = struct {
		result1 *alerts.Condition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) UpdateNrqlCondition(arg1 alerts.NrqlCondition) (*alerts.NrqlCondition, error) {
	fake.updateNrqlConditionMutex.Lock()
	ret, specificReturn := fake.updateNrqlConditionReturnsOnCall[len(fake.updateNrqlConditionArgsForCall)]
	fake.updateNrqlConditionArgsForCall = append(fake.updateNrqlConditionArgsForCall, struct {
		arg1 alerts.NrqlCondition
	}{arg1})
	fake.recordInvocation("UpdateNrqlCondition", []interface{}{arg1})
	fake.updateNrqlConditionMutex.Unlock()
	if fake.UpdateNrqlConditionStub != nil {
		return fake.UpdateNrqlConditionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateNrqlConditionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) UpdateNrqlConditionCallCount() int {
	fake.updateNrqlConditionMutex.RLock()
	defer fake.updateNrqlConditionMutex.RUnlock()
	return len(fake.updateNrqlConditionArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) UpdateNrqlConditionCalls(stub func(alerts.NrqlCondition) (*alerts.NrqlCondition, error)) {
	fake.updateNrqlConditionMutex.Lock()
	defer fake.updateNrqlConditionMutex.Unlock()
	fake.UpdateNrqlConditionStub = stub
}

func (fake *FakeNewRelicAlertsClient) UpdateNrqlConditionArgsForCall(i int) alerts.NrqlCondition {
	fake.updateNrqlConditionMutex.RLock()
	defer fake.updateNrqlConditionMutex.RUnlock()
	argsForCall := fake.updateNrqlConditionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) UpdateNrqlConditionReturns(result1 *alerts.NrqlCondition, result2 error) {
	fake.updateNrqlConditionMutex.Lock()
	defer fake.updateNrqlConditionMutex.Unlock()
	fake.UpdateNrqlConditionStub = nil
	fake.updateNrqlConditionReturns = struct {
		result1 *alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) UpdateNrqlConditionReturnsOnCall(i int, result1 *alerts.NrqlCondition, result2 error) {
	fake.updateNrqlConditionMutex.Lock()
	defer fake.updateNrqlConditionMutex.Unlock()
	fake.UpdateNrqlConditionStub = nil
	if fake.updateNrqlConditionReturnsOnCall == nil {
		fake.updateNrqlConditionReturnsOnCall = make(map[int]struct {
			result1 *alerts.NrqlCondition
			result2 error
		})
	}
	fake.updateNrqlConditionReturnsOnCall[i] = struct {
		result1 *alerts.NrqlCondition
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) UpdatePolicy(arg1 alerts.Policy) (*alerts.Policy, error) {
	fake.updatePolicyMutex.Lock()
	ret, specificReturn := fake.updatePolicyReturnsOnCall[len(fake.updatePolicyArgsForCall)]
	fake.updatePolicyArgsForCall = append(fake.updatePolicyArgsForCall, struct {
		arg1 alerts.Policy
	}{arg1})
	fake.recordInvocation("UpdatePolicy", []interface{}{arg1})
	fake.updatePolicyMutex.Unlock()
	if fake.UpdatePolicyStub != nil {
		return fake.UpdatePolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updatePolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNewRelicAlertsClient) UpdatePolicyCallCount() int {
	fake.updatePolicyMutex.RLock()
	defer fake.updatePolicyMutex.RUnlock()
	return len(fake.updatePolicyArgsForCall)
}

func (fake *FakeNewRelicAlertsClient) UpdatePolicyCalls(stub func(alerts.Policy) (*alerts.Policy, error)) {
	fake.updatePolicyMutex.Lock()
	defer fake.updatePolicyMutex.Unlock()
	fake.UpdatePolicyStub = stub
}

func (fake *FakeNewRelicAlertsClient) UpdatePolicyArgsForCall(i int) alerts.Policy {
	fake.updatePolicyMutex.RLock()
	defer fake.updatePolicyMutex.RUnlock()
	argsForCall := fake.updatePolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNewRelicAlertsClient) UpdatePolicyReturns(result1 *alerts.Policy, result2 error) {
	fake.updatePolicyMutex.Lock()
	defer fake.updatePolicyMutex.Unlock()
	fake.UpdatePolicyStub = nil
	fake.updatePolicyReturns = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) UpdatePolicyReturnsOnCall(i int, result1 *alerts.Policy, result2 error) {
	fake.updatePolicyMutex.Lock()
	defer fake.updatePolicyMutex.Unlock()
	fake.UpdatePolicyStub = nil
	if fake.updatePolicyReturnsOnCall == nil {
		fake.updatePolicyReturnsOnCall = make(map[int]struct {
			result1 *alerts.Policy
			result2 error
		})
	}
	fake.updatePolicyReturnsOnCall[i] = struct {
		result1 *alerts.Policy
		result2 error
	}{result1, result2}
}

func (fake *FakeNewRelicAlertsClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createConditionMutex.RLock()
	defer fake.createConditionMutex.RUnlock()
	fake.createNrqlConditionMutex.RLock()
	defer fake.createNrqlConditionMutex.RUnlock()
	fake.createPolicyMutex.RLock()
	defer fake.createPolicyMutex.RUnlock()
	fake.deleteConditionMutex.RLock()
	defer fake.deleteConditionMutex.RUnlock()
	fake.deleteNrqlConditionMutex.RLock()
	defer fake.deleteNrqlConditionMutex.RUnlock()
	fake.deletePolicyMutex.RLock()
	defer fake.deletePolicyMutex.RUnlock()
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	fake.listConditionsMutex.RLock()
	defer fake.listConditionsMutex.RUnlock()
	fake.listNrqlConditionsMutex.RLock()
	defer fake.listNrqlConditionsMutex.RUnlock()
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	fake.updateConditionMutex.RLock()
	defer fake.updateConditionMutex.RUnlock()
	fake.updateNrqlConditionMutex.RLock()
	defer fake.updateNrqlConditionMutex.RUnlock()
	fake.updatePolicyMutex.RLock()
	defer fake.updatePolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNewRelicAlertsClient) recordInvocation(key string, args []interface{}) {
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

var _ interfaces.NewRelicAlertsClient = new(FakeNewRelicAlertsClient)
