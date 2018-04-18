// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type DeployedConfigService struct {
	GetDeployedProductByNameStub        func(product string) (api.DeployedProductsFindOutput, error)
	getDeployedProductByNameMutex       sync.RWMutex
	getDeployedProductByNameArgsForCall []struct {
		product string
	}
	getDeployedProductByNameReturns struct {
		result1 api.DeployedProductsFindOutput
		result2 error
	}
	getDeployedProductByNameReturnsOnCall map[int]struct {
		result1 api.DeployedProductsFindOutput
		result2 error
	}
	ListDeployedProductJobsStub        func(productGUID string) (map[string]string, error)
	listDeployedProductJobsMutex       sync.RWMutex
	listDeployedProductJobsArgsForCall []struct {
		productGUID string
	}
	listDeployedProductJobsReturns struct {
		result1 map[string]string
		result2 error
	}
	listDeployedProductJobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetDeployedProductJobResourceConfigStub        func(productGUID, jobGUID string) (api.JobProperties, error)
	getDeployedProductJobResourceConfigMutex       sync.RWMutex
	getDeployedProductJobResourceConfigArgsForCall []struct {
		productGUID string
		jobGUID     string
	}
	getDeployedProductJobResourceConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getDeployedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	GetDeployedProductPropertiesStub        func(product string) (map[string]api.ResponseProperty, error)
	getDeployedProductPropertiesMutex       sync.RWMutex
	getDeployedProductPropertiesArgsForCall []struct {
		product string
	}
	getDeployedProductPropertiesReturns struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	getDeployedProductPropertiesReturnsOnCall map[int]struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	GetDeployedProductNetworksAndAZsStub        func(product string) (map[string]interface{}, error)
	getDeployedProductNetworksAndAZsMutex       sync.RWMutex
	getDeployedProductNetworksAndAZsArgsForCall []struct {
		product string
	}
	getDeployedProductNetworksAndAZsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	getDeployedProductNetworksAndAZsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	GetDeployedProductCredentialStub        func(product string, credentialName string) (api.CredentialOutput, error)
	getDeployedProductCredentialMutex       sync.RWMutex
	getDeployedProductCredentialArgsForCall []struct {
		product        string
		credentialName string
	}
	getDeployedProductCredentialReturns struct {
		result1 api.CredentialOutput
		result2 error
	}
	getDeployedProductCredentialReturnsOnCall map[int]struct {
		result1 api.CredentialOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeployedConfigService) GetDeployedProductByName(product string) (api.DeployedProductsFindOutput, error) {
	fake.getDeployedProductByNameMutex.Lock()
	ret, specificReturn := fake.getDeployedProductByNameReturnsOnCall[len(fake.getDeployedProductByNameArgsForCall)]
	fake.getDeployedProductByNameArgsForCall = append(fake.getDeployedProductByNameArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetDeployedProductByName", []interface{}{product})
	fake.getDeployedProductByNameMutex.Unlock()
	if fake.GetDeployedProductByNameStub != nil {
		return fake.GetDeployedProductByNameStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductByNameReturns.result1, fake.getDeployedProductByNameReturns.result2
}

func (fake *DeployedConfigService) GetDeployedProductByNameCallCount() int {
	fake.getDeployedProductByNameMutex.RLock()
	defer fake.getDeployedProductByNameMutex.RUnlock()
	return len(fake.getDeployedProductByNameArgsForCall)
}

func (fake *DeployedConfigService) GetDeployedProductByNameArgsForCall(i int) string {
	fake.getDeployedProductByNameMutex.RLock()
	defer fake.getDeployedProductByNameMutex.RUnlock()
	return fake.getDeployedProductByNameArgsForCall[i].product
}

func (fake *DeployedConfigService) GetDeployedProductByNameReturns(result1 api.DeployedProductsFindOutput, result2 error) {
	fake.GetDeployedProductByNameStub = nil
	fake.getDeployedProductByNameReturns = struct {
		result1 api.DeployedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductByNameReturnsOnCall(i int, result1 api.DeployedProductsFindOutput, result2 error) {
	fake.GetDeployedProductByNameStub = nil
	if fake.getDeployedProductByNameReturnsOnCall == nil {
		fake.getDeployedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.DeployedProductsFindOutput
			result2 error
		})
	}
	fake.getDeployedProductByNameReturnsOnCall[i] = struct {
		result1 api.DeployedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) ListDeployedProductJobs(productGUID string) (map[string]string, error) {
	fake.listDeployedProductJobsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductJobsReturnsOnCall[len(fake.listDeployedProductJobsArgsForCall)]
	fake.listDeployedProductJobsArgsForCall = append(fake.listDeployedProductJobsArgsForCall, struct {
		productGUID string
	}{productGUID})
	fake.recordInvocation("ListDeployedProductJobs", []interface{}{productGUID})
	fake.listDeployedProductJobsMutex.Unlock()
	if fake.ListDeployedProductJobsStub != nil {
		return fake.ListDeployedProductJobsStub(productGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDeployedProductJobsReturns.result1, fake.listDeployedProductJobsReturns.result2
}

func (fake *DeployedConfigService) ListDeployedProductJobsCallCount() int {
	fake.listDeployedProductJobsMutex.RLock()
	defer fake.listDeployedProductJobsMutex.RUnlock()
	return len(fake.listDeployedProductJobsArgsForCall)
}

func (fake *DeployedConfigService) ListDeployedProductJobsArgsForCall(i int) string {
	fake.listDeployedProductJobsMutex.RLock()
	defer fake.listDeployedProductJobsMutex.RUnlock()
	return fake.listDeployedProductJobsArgsForCall[i].productGUID
}

func (fake *DeployedConfigService) ListDeployedProductJobsReturns(result1 map[string]string, result2 error) {
	fake.ListDeployedProductJobsStub = nil
	fake.listDeployedProductJobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) ListDeployedProductJobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ListDeployedProductJobsStub = nil
	if fake.listDeployedProductJobsReturnsOnCall == nil {
		fake.listDeployedProductJobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.listDeployedProductJobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductJobResourceConfig(productGUID string, jobGUID string) (api.JobProperties, error) {
	fake.getDeployedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.getDeployedProductJobResourceConfigReturnsOnCall[len(fake.getDeployedProductJobResourceConfigArgsForCall)]
	fake.getDeployedProductJobResourceConfigArgsForCall = append(fake.getDeployedProductJobResourceConfigArgsForCall, struct {
		productGUID string
		jobGUID     string
	}{productGUID, jobGUID})
	fake.recordInvocation("GetDeployedProductJobResourceConfig", []interface{}{productGUID, jobGUID})
	fake.getDeployedProductJobResourceConfigMutex.Unlock()
	if fake.GetDeployedProductJobResourceConfigStub != nil {
		return fake.GetDeployedProductJobResourceConfigStub(productGUID, jobGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductJobResourceConfigReturns.result1, fake.getDeployedProductJobResourceConfigReturns.result2
}

func (fake *DeployedConfigService) GetDeployedProductJobResourceConfigCallCount() int {
	fake.getDeployedProductJobResourceConfigMutex.RLock()
	defer fake.getDeployedProductJobResourceConfigMutex.RUnlock()
	return len(fake.getDeployedProductJobResourceConfigArgsForCall)
}

func (fake *DeployedConfigService) GetDeployedProductJobResourceConfigArgsForCall(i int) (string, string) {
	fake.getDeployedProductJobResourceConfigMutex.RLock()
	defer fake.getDeployedProductJobResourceConfigMutex.RUnlock()
	return fake.getDeployedProductJobResourceConfigArgsForCall[i].productGUID, fake.getDeployedProductJobResourceConfigArgsForCall[i].jobGUID
}

func (fake *DeployedConfigService) GetDeployedProductJobResourceConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetDeployedProductJobResourceConfigStub = nil
	fake.getDeployedProductJobResourceConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductJobResourceConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.GetDeployedProductJobResourceConfigStub = nil
	if fake.getDeployedProductJobResourceConfigReturnsOnCall == nil {
		fake.getDeployedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getDeployedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductProperties(product string) (map[string]api.ResponseProperty, error) {
	fake.getDeployedProductPropertiesMutex.Lock()
	ret, specificReturn := fake.getDeployedProductPropertiesReturnsOnCall[len(fake.getDeployedProductPropertiesArgsForCall)]
	fake.getDeployedProductPropertiesArgsForCall = append(fake.getDeployedProductPropertiesArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetDeployedProductProperties", []interface{}{product})
	fake.getDeployedProductPropertiesMutex.Unlock()
	if fake.GetDeployedProductPropertiesStub != nil {
		return fake.GetDeployedProductPropertiesStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductPropertiesReturns.result1, fake.getDeployedProductPropertiesReturns.result2
}

func (fake *DeployedConfigService) GetDeployedProductPropertiesCallCount() int {
	fake.getDeployedProductPropertiesMutex.RLock()
	defer fake.getDeployedProductPropertiesMutex.RUnlock()
	return len(fake.getDeployedProductPropertiesArgsForCall)
}

func (fake *DeployedConfigService) GetDeployedProductPropertiesArgsForCall(i int) string {
	fake.getDeployedProductPropertiesMutex.RLock()
	defer fake.getDeployedProductPropertiesMutex.RUnlock()
	return fake.getDeployedProductPropertiesArgsForCall[i].product
}

func (fake *DeployedConfigService) GetDeployedProductPropertiesReturns(result1 map[string]api.ResponseProperty, result2 error) {
	fake.GetDeployedProductPropertiesStub = nil
	fake.getDeployedProductPropertiesReturns = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductPropertiesReturnsOnCall(i int, result1 map[string]api.ResponseProperty, result2 error) {
	fake.GetDeployedProductPropertiesStub = nil
	if fake.getDeployedProductPropertiesReturnsOnCall == nil {
		fake.getDeployedProductPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]api.ResponseProperty
			result2 error
		})
	}
	fake.getDeployedProductPropertiesReturnsOnCall[i] = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductNetworksAndAZs(product string) (map[string]interface{}, error) {
	fake.getDeployedProductNetworksAndAZsMutex.Lock()
	ret, specificReturn := fake.getDeployedProductNetworksAndAZsReturnsOnCall[len(fake.getDeployedProductNetworksAndAZsArgsForCall)]
	fake.getDeployedProductNetworksAndAZsArgsForCall = append(fake.getDeployedProductNetworksAndAZsArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetDeployedProductNetworksAndAZs", []interface{}{product})
	fake.getDeployedProductNetworksAndAZsMutex.Unlock()
	if fake.GetDeployedProductNetworksAndAZsStub != nil {
		return fake.GetDeployedProductNetworksAndAZsStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductNetworksAndAZsReturns.result1, fake.getDeployedProductNetworksAndAZsReturns.result2
}

func (fake *DeployedConfigService) GetDeployedProductNetworksAndAZsCallCount() int {
	fake.getDeployedProductNetworksAndAZsMutex.RLock()
	defer fake.getDeployedProductNetworksAndAZsMutex.RUnlock()
	return len(fake.getDeployedProductNetworksAndAZsArgsForCall)
}

func (fake *DeployedConfigService) GetDeployedProductNetworksAndAZsArgsForCall(i int) string {
	fake.getDeployedProductNetworksAndAZsMutex.RLock()
	defer fake.getDeployedProductNetworksAndAZsMutex.RUnlock()
	return fake.getDeployedProductNetworksAndAZsArgsForCall[i].product
}

func (fake *DeployedConfigService) GetDeployedProductNetworksAndAZsReturns(result1 map[string]interface{}, result2 error) {
	fake.GetDeployedProductNetworksAndAZsStub = nil
	fake.getDeployedProductNetworksAndAZsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductNetworksAndAZsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.GetDeployedProductNetworksAndAZsStub = nil
	if fake.getDeployedProductNetworksAndAZsReturnsOnCall == nil {
		fake.getDeployedProductNetworksAndAZsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getDeployedProductNetworksAndAZsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductCredential(product string, credentialName string) (api.CredentialOutput, error) {
	fake.getDeployedProductCredentialMutex.Lock()
	ret, specificReturn := fake.getDeployedProductCredentialReturnsOnCall[len(fake.getDeployedProductCredentialArgsForCall)]
	fake.getDeployedProductCredentialArgsForCall = append(fake.getDeployedProductCredentialArgsForCall, struct {
		product        string
		credentialName string
	}{product, credentialName})
	fake.recordInvocation("GetDeployedProductCredential", []interface{}{product, credentialName})
	fake.getDeployedProductCredentialMutex.Unlock()
	if fake.GetDeployedProductCredentialStub != nil {
		return fake.GetDeployedProductCredentialStub(product, credentialName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductCredentialReturns.result1, fake.getDeployedProductCredentialReturns.result2
}

func (fake *DeployedConfigService) GetDeployedProductCredentialCallCount() int {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return len(fake.getDeployedProductCredentialArgsForCall)
}

func (fake *DeployedConfigService) GetDeployedProductCredentialArgsForCall(i int) (string, string) {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return fake.getDeployedProductCredentialArgsForCall[i].product, fake.getDeployedProductCredentialArgsForCall[i].credentialName
}

func (fake *DeployedConfigService) GetDeployedProductCredentialReturns(result1 api.CredentialOutput, result2 error) {
	fake.GetDeployedProductCredentialStub = nil
	fake.getDeployedProductCredentialReturns = struct {
		result1 api.CredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) GetDeployedProductCredentialReturnsOnCall(i int, result1 api.CredentialOutput, result2 error) {
	fake.GetDeployedProductCredentialStub = nil
	if fake.getDeployedProductCredentialReturnsOnCall == nil {
		fake.getDeployedProductCredentialReturnsOnCall = make(map[int]struct {
			result1 api.CredentialOutput
			result2 error
		})
	}
	fake.getDeployedProductCredentialReturnsOnCall[i] = struct {
		result1 api.CredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *DeployedConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDeployedProductByNameMutex.RLock()
	defer fake.getDeployedProductByNameMutex.RUnlock()
	fake.listDeployedProductJobsMutex.RLock()
	defer fake.listDeployedProductJobsMutex.RUnlock()
	fake.getDeployedProductJobResourceConfigMutex.RLock()
	defer fake.getDeployedProductJobResourceConfigMutex.RUnlock()
	fake.getDeployedProductPropertiesMutex.RLock()
	defer fake.getDeployedProductPropertiesMutex.RUnlock()
	fake.getDeployedProductNetworksAndAZsMutex.RLock()
	defer fake.getDeployedProductNetworksAndAZsMutex.RUnlock()
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeployedConfigService) recordInvocation(key string, args []interface{}) {
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