// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/radar"
)

type FakeRadarDB struct {
	GetPipelineNameStub        func() string
	getPipelineNameMutex       sync.RWMutex
	getPipelineNameArgsForCall []struct{}
	getPipelineNameReturns     struct {
		result1 string
	}
	ScopedNameStub        func(string) string
	scopedNameMutex       sync.RWMutex
	scopedNameArgsForCall []struct {
		arg1 string
	}
	scopedNameReturns struct {
		result1 string
	}
	IsPausedStub        func() (bool, error)
	isPausedMutex       sync.RWMutex
	isPausedArgsForCall []struct{}
	isPausedReturns     struct {
		result1 bool
		result2 error
	}
	GetConfigStub        func() (atc.Config, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct{}
	getConfigReturns     struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}
	GetLatestVersionedResourceStub        func(resource db.SavedResource) (db.SavedVersionedResource, error)
	getLatestVersionedResourceMutex       sync.RWMutex
	getLatestVersionedResourceArgsForCall []struct {
		resource db.SavedResource
	}
	getLatestVersionedResourceReturns struct {
		result1 db.SavedVersionedResource
		result2 error
	}
	GetResourceStub        func(resourceName string) (db.SavedResource, error)
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		resourceName string
	}
	getResourceReturns struct {
		result1 db.SavedResource
		result2 error
	}
	PauseResourceStub        func(resourceName string) error
	pauseResourceMutex       sync.RWMutex
	pauseResourceArgsForCall []struct {
		resourceName string
	}
	pauseResourceReturns struct {
		result1 error
	}
	UnpauseResourceStub        func(resourceName string) error
	unpauseResourceMutex       sync.RWMutex
	unpauseResourceArgsForCall []struct {
		resourceName string
	}
	unpauseResourceReturns struct {
		result1 error
	}
	SaveResourceVersionsStub        func(atc.ResourceConfig, []atc.Version) error
	saveResourceVersionsMutex       sync.RWMutex
	saveResourceVersionsArgsForCall []struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}
	saveResourceVersionsReturns struct {
		result1 error
	}
	SetResourceCheckErrorStub        func(resource db.SavedResource, err error) error
	setResourceCheckErrorMutex       sync.RWMutex
	setResourceCheckErrorArgsForCall []struct {
		resource db.SavedResource
		err      error
	}
	setResourceCheckErrorReturns struct {
		result1 error
	}
	LeaseCheckStub        func(resource string, interval time.Duration) (db.Lease, bool, error)
	leaseCheckMutex       sync.RWMutex
	leaseCheckArgsForCall []struct {
		resource string
		interval time.Duration
	}
	leaseCheckReturns struct {
		result1 db.Lease
		result2 bool
		result3 error
	}
}

func (fake *FakeRadarDB) GetPipelineName() string {
	fake.getPipelineNameMutex.Lock()
	fake.getPipelineNameArgsForCall = append(fake.getPipelineNameArgsForCall, struct{}{})
	fake.getPipelineNameMutex.Unlock()
	if fake.GetPipelineNameStub != nil {
		return fake.GetPipelineNameStub()
	} else {
		return fake.getPipelineNameReturns.result1
	}
}

func (fake *FakeRadarDB) GetPipelineNameCallCount() int {
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	return len(fake.getPipelineNameArgsForCall)
}

func (fake *FakeRadarDB) GetPipelineNameReturns(result1 string) {
	fake.GetPipelineNameStub = nil
	fake.getPipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRadarDB) ScopedName(arg1 string) string {
	fake.scopedNameMutex.Lock()
	fake.scopedNameArgsForCall = append(fake.scopedNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.scopedNameMutex.Unlock()
	if fake.ScopedNameStub != nil {
		return fake.ScopedNameStub(arg1)
	} else {
		return fake.scopedNameReturns.result1
	}
}

func (fake *FakeRadarDB) ScopedNameCallCount() int {
	fake.scopedNameMutex.RLock()
	defer fake.scopedNameMutex.RUnlock()
	return len(fake.scopedNameArgsForCall)
}

func (fake *FakeRadarDB) ScopedNameArgsForCall(i int) string {
	fake.scopedNameMutex.RLock()
	defer fake.scopedNameMutex.RUnlock()
	return fake.scopedNameArgsForCall[i].arg1
}

func (fake *FakeRadarDB) ScopedNameReturns(result1 string) {
	fake.ScopedNameStub = nil
	fake.scopedNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRadarDB) IsPaused() (bool, error) {
	fake.isPausedMutex.Lock()
	fake.isPausedArgsForCall = append(fake.isPausedArgsForCall, struct{}{})
	fake.isPausedMutex.Unlock()
	if fake.IsPausedStub != nil {
		return fake.IsPausedStub()
	} else {
		return fake.isPausedReturns.result1, fake.isPausedReturns.result2
	}
}

func (fake *FakeRadarDB) IsPausedCallCount() int {
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	return len(fake.isPausedArgsForCall)
}

func (fake *FakeRadarDB) IsPausedReturns(result1 bool, result2 error) {
	fake.IsPausedStub = nil
	fake.isPausedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRadarDB) GetConfig() (atc.Config, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct{}{})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub()
	} else {
		return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3
	}
}

