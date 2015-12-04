// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakePipelineDBFactory struct {
	BuildStub        func(pipeline db.SavedPipeline) db.PipelineDB
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		pipeline db.SavedPipeline
	}
	buildReturns struct {
		result1 db.PipelineDB
	}
	BuildWithTeamNameAndNameStub        func(teamName, pipelineName string) (db.PipelineDB, error)
	buildWithTeamNameAndNameMutex       sync.RWMutex
	buildWithTeamNameAndNameArgsForCall []struct {
		teamName     string
		pipelineName string
	}
	buildWithTeamNameAndNameReturns struct {
		result1 db.PipelineDB
		result2 error
	}
	BuildDefaultStub        func() (db.PipelineDB, bool, error)
	buildDefaultMutex       sync.RWMutex
	buildDefaultArgsForCall []struct{}
	buildDefaultReturns     struct {
		result1 db.PipelineDB
		result2 bool
		result3 error
	}
}

func (fake *FakePipelineDBFactory) Build(pipeline db.SavedPipeline) db.PipelineDB {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		pipeline db.SavedPipeline
	}{pipeline})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(pipeline)
	} else {
		return fake.buildReturns.result1
	}
}

func (fake *FakePipelineDBFactory) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakePipelineDBFactory) BuildArgsForCall(i int) db.SavedPipeline {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].pipeline
}

func (fake *FakePipelineDBFactory) BuildReturns(result1 db.PipelineDB) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 db.PipelineDB
	}{result1}
}

func (fake *FakePipelineDBFactory) BuildWithTeamNameAndName(teamName string, pipelineName string) (db.PipelineDB, error) {
	fake.buildWithTeamNameAndNameMutex.Lock()
	fake.buildWithTeamNameAndNameArgsForCall = append(fake.buildWithTeamNameAndNameArgsForCall, struct {
		teamName     string
		pipelineName string
	}{teamName, pipelineName})
	fake.buildWithTeamNameAndNameMutex.Unlock()
	if fake.BuildWithTeamNameAndNameStub != nil {
		return fake.BuildWithTeamNameAndNameStub(teamName, pipelineName)
	} else {
		return fake.buildWithTeamNameAndNameReturns.result1, fake.buildWithTeamNameAndNameReturns.result2
	}
}

func (fake *FakePipelineDBFactory) BuildWithTeamNameAndNameCallCount() int {
	fake.buildWithTeamNameAndNameMutex.RLock()
	defer fake.buildWithTeamNameAndNameMutex.RUnlock()
	return len(fake.buildWithTeamNameAndNameArgsForCall)
}

func (fake *FakePipelineDBFactory) BuildWithTeamNameAndNameArgsForCall(i int) (string, string) {
	fake.buildWithTeamNameAndNameMutex.RLock()
	defer fake.buildWithTeamNameAndNameMutex.RUnlock()
	return fake.buildWithTeamNameAndNameArgsForCall[i].teamName, fake.buildWithTeamNameAndNameArgsForCall[i].pipelineName
}

func (fake *FakePipelineDBFactory) BuildWithTeamNameAndNameReturns(result1 db.PipelineDB, result2 error) {
	fake.BuildWithTeamNameAndNameStub = nil
	fake.buildWithTeamNameAndNameReturns = struct {
		result1 db.PipelineDB
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDBFactory) BuildDefault() (db.PipelineDB, bool, error) {
	fake.buildDefaultMutex.Lock()
	fake.buildDefaultArgsForCall = append(fake.buildDefaultArgsForCall, struct{}{})
	fake.buildDefaultMutex.Unlock()
	if fake.BuildDefaultStub != nil {
		return fake.BuildDefaultStub()
	} else {
		return fake.buildDefaultReturns.result1, fake.buildDefaultReturns.result2, fake.buildDefaultReturns.result3
	}
}

func (fake *FakePipelineDBFactory) BuildDefaultCallCount() int {
	fake.buildDefaultMutex.RLock()
	defer fake.buildDefaultMutex.RUnlock()
	return len(fake.buildDefaultArgsForCall)
}

func (fake *FakePipelineDBFactory) BuildDefaultReturns(result1 db.PipelineDB, result2 bool, result3 error) {
	fake.BuildDefaultStub = nil
	fake.buildDefaultReturns = struct {
		result1 db.PipelineDB
		result2 bool
		result3 error
	}{result1, result2, result3}
}

var _ db.PipelineDBFactory = new(FakePipelineDBFactory)
