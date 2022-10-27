package sqs_test

import (
	"testing"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/sqs"
	"github.com/stretchr/testify/assert"
)

func TestConnectionAws(t *testing.T) {
	sqsTest := sqs.NewSendOrderServiceEventSqs()
	order := entities.OrderService{
		Id:          1,
		Number:      "20221013171981",
		Description: "referente a manutenção de smartphone",
		StatusId:    1,
		Paid:        true,
		Items: []entities.OrderItem{
			{
				Id:        1,
				ProductId: 1,
				OrderId:   1,
				Type:      "service",
				Quantity:  1,
				Value:     130.00,
			},
			{
				Id:        2,
				ProductId: 1,
				OrderId:   1,
				Type:      "product",
				Quantity:  1,
				Value:     200.00,
			},
		},
		Equipments: []entities.Equipment{
			{
				Id:          1,
				Description: "Iphone X",
				Defect:      "Não Liga",
				OrderId:     1,
			},
		},
	}
	err := sqsTest.Send(order)
	assert.NoError(t, err)
}

func TestReceiveMessage(t *testing.T) {
	sqsReceive := sqs.NewListenOrderServiceEventSqs()
	sqsReceive.GetEvents()
}
