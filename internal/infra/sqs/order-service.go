package sqs

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type SendOrderServiceEventSqs struct {
	sqs *sqs.SQS
}

func NewSendOrderServiceEventSqs() *SendOrderServiceEventSqs {
	session, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewSharedCredentials("", "cybertech"),
		},
	)

	if err != nil {
		panic(err)
	}

	sqs := sqs.New(session)
	return &SendOrderServiceEventSqs{
		sqs: sqs,
	}
}

func (s *SendOrderServiceEventSqs) Send(order entities.OrderService) error {
	qURl := `https://sqs.us-east-1.amazonaws.com/608105930645/order_processed`
	o, _ := json.Marshal(order)
	result, err := s.sqs.SendMessage(
		&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"Title": &sqs.MessageAttributeValue{
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
	sqs *sqs.SQS
}

func NewListenOrderServiceEventSqs() *ListenOrderServiceEventSqs {
	session := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))
	sqs := sqs.New(session)
	return &ListenOrderServiceEventSqs{
		sqs: sqs,
	}
}

func (l *ListenOrderServiceEventSqs) GetEvents() {
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
		qURl := `https://sqs.us-east-1.amazonaws.com/608105930645/order_processed`
		result, err := l.sqs.ReceiveMessage(
			&sqs.ReceiveMessageInput{
				AttributeNames: []*string{
					aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
				},
				MessageAttributeNames: []*string{
					aws.String(sqs.QueueAttributeNameAll),
				},
				QueueUrl:            &qURl,
				MaxNumberOfMessages: aws.Int64(10),
				VisibilityTimeout:   aws.Int64(60),
				WaitTimeSeconds:     aws.Int64(0),
			},
		)

		if err != nil {
			log.Println("failed to fetch messages: %v", err)
		}

		for _, message := range result.Messages {
			ch <- message
		}
	}
}

func (l *ListenOrderServiceEventSqs) processMessage(msg *sqs.Message) error {

}

func (l *ListenOrderServiceEventSqs) deleteMessage(msg *sqs.Message) error {

}
