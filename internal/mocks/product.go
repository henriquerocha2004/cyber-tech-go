package mocks

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (p *ProductRepositoryMock) Create(product entities.Product) error {
	args := p.Called(product)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (p *ProductRepositoryMock) Update(product entities.Product) error {
	args := p.Called(product)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (p *ProductRepositoryMock) Delete(id int) error {
	args := p.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (p *ProductRepositoryMock) FindOne(id int) (entities.Product, error) {
	args := p.Called(id)
	if args.Get(0) == nil {
		return args.Get(0).(entities.Product), nil
	}
	return args.Get(0).(entities.Product), args.Get(1).(error)
}

func (p *ProductRepositoryMock) FindAll() ([]entities.Product, error) {
	args := p.Called()
	if args.Get(0) == nil {
		return args.Get(0).([]entities.Product), nil
	}
	return args.Get(0).([]entities.Product), args.Get(1).(error)
}
