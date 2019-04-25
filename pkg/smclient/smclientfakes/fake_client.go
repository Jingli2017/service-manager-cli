// Code generated by counterfeiter. DO NOT EDIT.
package smclientfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/Peripli/service-manager-cli/pkg/smclient"
	"github.com/Peripli/service-manager-cli/pkg/types"
)

type FakeClient struct {
	CallStub        func(string, string, io.Reader) (*http.Response, error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 io.Reader
	}
	callReturns struct {
		result1 *http.Response
		result2 error
	}
	callReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	DeleteBrokerStub        func(string) error
	deleteBrokerMutex       sync.RWMutex
	deleteBrokerArgsForCall []struct {
		arg1 string
	}
	deleteBrokerReturns struct {
		result1 error
	}
	deleteBrokerReturnsOnCall map[int]struct {
		result1 error
	}
	DeletePlatformStub        func(string) error
	deletePlatformMutex       sync.RWMutex
	deletePlatformArgsForCall []struct {
		arg1 string
	}
	deletePlatformReturns struct {
		result1 error
	}
	deletePlatformReturnsOnCall map[int]struct {
		result1 error
	}
	GetInfoStub        func() (*types.Info, error)
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct {
	}
	getInfoReturns struct {
		result1 *types.Info
		result2 error
	}
	getInfoReturnsOnCall map[int]struct {
		result1 *types.Info
		result2 error
	}
	ListBrokersStub        func() (*types.Brokers, error)
	listBrokersMutex       sync.RWMutex
	listBrokersArgsForCall []struct {
	}
	listBrokersReturns struct {
		result1 *types.Brokers
		result2 error
	}
	listBrokersReturnsOnCall map[int]struct {
		result1 *types.Brokers
		result2 error
	}
	ListBrokersWithQueryStub        func(string, string) (*types.Brokers, error)
	listBrokersWithQueryMutex       sync.RWMutex
	listBrokersWithQueryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listBrokersWithQueryReturns struct {
		result1 *types.Brokers
		result2 error
	}
	listBrokersWithQueryReturnsOnCall map[int]struct {
		result1 *types.Brokers
		result2 error
	}
	ListOfferingsStub        func() (*types.ServiceOfferings, error)
	listOfferingsMutex       sync.RWMutex
	listOfferingsArgsForCall []struct {
	}
	listOfferingsReturns struct {
		result1 *types.ServiceOfferings
		result2 error
	}
	listOfferingsReturnsOnCall map[int]struct {
		result1 *types.ServiceOfferings
		result2 error
	}
	ListOfferingsWithQueryStub        func(string, string) (*types.ServiceOfferings, error)
	listOfferingsWithQueryMutex       sync.RWMutex
	listOfferingsWithQueryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listOfferingsWithQueryReturns struct {
		result1 *types.ServiceOfferings
		result2 error
	}
	listOfferingsWithQueryReturnsOnCall map[int]struct {
		result1 *types.ServiceOfferings
		result2 error
	}
	ListPlatformsStub        func() (*types.Platforms, error)
	listPlatformsMutex       sync.RWMutex
	listPlatformsArgsForCall []struct {
	}
	listPlatformsReturns struct {
		result1 *types.Platforms
		result2 error
	}
	listPlatformsReturnsOnCall map[int]struct {
		result1 *types.Platforms
		result2 error
	}
	ListPlatformsWithQueryStub        func(string, string) (*types.Platforms, error)
	listPlatformsWithQueryMutex       sync.RWMutex
	listPlatformsWithQueryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listPlatformsWithQueryReturns struct {
		result1 *types.Platforms
		result2 error
	}
	listPlatformsWithQueryReturnsOnCall map[int]struct {
		result1 *types.Platforms
		result2 error
	}
	RegisterBrokerStub        func(*types.Broker) (*types.Broker, error)
	registerBrokerMutex       sync.RWMutex
	registerBrokerArgsForCall []struct {
		arg1 *types.Broker
	}
	registerBrokerReturns struct {
		result1 *types.Broker
		result2 error
	}
	registerBrokerReturnsOnCall map[int]struct {
		result1 *types.Broker
		result2 error
	}
	RegisterPlatformStub        func(*types.Platform) (*types.Platform, error)
	registerPlatformMutex       sync.RWMutex
	registerPlatformArgsForCall []struct {
		arg1 *types.Platform
	}
	registerPlatformReturns struct {
		result1 *types.Platform
		result2 error
	}
	registerPlatformReturnsOnCall map[int]struct {
		result1 *types.Platform
		result2 error
	}
	UpdateBrokerStub        func(string, *types.Broker) (*types.Broker, error)
	updateBrokerMutex       sync.RWMutex
	updateBrokerArgsForCall []struct {
		arg1 string
		arg2 *types.Broker
	}
	updateBrokerReturns struct {
		result1 *types.Broker
		result2 error
	}
	updateBrokerReturnsOnCall map[int]struct {
		result1 *types.Broker
		result2 error
	}
	UpdatePlatformStub        func(string, *types.Platform) (*types.Platform, error)
	updatePlatformMutex       sync.RWMutex
	updatePlatformArgsForCall []struct {
		arg1 string
		arg2 *types.Platform
	}
	updatePlatformReturns struct {
		result1 *types.Platform
		result2 error
	}
	updatePlatformReturnsOnCall map[int]struct {
		result1 *types.Platform
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Call(arg1 string, arg2 string, arg3 io.Reader) (*http.Response, error) {
	fake.callMutex.Lock()
	ret, specificReturn := fake.callReturnsOnCall[len(fake.callArgsForCall)]
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 io.Reader
	}{arg1, arg2, arg3})
	fake.recordInvocation("Call", []interface{}{arg1, arg2, arg3})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.callReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *FakeClient) CallCalls(stub func(string, string, io.Reader) (*http.Response, error)) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = stub
}

