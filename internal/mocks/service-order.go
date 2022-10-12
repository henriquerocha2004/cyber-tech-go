package mocks

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/stretchr/testify/mock"
)

type OrderServiceRepositoryCommandMock struct {
	mock.Mock
}

func (o *OrderServiceRepositoryCommandMock) Create(serviceOrder entities.OrderService) (int, error) {
	args := o.Called(serviceOrder)
	if args.Get(1) == nil {
		return args.Get(0).(int), nil
	}
	return 0, args.Get(0).(error)
}

func (o *OrderServiceRepositoryCommandMock) Update(serviceOrder entities.OrderService) error {
	args := o.Called(serviceOrder)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

type OrderServiceRepositoryQueryMock struct {
	mock.Mock
}

func (o *OrderServiceRepositoryQueryMock) FindOne(id int) (entities.OrderService, error) {
	args := o.Called(id)
	if args.Get(0) == nil {
		return args.Get(0).(entities.OrderService), nil
	}
	return args.Get(0).(entities.OrderService), args.Get(1).(error)
}

func (o *OrderServiceRepositoryQueryMock) FindAll() ([]entities.OrderService, error) {
	args := o.Called()
	if args.Get(0) == nil {
		return args.Get(0).([]entities.OrderService), nil
	}
	return args.Get(0).([]entities.OrderService), args.Get(1).(error)
}

func (o *OrderServiceRepositoryQueryMock) GetOrderItems(orderId int) ([]entities.OrderItem, error) {
	args := o.Called(orderId)
	if args.Get(0) == nil {
		return args.Get(0).([]entities.OrderItem), nil
	}
	return args.Get(0).([]entities.OrderItem), args.Get(1).(error)
}

func (o *OrderServiceRepositoryQueryMock) GetEquipments(orderId int) ([]entities.Equipment, error) {
	args := o.Called(orderId)
	if args.Get(0) == nil {
		return args.Get(0).([]entities.Equipment), nil
	}
	return args.Get(0).([]entities.Equipment), args.Get(1).(error)
}
