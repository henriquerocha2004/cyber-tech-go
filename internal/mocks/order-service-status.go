package mocks

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/stretchr/testify/mock"
)

type OrderServiceStatusRepositoryMock struct {
	mock.Mock
}

func (o *OrderServiceStatusRepositoryMock) Create(status entities.OrderServiceStatus) error {
	args := o.Called(status)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (o *OrderServiceStatusRepositoryMock) Update(status entities.OrderServiceStatus) error {
	args := o.Called(status)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (o *OrderServiceStatusRepositoryMock) Delete(id int) error {
	args := o.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (o *OrderServiceStatusRepositoryMock) FindOne(id int) (entities.OrderServiceStatus, error) {
	args := o.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).(entities.OrderServiceStatus), nil
	}
	return args.Get(0).(entities.OrderServiceStatus), args.Get(1).(error)
}

func (o *OrderServiceStatusRepositoryMock) FindAll() ([]entities.OrderServiceStatus, error) {
	args := o.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]entities.OrderServiceStatus), nil
	}
	return args.Get(0).([]entities.OrderServiceStatus), args.Get(1).(error)
}