func (fake *FakeRadarDB) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeRadarDB) GetConfigReturns(result1 atc.Config, result2 db.ConfigVersion, result3 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRadarDB) GetLatestVersionedResource(resource db.SavedResource) (db.SavedVersionedResource, error) {
	fake.getLatestVersionedResourceMutex.Lock()
	fake.getLatestVersionedResourceArgsForCall = append(fake.getLatestVersionedResourceArgsForCall, struct {
		resource db.SavedResource
	}{resource})
	fake.getLatestVersionedResourceMutex.Unlock()
	if fake.GetLatestVersionedResourceStub != nil {
		return fake.GetLatestVersionedResourceStub(resource)
	} else {
		return fake.getLatestVersionedResourceReturns.result1, fake.getLatestVersionedResourceReturns.result2
	}
}

func (fake *FakeRadarDB) GetLatestVersionedResourceCallCount() int {
	fake.getLatestVersionedResourceMutex.RLock()
	defer fake.getLatestVersionedResourceMutex.RUnlock()
	return len(fake.getLatestVersionedResourceArgsForCall)
}

func (fake *FakeRadarDB) GetLatestVersionedResourceArgsForCall(i int) db.SavedResource {
	fake.getLatestVersionedResourceMutex.RLock()
	defer fake.getLatestVersionedResourceMutex.RUnlock()
	return fake.getLatestVersionedResourceArgsForCall[i].resource
}

