package sqs

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/henriquerocha2004/cyber-tech-go/cmd/api/handlers"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type SendOrderServiceEventSqs struct {
	sqs *sqs.SQS
}

func NewSendOrderServiceEventSqs() *SendOrderServiceEventSqs {
	session := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	sqs := sqs.New(session)
	return &SendOrderServiceEventSqs{
		sqs: sqs,
	}
}

func (s *SendOrderServiceEventSqs) Send(order entities.OrderService) error {
	qURl := viper.Get("queue.url_order_service_queue").(string)
	o, _ := json.Marshal(order)
	result, err := s.sqs.SendMessage(
		&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"Title": {
					DataType:    aws.String("String"),
					StringValue: aws.String("order created"),
				},
			},
			MessageBody: aws.String(string(o)),
			QueueUrl:    &qURl,
		},
	)

	if err != nil {
		return err
	}

	log.Println(result.MessageId)
	return nil
}

type ListenOrderServiceEventSqs struct {
	sqs               *sqs.SQS
	queueUrl          string
	orderServiceQueue *handlers.OrderServiceQueueHandler
}

func NewListenOrderServiceEventSqs(orderServiceQueueHandler *handlers.OrderServiceQueueHandler) *ListenOrderServiceEventSqs {
	session := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))
	sqs := sqs.New(session)
	url := viper.Get("queue.url_order_service_queue").(string)
	return &ListenOrderServiceEventSqs{
		sqs:               sqs,
		queueUrl:          url,
		orderServiceQueue: orderServiceQueueHandler,
	}
}

func (l *ListenOrderServiceEventSqs) GetEvents() {
	log.Println("listen events...")
	messages := make(chan *sqs.Message, 2)
	go l.pollingMessages(messages)
	for message := range messages {
		err := l.processMessage(message)
		if err != nil {
			log.Println(err)
		}
		l.deleteMessage(message)
	}
}

func (l *ListenOrderServiceEventSqs) pollingMessages(ch chan<- *sqs.Message) {
	for {
		result, err := l.sqs.ReceiveMessage(
			&sqs.ReceiveMessageInput{
				AttributeNames: []*string{
					aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
				},
				MessageAttributeNames: []*string{
					aws.String(sqs.QueueAttributeNameAll),
				},
				QueueUrl:            &l.queueUrl,
				MaxNumberOfMessages: aws.Int64(10),
				VisibilityTimeout:   aws.Int64(60),
				WaitTimeSeconds:     aws.Int64(0),
			},
		)
		if err != nil {
			log.Println(err)
		}

		for _, message := range result.Messages {
			ch <- message
		}

		time.Sleep(time.Second * 15)
	}
}

func (l *ListenOrderServiceEventSqs) processMessage(msg *sqs.Message) error {
	return l.orderServiceQueue.Distribute(*msg.Body)
}

func (l *ListenOrderServiceEventSqs) deleteMessage(msg *sqs.Message) {
	_, err := l.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(l.queueUrl),
		ReceiptHandle: msg.ReceiptHandle,
	})

	log.Println("Error in delete message: " + err.Error())
}
