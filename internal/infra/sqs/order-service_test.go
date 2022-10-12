package sqs_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/sqs"
	"github.com/stretchr/testify/assert"
)

func TestConnectionAws(t *testing.T) {
	sqsTest := sqs.NewSendOrderServiceEventSqs()
	order := entities.OrderService{
		Id:          1,
		Number:      "2020202125455555",
		Description: "Alguma descricao",
		StatusId:    1,
		Paid:        false,
	}

	err := sqsTest.Send(order)
	assert.NoError(t, err)
}

func TestReceiveMessage(t *testing.T) {
	sqsReceive := sqs.NewListenOrderServiceEventSqs()
	messages, err := sqsReceive.GetEvent()
	assert.NoError(t, err)
	m := messages[0].Body
	log.Println(*m)
	var order entities.OrderService
	_ = json.Unmarshal([]byte(*m), &order)
	log.Println(order)
	log.Println(messages[0].ReceiptHandle)
	assert.NotEmpty(t, messages)
}