func (fake *FakeRadarDB) GetLatestVersionedResourceReturns(result1 db.SavedVersionedResource, result2 error) {
	fake.GetLatestVersionedResourceStub = nil
	fake.getLatestVersionedResourceReturns = struct {
		result1 db.SavedVersionedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeRadarDB) GetResource(resourceName string) (db.SavedResource, error) {
	fake.getResourceMutex.Lock()
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.getResourceMutex.Unlock()
	if fake.GetResourceStub != nil {
		return fake.GetResourceStub(resourceName)
	} else {
		return fake.getResourceReturns.result1, fake.getResourceReturns.result2
	}
}

func (fake *FakeRadarDB) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeRadarDB) GetResourceArgsForCall(i int) string {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return fake.getResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) GetResourceReturns(result1 db.SavedResource, result2 error) {
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 db.SavedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeRadarDB) PauseResource(resourceName string) error {
	fake.pauseResourceMutex.Lock()
	fake.pauseResourceArgsForCall = append(fake.pauseResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.pauseResourceMutex.Unlock()
	if fake.PauseResourceStub != nil {
		return fake.PauseResourceStub(resourceName)
	} else {
		return fake.pauseResourceReturns.result1
	}
}

func (fake *FakeRadarDB) PauseResourceCallCount() int {
	fake.pauseResourceMutex.RLock()
	defer fake.pauseResourceMutex.RUnlock()
	return len(fake.pauseResourceArgsForCall)
}

func (fake *FakeRadarDB) PauseResourceArgsForCall(i int) string {
	fake.pauseResourceMutex.RLock()
	defer fake.pauseResourceMutex.RUnlock()
	return fake.pauseResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) PauseResourceReturns(result1 error) {
	fake.PauseResourceStub = nil
	fake.pauseResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) UnpauseResource(resourceName string) error {
	fake.unpauseResourceMutex.Lock()
	fake.unpauseResourceArgsForCall = append(fake.unpauseResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.unpauseResourceMutex.Unlock()
	if fake.UnpauseResourceStub != nil {
		return fake.UnpauseResourceStub(resourceName)
	} else {
		return fake.unpauseResourceReturns.result1
	}
}

func (fake *FakeRadarDB) UnpauseResourceCallCount() int {
	fake.unpauseResourceMutex.RLock()
	defer fake.unpauseResourceMutex.RUnlock()
	return len(fake.unpauseResourceArgsForCall)
}

func (fake *FakeRadarDB) UnpauseResourceArgsForCall(i int) string {
	fake.unpauseResourceMutex.RLock()
	defer fake.unpauseResourceMutex.RUnlock()
	return fake.unpauseResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) UnpauseResourceReturns(result1 error) {
	fake.UnpauseResourceStub = nil
	fake.unpauseResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) SaveResourceVersions(arg1 atc.ResourceConfig, arg2 []atc.Version) error {
	fake.saveResourceVersionsMutex.Lock()
	fake.saveResourceVersionsArgsForCall = append(fake.saveResourceVersionsArgsForCall, struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}{arg1, arg2})
	fake.saveResourceVersionsMutex.Unlock()
	if fake.SaveResourceVersionsStub != nil {
		return fake.SaveResourceVersionsStub(arg1, arg2)
	} else {
		return fake.saveResourceVersionsReturns.result1
	}
}

func (fake *FakeRadarDB) SaveResourceVersionsCallCount() int {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return len(fake.saveResourceVersionsArgsForCall)
}

func (fake *FakeRadarDB) SaveResourceVersionsArgsForCall(i int) (atc.ResourceConfig, []atc.Version) {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return fake.saveResourceVersionsArgsForCall[i].arg1, fake.saveResourceVersionsArgsForCall[i].arg2
}

func (fake *FakeRadarDB) SaveResourceVersionsReturns(result1 error) {
	fake.SaveResourceVersionsStub = nil
	fake.saveResourceVersionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) SetResourceCheckError(resource db.SavedResource, err error) error {
	fake.setResourceCheckErrorMutex.Lock()
	fake.setResourceCheckErrorArgsForCall = append(fake.setResourceCheckErrorArgsForCall, struct {
		resource db.SavedResource
		err      error
	}{resource, err})
	fake.setResourceCheckErrorMutex.Unlock()
	if fake.SetResourceCheckErrorStub != nil {
		return fake.SetResourceCheckErrorStub(resource, err)
	} else {
		return fake.setResourceCheckErrorReturns.result1
	}
}

func (fake *FakeRadarDB) SetResourceCheckErrorCallCount() int {
	fake.setResourceCheckErrorMutex.RLock()
	defer fake.setResourceCheckErrorMutex.RUnlock()
	return len(fake.setResourceCheckErrorArgsForCall)
}

func (fake *FakeRadarDB) SetResourceCheckErrorArgsForCall(i int) (db.SavedResource, error) {
	fake.setResourceCheckErrorMutex.RLock()
	defer fake.setResourceCheckErrorMutex.RUnlock()
	return fake.setResourceCheckErrorArgsForCall[i].resource, fake.setResourceCheckErrorArgsForCall[i].err
}

func (fake *FakeRadarDB) SetResourceCheckErrorReturns(result1 error) {
	fake.SetResourceCheckErrorStub = nil
	fake.setResourceCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) LeaseCheck(resource string, interval time.Duration) (db.Lease, bool, error) {
	fake.leaseCheckMutex.Lock()
	fake.leaseCheckArgsForCall = append(fake.leaseCheckArgsForCall, struct {
		resource string
		interval time.Duration
	}{resource, interval})
	fake.leaseCheckMutex.Unlock()
	if fake.LeaseCheckStub != nil {
		return fake.LeaseCheckStub(resource, interval)
	} else {
		return fake.leaseCheckReturns.result1, fake.leaseCheckReturns.result2, fake.leaseCheckReturns.result3
	}
}

func (fake *FakeRadarDB) LeaseCheckCallCount() int {
	fake.leaseCheckMutex.RLock()
	defer fake.leaseCheckMutex.RUnlock()
	return len(fake.leaseCheckArgsForCall)
}

func (fake *FakeRadarDB) LeaseCheckArgsForCall(i int) (string, time.Duration) {
	fake.leaseCheckMutex.RLock()
	defer fake.leaseCheckMutex.RUnlock()
	return fake.leaseCheckArgsForCall[i].resource, fake.leaseCheckArgsForCall[i].interval
}

func (fake *FakeRadarDB) LeaseCheckReturns(result1 db.Lease, result2 bool, result3 error) {
	fake.LeaseCheckStub = nil
	fake.leaseCheckReturns = struct {
		result1 db.Lease
		result2 bool
		result3 error
	}{result1, result2, result3}
}

var _ radar.RadarDB = new(FakeRadarDB)