func (fake *FakeClient) CallArgsForCall(i int) (string, string, io.Reader) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	argsForCall := fake.callArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) CallReturns(result1 *http.Response, result2 error) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CallReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = nil
	if fake.callReturnsOnCall == nil {
		fake.callReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.callReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeleteBroker(arg1 string) error {
	fake.deleteBrokerMutex.Lock()
	ret, specificReturn := fake.deleteBrokerReturnsOnCall[len(fake.deleteBrokerArgsForCall)]
	fake.deleteBrokerArgsForCall = append(fake.deleteBrokerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteBroker", []interface{}{arg1})
	fake.deleteBrokerMutex.Unlock()
	if fake.DeleteBrokerStub != nil {
		return fake.DeleteBrokerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteBrokerReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DeleteBrokerCallCount() int {
	fake.deleteBrokerMutex.RLock()
	defer fake.deleteBrokerMutex.RUnlock()
	return len(fake.deleteBrokerArgsForCall)
}

func (fake *FakeClient) DeleteBrokerCalls(stub func(string) error) {
	fake.deleteBrokerMutex.Lock()
	defer fake.deleteBrokerMutex.Unlock()
	fake.DeleteBrokerStub = stub
}

func (fake *FakeClient) DeleteBrokerArgsForCall(i int) string {
	fake.deleteBrokerMutex.RLock()
	defer fake.deleteBrokerMutex.RUnlock()
	argsForCall := fake.deleteBrokerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DeleteBrokerReturns(result1 error) {
	fake.deleteBrokerMutex.Lock()
	defer fake.deleteBrokerMutex.Unlock()
	fake.DeleteBrokerStub = nil
	fake.deleteBrokerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteBrokerReturnsOnCall(i int, result1 error) {
	fake.deleteBrokerMutex.Lock()
	defer fake.deleteBrokerMutex.Unlock()
	fake.DeleteBrokerStub = nil
	if fake.deleteBrokerReturnsOnCall == nil {
		fake.deleteBrokerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteBrokerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeletePlatform(arg1 string) error {
	fake.deletePlatformMutex.Lock()
	ret, specificReturn := fake.deletePlatformReturnsOnCall[len(fake.deletePlatformArgsForCall)]
	fake.deletePlatformArgsForCall = append(fake.deletePlatformArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeletePlatform", []interface{}{arg1})
	fake.deletePlatformMutex.Unlock()
	if fake.DeletePlatformStub != nil {
		return fake.DeletePlatformStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deletePlatformReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DeletePlatformCallCount() int {
	fake.deletePlatformMutex.RLock()
	defer fake.deletePlatformMutex.RUnlock()
	return len(fake.deletePlatformArgsForCall)
}

func (fake *FakeClient) DeletePlatformCalls(stub func(string) error) {
	fake.deletePlatformMutex.Lock()
	defer fake.deletePlatformMutex.Unlock()
	fake.DeletePlatformStub = stub
}

func (fake *FakeClient) DeletePlatformArgsForCall(i int) string {
	fake.deletePlatformMutex.RLock()
	defer fake.deletePlatformMutex.RUnlock()
	argsForCall := fake.deletePlatformArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DeletePlatformReturns(result1 error) {
	fake.deletePlatformMutex.Lock()
	defer fake.deletePlatformMutex.Unlock()
	fake.DeletePlatformStub = nil
	fake.deletePlatformReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeletePlatformReturnsOnCall(i int, result1 error) {
	fake.deletePlatformMutex.Lock()
	defer fake.deletePlatformMutex.Unlock()
	fake.DeletePlatformStub = nil
	if fake.deletePlatformReturnsOnCall == nil {
		fake.deletePlatformReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePlatformReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetInfo() (*types.Info, error) {
	fake.getInfoMutex.Lock()
	ret, specificReturn := fake.getInfoReturnsOnCall[len(fake.getInfoArgsForCall)]
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct {
	}{})
	fake.recordInvocation("GetInfo", []interface{}{})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeClient) GetInfoCalls(stub func() (*types.Info, error)) {
	fake.getInfoMutex.Lock()
	defer fake.getInfoMutex.Unlock()
	fake.GetInfoStub = stub
}

func (fake *FakeClient) GetInfoReturns(result1 *types.Info, result2 error) {
	fake.getInfoMutex.Lock()
	defer fake.getInfoMutex.Unlock()
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 *types.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetInfoReturnsOnCall(i int, result1 *types.Info, result2 error) {
	fake.getInfoMutex.Lock()
	defer fake.getInfoMutex.Unlock()
	fake.GetInfoStub = nil
	if fake.getInfoReturnsOnCall == nil {
		fake.getInfoReturnsOnCall = make(map[int]struct {
			result1 *types.Info
			result2 error
		})
	}
	fake.getInfoReturnsOnCall[i] = struct {
		result1 *types.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListBrokers() (*types.Brokers, error) {
	fake.listBrokersMutex.Lock()
	ret, specificReturn := fake.listBrokersReturnsOnCall[len(fake.listBrokersArgsForCall)]
	fake.listBrokersArgsForCall = append(fake.listBrokersArgsForCall, struct {
	}{})
	fake.recordInvocation("ListBrokers", []interface{}{})
	fake.listBrokersMutex.Unlock()
	if fake.ListBrokersStub != nil {
		return fake.ListBrokersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listBrokersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListBrokersCallCount() int {
	fake.listBrokersMutex.RLock()
	defer fake.listBrokersMutex.RUnlock()
	return len(fake.listBrokersArgsForCall)
}

func (fake *FakeClient) ListBrokersCalls(stub func() (*types.Brokers, error)) {
	fake.listBrokersMutex.Lock()
	defer fake.listBrokersMutex.Unlock()
	fake.ListBrokersStub = stub
}

func (fake *FakeClient) ListBrokersReturns(result1 *types.Brokers, result2 error) {
	fake.listBrokersMutex.Lock()
	defer fake.listBrokersMutex.Unlock()
	fake.ListBrokersStub = nil
	fake.listBrokersReturns = struct {
		result1 *types.Brokers
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListBrokersReturnsOnCall(i int, result1 *types.Brokers, result2 error) {
	fake.listBrokersMutex.Lock()
	defer fake.listBrokersMutex.Unlock()
	fake.ListBrokersStub = nil
	if fake.listBrokersReturnsOnCall == nil {
		fake.listBrokersReturnsOnCall = make(map[int]struct {
			result1 *types.Brokers
			result2 error
		})
	}
	fake.listBrokersReturnsOnCall[i] = struct {
		result1 *types.Brokers
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListBrokersWithQuery(arg1 string, arg2 string) (*types.Brokers, error) {
	fake.listBrokersWithQueryMutex.Lock()
	ret, specificReturn := fake.listBrokersWithQueryReturnsOnCall[len(fake.listBrokersWithQueryArgsForCall)]
	fake.listBrokersWithQueryArgsForCall = append(fake.listBrokersWithQueryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListBrokersWithQuery", []interface{}{arg1, arg2})
	fake.listBrokersWithQueryMutex.Unlock()
	if fake.ListBrokersWithQueryStub != nil {
		return fake.ListBrokersWithQueryStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listBrokersWithQueryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListBrokersWithQueryCallCount() int {
	fake.listBrokersWithQueryMutex.RLock()
	defer fake.listBrokersWithQueryMutex.RUnlock()
	return len(fake.listBrokersWithQueryArgsForCall)
}

func (fake *FakeClient) ListBrokersWithQueryCalls(stub func(string, string) (*types.Brokers, error)) {
	fake.listBrokersWithQueryMutex.Lock()
	defer fake.listBrokersWithQueryMutex.Unlock()
	fake.ListBrokersWithQueryStub = stub
}

func (fake *FakeClient) ListBrokersWithQueryArgsForCall(i int) (string, string) {
	fake.listBrokersWithQueryMutex.RLock()
	defer fake.listBrokersWithQueryMutex.RUnlock()
	argsForCall := fake.listBrokersWithQueryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ListBrokersWithQueryReturns(result1 *types.Brokers, result2 error) {
	fake.listBrokersWithQueryMutex.Lock()
	defer fake.listBrokersWithQueryMutex.Unlock()
	fake.ListBrokersWithQueryStub = nil
	fake.listBrokersWithQueryReturns = struct {
		result1 *types.Brokers
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListBrokersWithQueryReturnsOnCall(i int, result1 *types.Brokers, result2 error) {
	fake.listBrokersWithQueryMutex.Lock()
	defer fake.listBrokersWithQueryMutex.Unlock()
	fake.ListBrokersWithQueryStub = nil
	if fake.listBrokersWithQueryReturnsOnCall == nil {
		fake.listBrokersWithQueryReturnsOnCall = make(map[int]struct {
			result1 *types.Brokers
			result2 error
		})
	}
	fake.listBrokersWithQueryReturnsOnCall[i] = struct {
		result1 *types.Brokers
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListOfferings() (*types.ServiceOfferings, error) {
	fake.listOfferingsMutex.Lock()
	ret, specificReturn := fake.listOfferingsReturnsOnCall[len(fake.listOfferingsArgsForCall)]
	fake.listOfferingsArgsForCall = append(fake.listOfferingsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListOfferings", []interface{}{})
	fake.listOfferingsMutex.Unlock()
	if fake.ListOfferingsStub != nil {
		return fake.ListOfferingsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOfferingsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListOfferingsCallCount() int {
	fake.listOfferingsMutex.RLock()
	defer fake.listOfferingsMutex.RUnlock()
	return len(fake.listOfferingsArgsForCall)
}

func (fake *FakeClient) ListOfferingsCalls(stub func() (*types.ServiceOfferings, error)) {
	fake.listOfferingsMutex.Lock()
	defer fake.listOfferingsMutex.Unlock()
	fake.ListOfferingsStub = stub
}

func (fake *FakeClient) ListOfferingsReturns(result1 *types.ServiceOfferings, result2 error) {
	fake.listOfferingsMutex.Lock()
	defer fake.listOfferingsMutex.Unlock()
	fake.ListOfferingsStub = nil
	fake.listOfferingsReturns = struct {
		result1 *types.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListOfferingsReturnsOnCall(i int, result1 *types.ServiceOfferings, result2 error) {
	fake.listOfferingsMutex.Lock()
	defer fake.listOfferingsMutex.Unlock()
	fake.ListOfferingsStub = nil
	if fake.listOfferingsReturnsOnCall == nil {
		fake.listOfferingsReturnsOnCall = make(map[int]struct {
			result1 *types.ServiceOfferings
			result2 error
		})
	}
	fake.listOfferingsReturnsOnCall[i] = struct {
		result1 *types.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListOfferingsWithQuery(arg1 string, arg2 string) (*types.ServiceOfferings, error) {
	fake.listOfferingsWithQueryMutex.Lock()
	ret, specificReturn := fake.listOfferingsWithQueryReturnsOnCall[len(fake.listOfferingsWithQueryArgsForCall)]
	fake.listOfferingsWithQueryArgsForCall = append(fake.listOfferingsWithQueryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListOfferingsWithQuery", []interface{}{arg1, arg2})
	fake.listOfferingsWithQueryMutex.Unlock()
	if fake.ListOfferingsWithQueryStub != nil {
		return fake.ListOfferingsWithQueryStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOfferingsWithQueryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListOfferingsWithQueryCallCount() int {
	fake.listOfferingsWithQueryMutex.RLock()
	defer fake.listOfferingsWithQueryMutex.RUnlock()
	return len(fake.listOfferingsWithQueryArgsForCall)
}

func (fake *FakeClient) ListOfferingsWithQueryCalls(stub func(string, string) (*types.ServiceOfferings, error)) {
	fake.listOfferingsWithQueryMutex.Lock()
	defer fake.listOfferingsWithQueryMutex.Unlock()
	fake.ListOfferingsWithQueryStub = stub
}

func (fake *FakeClient) ListOfferingsWithQueryArgsForCall(i int) (string, string) {
	fake.listOfferingsWithQueryMutex.RLock()
	defer fake.listOfferingsWithQueryMutex.RUnlock()
	argsForCall := fake.listOfferingsWithQueryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ListOfferingsWithQueryReturns(result1 *types.ServiceOfferings, result2 error) {
	fake.listOfferingsWithQueryMutex.Lock()
	defer fake.listOfferingsWithQueryMutex.Unlock()
	fake.ListOfferingsWithQueryStub = nil
	fake.listOfferingsWithQueryReturns = struct {
		result1 *types.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListOfferingsWithQueryReturnsOnCall(i int, result1 *types.ServiceOfferings, result2 error) {
	fake.listOfferingsWithQueryMutex.Lock()
	defer fake.listOfferingsWithQueryMutex.Unlock()
	fake.ListOfferingsWithQueryStub = nil
	if fake.listOfferingsWithQueryReturnsOnCall == nil {
		fake.listOfferingsWithQueryReturnsOnCall = make(map[int]struct {
			result1 *types.ServiceOfferings
			result2 error
		})
	}
	fake.listOfferingsWithQueryReturnsOnCall[i] = struct {
		result1 *types.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListPlatforms() (*types.Platforms, error) {
	fake.listPlatformsMutex.Lock()
	ret, specificReturn := fake.listPlatformsReturnsOnCall[len(fake.listPlatformsArgsForCall)]
	fake.listPlatformsArgsForCall = append(fake.listPlatformsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListPlatforms", []interface{}{})
	fake.listPlatformsMutex.Unlock()
	if fake.ListPlatformsStub != nil {
		return fake.ListPlatformsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listPlatformsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListPlatformsCallCount() int {
	fake.listPlatformsMutex.RLock()
	defer fake.listPlatformsMutex.RUnlock()
	return len(fake.listPlatformsArgsForCall)
}

func (fake *FakeClient) ListPlatformsCalls(stub func() (*types.Platforms, error)) {
	fake.listPlatformsMutex.Lock()
	defer fake.listPlatformsMutex.Unlock()
	fake.ListPlatformsStub = stub
}

func (fake *FakeClient) ListPlatformsReturns(result1 *types.Platforms, result2 error) {
	fake.listPlatformsMutex.Lock()
	defer fake.listPlatformsMutex.Unlock()
	fake.ListPlatformsStub = nil
	fake.listPlatformsReturns = struct {
		result1 *types.Platforms
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListPlatformsReturnsOnCall(i int, result1 *types.Platforms, result2 error) {
	fake.listPlatformsMutex.Lock()
	defer fake.listPlatformsMutex.Unlock()
	fake.ListPlatformsStub = nil
	if fake.listPlatformsReturnsOnCall == nil {
		fake.listPlatformsReturnsOnCall = make(map[int]struct {
			result1 *types.Platforms
			result2 error
		})
	}
	fake.listPlatformsReturnsOnCall[i] = struct {
		result1 *types.Platforms
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListPlatformsWithQuery(arg1 string, arg2 string) (*types.Platforms, error) {
	fake.listPlatformsWithQueryMutex.Lock()
	ret, specificReturn := fake.listPlatformsWithQueryReturnsOnCall[len(fake.listPlatformsWithQueryArgsForCall)]
	fake.listPlatformsWithQueryArgsForCall = append(fake.listPlatformsWithQueryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListPlatformsWithQuery", []interface{}{arg1, arg2})
	fake.listPlatformsWithQueryMutex.Unlock()
	if fake.ListPlatformsWithQueryStub != nil {
		return fake.ListPlatformsWithQueryStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listPlatformsWithQueryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListPlatformsWithQueryCallCount() int {
	fake.listPlatformsWithQueryMutex.RLock()
	defer fake.listPlatformsWithQueryMutex.RUnlock()
	return len(fake.listPlatformsWithQueryArgsForCall)
}

func (fake *FakeClient) ListPlatformsWithQueryCalls(stub func(string, string) (*types.Platforms, error)) {
	fake.listPlatformsWithQueryMutex.Lock()
	defer fake.listPlatformsWithQueryMutex.Unlock()
	fake.ListPlatformsWithQueryStub = stub
}

func (fake *FakeClient) ListPlatformsWithQueryArgsForCall(i int) (string, string) {
	fake.listPlatformsWithQueryMutex.RLock()
	defer fake.listPlatformsWithQueryMutex.RUnlock()
	argsForCall := fake.listPlatformsWithQueryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ListPlatformsWithQueryReturns(result1 *types.Platforms, result2 error) {
	fake.listPlatformsWithQueryMutex.Lock()
	defer fake.listPlatformsWithQueryMutex.Unlock()
	fake.ListPlatformsWithQueryStub = nil
	fake.listPlatformsWithQueryReturns = struct {
		result1 *types.Platforms
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListPlatformsWithQueryReturnsOnCall(i int, result1 *types.Platforms, result2 error) {
	fake.listPlatformsWithQueryMutex.Lock()
	defer fake.listPlatformsWithQueryMutex.Unlock()
	fake.ListPlatformsWithQueryStub = nil
	if fake.listPlatformsWithQueryReturnsOnCall == nil {
		fake.listPlatformsWithQueryReturnsOnCall = make(map[int]struct {
			result1 *types.Platforms
			result2 error
		})
	}
	fake.listPlatformsWithQueryReturnsOnCall[i] = struct {
		result1 *types.Platforms
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RegisterBroker(arg1 *types.Broker) (*types.Broker, error) {
	fake.registerBrokerMutex.Lock()
	ret, specificReturn := fake.registerBrokerReturnsOnCall[len(fake.registerBrokerArgsForCall)]
	fake.registerBrokerArgsForCall = append(fake.registerBrokerArgsForCall, struct {
		arg1 *types.Broker
	}{arg1})
	fake.recordInvocation("RegisterBroker", []interface{}{arg1})
	fake.registerBrokerMutex.Unlock()
	if fake.RegisterBrokerStub != nil {
		return fake.RegisterBrokerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.registerBrokerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) RegisterBrokerCallCount() int {
	fake.registerBrokerMutex.RLock()
	defer fake.registerBrokerMutex.RUnlock()
	return len(fake.registerBrokerArgsForCall)
}

func (fake *FakeClient) RegisterBrokerCalls(stub func(*types.Broker) (*types.Broker, error)) {
	fake.registerBrokerMutex.Lock()
	defer fake.registerBrokerMutex.Unlock()
	fake.RegisterBrokerStub = stub
}

func (fake *FakeClient) RegisterBrokerArgsForCall(i int) *types.Broker {
	fake.registerBrokerMutex.RLock()
	defer fake.registerBrokerMutex.RUnlock()
	argsForCall := fake.registerBrokerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) RegisterBrokerReturns(result1 *types.Broker, result2 error) {
	fake.registerBrokerMutex.Lock()
	defer fake.registerBrokerMutex.Unlock()
	fake.RegisterBrokerStub = nil
	fake.registerBrokerReturns = struct {
		result1 *types.Broker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RegisterBrokerReturnsOnCall(i int, result1 *types.Broker, result2 error) {
	fake.registerBrokerMutex.Lock()
	defer fake.registerBrokerMutex.Unlock()
	fake.RegisterBrokerStub = nil
	if fake.registerBrokerReturnsOnCall == nil {
		fake.registerBrokerReturnsOnCall = make(map[int]struct {
			result1 *types.Broker
			result2 error
		})
	}
	fake.registerBrokerReturnsOnCall[i] = struct {
		result1 *types.Broker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RegisterPlatform(arg1 *types.Platform) (*types.Platform, error) {
	fake.registerPlatformMutex.Lock()
	ret, specificReturn := fake.registerPlatformReturnsOnCall[len(fake.registerPlatformArgsForCall)]
	fake.registerPlatformArgsForCall = append(fake.registerPlatformArgsForCall, struct {
		arg1 *types.Platform
	}{arg1})
	fake.recordInvocation("RegisterPlatform", []interface{}{arg1})
	fake.registerPlatformMutex.Unlock()
	if fake.RegisterPlatformStub != nil {
		return fake.RegisterPlatformStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.registerPlatformReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) RegisterPlatformCallCount() int {
	fake.registerPlatformMutex.RLock()
	defer fake.registerPlatformMutex.RUnlock()
	return len(fake.registerPlatformArgsForCall)
}

func (fake *FakeClient) RegisterPlatformCalls(stub func(*types.Platform) (*types.Platform, error)) {
	fake.registerPlatformMutex.Lock()
	defer fake.registerPlatformMutex.Unlock()
	fake.RegisterPlatformStub = stub
}

func (fake *FakeClient) RegisterPlatformArgsForCall(i int) *types.Platform {
	fake.registerPlatformMutex.RLock()
	defer fake.registerPlatformMutex.RUnlock()
	argsForCall := fake.registerPlatformArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) RegisterPlatformReturns(result1 *types.Platform, result2 error) {
	fake.registerPlatformMutex.Lock()
	defer fake.registerPlatformMutex.Unlock()
	fake.RegisterPlatformStub = nil
	fake.registerPlatformReturns = struct {
		result1 *types.Platform
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RegisterPlatformReturnsOnCall(i int, result1 *types.Platform, result2 error) {
	fake.registerPlatformMutex.Lock()
	defer fake.registerPlatformMutex.Unlock()
	fake.RegisterPlatformStub = nil
	if fake.registerPlatformReturnsOnCall == nil {
		fake.registerPlatformReturnsOnCall = make(map[int]struct {
			result1 *types.Platform
			result2 error
		})
	}
	fake.registerPlatformReturnsOnCall[i] = struct {
		result1 *types.Platform
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateBroker(arg1 string, arg2 *types.Broker) (*types.Broker, error) {
	fake.updateBrokerMutex.Lock()
	ret, specificReturn := fake.updateBrokerReturnsOnCall[len(fake.updateBrokerArgsForCall)]
	fake.updateBrokerArgsForCall = append(fake.updateBrokerArgsForCall, struct {
		arg1 string
		arg2 *types.Broker
	}{arg1, arg2})
	fake.recordInvocation("UpdateBroker", []interface{}{arg1, arg2})
	fake.updateBrokerMutex.Unlock()
	if fake.UpdateBrokerStub != nil {
		return fake.UpdateBrokerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateBrokerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) UpdateBrokerCallCount() int {
	fake.updateBrokerMutex.RLock()
	defer fake.updateBrokerMutex.RUnlock()
	return len(fake.updateBrokerArgsForCall)
}

func (fake *FakeClient) UpdateBrokerCalls(stub func(string, *types.Broker) (*types.Broker, error)) {
	fake.updateBrokerMutex.Lock()
	defer fake.updateBrokerMutex.Unlock()
	fake.UpdateBrokerStub = stub
}

func (fake *FakeClient) UpdateBrokerArgsForCall(i int) (string, *types.Broker) {
	fake.updateBrokerMutex.RLock()
	defer fake.updateBrokerMutex.RUnlock()
	argsForCall := fake.updateBrokerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) UpdateBrokerReturns(result1 *types.Broker, result2 error) {
	fake.updateBrokerMutex.Lock()
	defer fake.updateBrokerMutex.Unlock()
	fake.UpdateBrokerStub = nil
	fake.updateBrokerReturns = struct {
		result1 *types.Broker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateBrokerReturnsOnCall(i int, result1 *types.Broker, result2 error) {
	fake.updateBrokerMutex.Lock()
	defer fake.updateBrokerMutex.Unlock()
	fake.UpdateBrokerStub = nil
	if fake.updateBrokerReturnsOnCall == nil {
		fake.updateBrokerReturnsOnCall = make(map[int]struct {
			result1 *types.Broker
			result2 error
		})
	}
	fake.updateBrokerReturnsOnCall[i] = struct {
		result1 *types.Broker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdatePlatform(arg1 string, arg2 *types.Platform) (*types.Platform, error) {
	fake.updatePlatformMutex.Lock()
	ret, specificReturn := fake.updatePlatformReturnsOnCall[len(fake.updatePlatformArgsForCall)]
	fake.updatePlatformArgsForCall = append(fake.updatePlatformArgsForCall, struct {
		arg1 string
		arg2 *types.Platform
	}{arg1, arg2})
	fake.recordInvocation("UpdatePlatform", []interface{}{arg1, arg2})
	fake.updatePlatformMutex.Unlock()
	if fake.UpdatePlatformStub != nil {
		return fake.UpdatePlatformStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updatePlatformReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) UpdatePlatformCallCount() int {
	fake.updatePlatformMutex.RLock()
	defer fake.updatePlatformMutex.RUnlock()
	return len(fake.updatePlatformArgsForCall)
}

func (fake *FakeClient) UpdatePlatformCalls(stub func(string, *types.Platform) (*types.Platform, error)) {
	fake.updatePlatformMutex.Lock()
	defer fake.updatePlatformMutex.Unlock()
	fake.UpdatePlatformStub = stub
}

func (fake *FakeClient) UpdatePlatformArgsForCall(i int) (string, *types.Platform) {
	fake.updatePlatformMutex.RLock()
	defer fake.updatePlatformMutex.RUnlock()
	argsForCall := fake.updatePlatformArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) UpdatePlatformReturns(result1 *types.Platform, result2 error) {
	fake.updatePlatformMutex.Lock()
	defer fake.updatePlatformMutex.Unlock()
	fake.UpdatePlatformStub = nil
	fake.updatePlatformReturns = struct {
		result1 *types.Platform
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdatePlatformReturnsOnCall(i int, result1 *types.Platform, result2 error) {
	fake.updatePlatformMutex.Lock()
	defer fake.updatePlatformMutex.Unlock()
	fake.UpdatePlatformStub = nil
	if fake.updatePlatformReturnsOnCall == nil {
		fake.updatePlatformReturnsOnCall = make(map[int]struct {
			result1 *types.Platform
			result2 error
		})
	}
	fake.updatePlatformReturnsOnCall[i] = struct {
		result1 *types.Platform
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	fake.deleteBrokerMutex.RLock()
	defer fake.deleteBrokerMutex.RUnlock()
	fake.deletePlatformMutex.RLock()
	defer fake.deletePlatformMutex.RUnlock()
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	fake.listBrokersMutex.RLock()
	defer fake.listBrokersMutex.RUnlock()
	fake.listBrokersWithQueryMutex.RLock()
	defer fake.listBrokersWithQueryMutex.RUnlock()
	fake.listOfferingsMutex.RLock()
	defer fake.listOfferingsMutex.RUnlock()
	fake.listOfferingsWithQueryMutex.RLock()
	defer fake.listOfferingsWithQueryMutex.RUnlock()
	fake.listPlatformsMutex.RLock()
	defer fake.listPlatformsMutex.RUnlock()
	fake.listPlatformsWithQueryMutex.RLock()
	defer fake.listPlatformsWithQueryMutex.RUnlock()
	fake.registerBrokerMutex.RLock()
	defer fake.registerBrokerMutex.RUnlock()
	fake.registerPlatformMutex.RLock()
	defer fake.registerPlatformMutex.RUnlock()
	fake.updateBrokerMutex.RLock()
	defer fake.updateBrokerMutex.RUnlock()
	fake.updatePlatformMutex.RLock()
	defer fake.updatePlatformMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ smclient.Client = new(FakeClient)
