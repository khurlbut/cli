// This file was generated by counterfeiter
package fakes

import (
	. "github.com/cloudfoundry/cli/cf/actors"
	"github.com/cloudfoundry/cli/cf/models"
	"sync"
)

type FakeServiceActor struct {
	GetAllBrokersWithDependenciesStub        func() ([]models.ServiceBroker, error)
	getAllBrokersWithDependenciesMutex       sync.RWMutex
	getAllBrokersWithDependenciesArgsForCall []struct{}
	getAllBrokersWithDependenciesReturns     struct {
		result1 []models.ServiceBroker
		result2 error
	}
	GetBrokerWithDependenciesStub        func(string) ([]models.ServiceBroker, error)
	getBrokerWithDependenciesMutex       sync.RWMutex
	getBrokerWithDependenciesArgsForCall []struct {
		arg1 string
	}
	getBrokerWithDependenciesReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
	GetBrokerWithSingleServiceStub        func(string) ([]models.ServiceBroker, error)
	getBrokerWithSingleServiceMutex       sync.RWMutex
	getBrokerWithSingleServiceArgsForCall []struct {
		arg1 string
	}
	getBrokerWithSingleServiceReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
}

func (fake *FakeServiceActor) GetAllBrokersWithDependencies() ([]models.ServiceBroker, error) {
	fake.getAllBrokersWithDependenciesMutex.Lock()
	defer fake.getAllBrokersWithDependenciesMutex.Unlock()
	fake.getAllBrokersWithDependenciesArgsForCall = append(fake.getAllBrokersWithDependenciesArgsForCall, struct{}{})
	if fake.GetAllBrokersWithDependenciesStub != nil {
		return fake.GetAllBrokersWithDependenciesStub()
	} else {
		return fake.getAllBrokersWithDependenciesReturns.result1, fake.getAllBrokersWithDependenciesReturns.result2
	}
}

func (fake *FakeServiceActor) GetAllBrokersWithDependenciesCallCount() int {
	fake.getAllBrokersWithDependenciesMutex.RLock()
	defer fake.getAllBrokersWithDependenciesMutex.RUnlock()
	return len(fake.getAllBrokersWithDependenciesArgsForCall)
}

func (fake *FakeServiceActor) GetAllBrokersWithDependenciesReturns(result1 []models.ServiceBroker, result2 error) {
	fake.getAllBrokersWithDependenciesReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceActor) GetBrokerWithDependencies(arg1 string) ([]models.ServiceBroker, error) {
	fake.getBrokerWithDependenciesMutex.Lock()
	defer fake.getBrokerWithDependenciesMutex.Unlock()
	fake.getBrokerWithDependenciesArgsForCall = append(fake.getBrokerWithDependenciesArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetBrokerWithDependenciesStub != nil {
		return fake.GetBrokerWithDependenciesStub(arg1)
	} else {
		return fake.getBrokerWithDependenciesReturns.result1, fake.getBrokerWithDependenciesReturns.result2
	}
}

func (fake *FakeServiceActor) GetBrokerWithDependenciesCallCount() int {
	fake.getBrokerWithDependenciesMutex.RLock()
	defer fake.getBrokerWithDependenciesMutex.RUnlock()
	return len(fake.getBrokerWithDependenciesArgsForCall)
}

func (fake *FakeServiceActor) GetBrokerWithDependenciesArgsForCall(i int) string {
	fake.getBrokerWithDependenciesMutex.RLock()
	defer fake.getBrokerWithDependenciesMutex.RUnlock()
	return fake.getBrokerWithDependenciesArgsForCall[i].arg1
}

func (fake *FakeServiceActor) GetBrokerWithDependenciesReturns(result1 []models.ServiceBroker, result2 error) {
	fake.getBrokerWithDependenciesReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceActor) GetBrokerWithSingleService(arg1 string) ([]models.ServiceBroker, error) {
	fake.getBrokerWithSingleServiceMutex.Lock()
	defer fake.getBrokerWithSingleServiceMutex.Unlock()
	fake.getBrokerWithSingleServiceArgsForCall = append(fake.getBrokerWithSingleServiceArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetBrokerWithSingleServiceStub != nil {
		return fake.GetBrokerWithSingleServiceStub(arg1)
	} else {
		return fake.getBrokerWithSingleServiceReturns.result1, fake.getBrokerWithSingleServiceReturns.result2
	}
}

func (fake *FakeServiceActor) GetBrokerWithSingleServiceCallCount() int {
	fake.getBrokerWithSingleServiceMutex.RLock()
	defer fake.getBrokerWithSingleServiceMutex.RUnlock()
	return len(fake.getBrokerWithSingleServiceArgsForCall)
}

func (fake *FakeServiceActor) GetBrokerWithSingleServiceArgsForCall(i int) string {
	fake.getBrokerWithSingleServiceMutex.RLock()
	defer fake.getBrokerWithSingleServiceMutex.RUnlock()
	return fake.getBrokerWithSingleServiceArgsForCall[i].arg1
}

func (fake *FakeServiceActor) GetBrokerWithSingleServiceReturns(result1 []models.ServiceBroker, result2 error) {
	fake.getBrokerWithSingleServiceReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

var _ ServiceActor = new(FakeServiceActor)
