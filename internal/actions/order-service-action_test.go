package actions_test

import (
	"testing"

	"github.com/henriquerocha2004/cyber-tech-go/internal/actions"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/henriquerocha2004/cyber-tech-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	orderServiceStatus = entities.OrderServiceStatus{
		Id:              1,
		Description:     "Aberto",
		LaunchFinancial: false,
		Color:           "red",
	}
	orderServiceInput = actions.ServiceOrderInput{
		Description: "Ordem referente a manutenção de micro",
		StatusId:    1,
		Paid:        false,
		Items: []actions.ItemsInput{
			{
				ProductId: 1,
				Type:      "product",
				Quantity:  1,
				Value:     300.00,
			},
			{
				ProductId: 2,
				Type:      "service",
				Quantity:  1,
				Value:     80.00,
			},
		},
		Equipments: []actions.EquipmentInput{
			{
				Description:  "CPU Login",
				Defect:       "Não Liga",
				Observations: "Cpu com lacres violados",
			},
		},
	}
	orderServiceInputWithoutStatusId = actions.ServiceOrderInput{
		Description: "Ordem referente a manutenção de micro",
		Paid:        false,
		StatusId:    13,
	}
)

func TestOrderService(t *testing.T) {
	orderServiceCommandRepo := new(mocks.OrderServiceRepositoryCommandMock)
	orderServiceCommandRepo.On("Create", mock.Anything).Return(1, nil)
	orderServiceQueryRepo := new(mocks.OrderServiceRepositoryQueryMock)
	productRepository := new(mocks.ProductRepositoryMock)
	orderServiceStatusRepo := new(mocks.OrderServiceStatusRepositoryMock)

	t.Run("should create a service order", func(t *testing.T) {
		orderServiceStatusRepo.On("FindOne", 1).Return(orderServiceStatus, nil)
		orderServiceActions := actions.NewServiceOrderActions(
			orderServiceCommandRepo,
			orderServiceQueryRepo,
			orderServiceStatusRepo,
			productRepository,
		)
		output := orderServiceActions.Create(orderServiceInput)
		assert.False(t, output.Error)
		assert.NotNil(t, output.Data)
	})
	t.Run("should calculate total of items", func(t *testing.T) {
		order := &entities.OrderService{
			Number:      "000211122",
			Description: "Manutenção de computador",
			Items: []entities.OrderItem{
				{
					ProductId: 1,
					Type:      "product",
					Quantity:  1,
					Value:     300.00,
				},
				{
					ProductId: 2,
					Type:      "service",
					Quantity:  1,
					Value:     80.00,
				},
			},
		}

		total, err := order.GetTotal()
		assert.NoError(t, err)
		assert.Equal(t, 380.00, total)
	})
	t.Run("should return error if not send order status", func(t *testing.T) {
		orderServiceStatusRepo.On("FindOne", 13).Return(entities.OrderServiceStatus{}, nil)
		orderServiceActions := actions.NewServiceOrderActions(
			orderServiceCommandRepo,
			orderServiceQueryRepo,
			orderServiceStatusRepo,
			productRepository,
		)
		output := orderServiceActions.Create(orderServiceInputWithoutStatusId)
		assert.True(t, output.Error)
		assert.Equal(t, "Failed to create order: Failed to get status order", output.Message)
	})
	t.Run("should return error if order number not informed when update order", func(t *testing.T) {
		orderServiceActions := actions.NewServiceOrderActions(
			orderServiceCommandRepo,
			orderServiceQueryRepo,
			orderServiceStatusRepo,
			productRepository,
		)

		output := orderServiceActions.Update(orderServiceInput)
		assert.True(t, output.Error)
		assert.Equal(t, "Failed to create order: Order Number not informed", output.Message)
	})
}
