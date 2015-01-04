// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/exec"
)

type FakeStep struct {
	UsingStub        func(exec.ArtifactSource) exec.ArtifactSource
	usingMutex       sync.RWMutex
	usingArgsForCall []struct {
		arg1 exec.ArtifactSource
	}
	usingReturns struct {
		result1 exec.ArtifactSource
	}
}

func (fake *FakeStep) Using(arg1 exec.ArtifactSource) exec.ArtifactSource {
	fake.usingMutex.Lock()
	fake.usingArgsForCall = append(fake.usingArgsForCall, struct {
		arg1 exec.ArtifactSource
	}{arg1})
	fake.usingMutex.Unlock()
	if fake.UsingStub != nil {
		return fake.UsingStub(arg1)
	} else {
		return fake.usingReturns.result1
	}
}

func (fake *FakeStep) UsingCallCount() int {
	fake.usingMutex.RLock()
	defer fake.usingMutex.RUnlock()
	return len(fake.usingArgsForCall)
}

func (fake *FakeStep) UsingArgsForCall(i int) exec.ArtifactSource {
	fake.usingMutex.RLock()
	defer fake.usingMutex.RUnlock()
	return fake.usingArgsForCall[i].arg1
}

func (fake *FakeStep) UsingReturns(result1 exec.ArtifactSource) {
	fake.UsingStub = nil
	fake.usingReturns = struct {
		result1 exec.ArtifactSource
	}{result1}
}

var _ exec.Step = new(FakeStep)