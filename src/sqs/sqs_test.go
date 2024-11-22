package sqs

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	sqs "github.com/aws/aws-sdk-go/service/sqs"
)

func TestSendMessage(t *testing.T) {

	// Crie uma sessão AWS
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		t.Fatalf("Erro ao criar sessão: %v", err)
	}

	svc := sqs.New(sess)

	// Chame a função SendMessage
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String("Olá isso é uma mensagem de TESTE!!!"),
		QueueUrl:    aws.String(SQS_URL),
	})
	if err != nil {
		t.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	// Verifique o resultado
	if result.MessageId == nil {
		t.Errorf("Esperava um MessageId, mas recebeu nil")
	}
}
