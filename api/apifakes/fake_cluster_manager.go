// This file was generated by counterfeiter
package apifakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/switchboard/api"
)

type FakeClusterManager struct {
	AsJSONStub        func() api.ClusterJSON
	asJSONMutex       sync.RWMutex
	asJSONArgsForCall []struct{}
	asJSONReturns     struct {
		result1 api.ClusterJSON
	}
	EnableTrafficStub        func(string)
	enableTrafficMutex       sync.RWMutex
	enableTrafficArgsForCall []struct {
		arg1 string
	}
	DisableTrafficStub        func(string)
	disableTrafficMutex       sync.RWMutex
	disableTrafficArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClusterManager) AsJSON() api.ClusterJSON {
	fake.asJSONMutex.Lock()
	fake.asJSONArgsForCall = append(fake.asJSONArgsForCall, struct{}{})
	fake.recordInvocation("AsJSON", []interface{}{})
	fake.asJSONMutex.Unlock()
	if fake.AsJSONStub != nil {
		return fake.AsJSONStub()
	} else {
		return fake.asJSONReturns.result1
	}
}

func (fake *FakeClusterManager) AsJSONCallCount() int {
	fake.asJSONMutex.RLock()
	defer fake.asJSONMutex.RUnlock()
	return len(fake.asJSONArgsForCall)
}

func (fake *FakeClusterManager) AsJSONReturns(result1 api.ClusterJSON) {
	fake.AsJSONStub = nil
	fake.asJSONReturns = struct {
		result1 api.ClusterJSON
	}{result1}
}

func (fake *FakeClusterManager) EnableTraffic(arg1 string) {
	fake.enableTrafficMutex.Lock()
	fake.enableTrafficArgsForCall = append(fake.enableTrafficArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("EnableTraffic", []interface{}{arg1})
	fake.enableTrafficMutex.Unlock()
	if fake.EnableTrafficStub != nil {
		fake.EnableTrafficStub(arg1)
	}
}

func (fake *FakeClusterManager) EnableTrafficCallCount() int {
	fake.enableTrafficMutex.RLock()
	defer fake.enableTrafficMutex.RUnlock()
	return len(fake.enableTrafficArgsForCall)
}

func (fake *FakeClusterManager) EnableTrafficArgsForCall(i int) string {
	fake.enableTrafficMutex.RLock()
	defer fake.enableTrafficMutex.RUnlock()
	return fake.enableTrafficArgsForCall[i].arg1
}

func (fake *FakeClusterManager) DisableTraffic(arg1 string) {
	fake.disableTrafficMutex.Lock()
	fake.disableTrafficArgsForCall = append(fake.disableTrafficArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DisableTraffic", []interface{}{arg1})
	fake.disableTrafficMutex.Unlock()
	if fake.DisableTrafficStub != nil {
		fake.DisableTrafficStub(arg1)
	}
}

func (fake *FakeClusterManager) DisableTrafficCallCount() int {
	fake.disableTrafficMutex.RLock()
	defer fake.disableTrafficMutex.RUnlock()
	return len(fake.disableTrafficArgsForCall)
}

func (fake *FakeClusterManager) DisableTrafficArgsForCall(i int) string {
	fake.disableTrafficMutex.RLock()
	defer fake.disableTrafficMutex.RUnlock()
	return fake.disableTrafficArgsForCall[i].arg1
}

func (fake *FakeClusterManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.asJSONMutex.RLock()
	defer fake.asJSONMutex.RUnlock()
	fake.enableTrafficMutex.RLock()
	defer fake.enableTrafficMutex.RUnlock()
	fake.disableTrafficMutex.RLock()
	defer fake.disableTrafficMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClusterManager) recordInvocation(key string, args []interface{}) {
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

var _ api.ClusterManager = new(FakeClusterManager)
