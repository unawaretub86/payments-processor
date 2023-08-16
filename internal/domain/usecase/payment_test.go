package usecase_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"

	"github.com/unawaretub86/payments-processor/internal/domain/entities"
	"github.com/unawaretub86/payments-processor/internal/domain/usecase"
	"github.com/unawaretub86/payments-processor/internal/domain/usecase/mocks"
)

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := &mocks.Mocks{}

	validBody := `{"order_id": "1234pruebaCompleta"}`

	requestID := "1234567890"

	mockRepo.ConvertPaymentRequestFunc = func(order *entities.PaymentRequest, requestId string) (*string, error) {
		orderId := "order_id"
		return &orderId, nil
	}

	err := usecase.ConvertPaymentRequest(validBody, requestID)

	mockSQS := mocks.NewMockSQS("us-east-2")
	queueURL := "https://queue.amazonaws.com/80398EXAMPLE/MyQueue"

	messageAttributes := map[string]*sqs.MessageAttributeValue{
		"Source": {
			DataType:    aws.String("String"),
			StringValue: aws.String("payments-processor"),
		},
	}

	_, err = mockSQS.SendMessage(&sqs.SendMessageInput{
		MessageBody:       aws.String(`{"order_id": "1234567890"}`),
		QueueUrl:          &queueURL,
		MessageAttributes: messageAttributes,
	})
	if err != nil {
		t.Errorf("Error sending message: %v", err)
	}

	if err != nil {
		t.Errorf("Error updating payment: %v", err)
	}

	assert.NoError(t, err)
}

func TestCreatePayment_Error(t *testing.T) {
	mockRepo := &mocks.Mocks{}

	invalidBody := `{invalid_field: 1234567890}`
	requestID := "1234567890"

	expectedError := fmt.Errorf("invalid input data")

	mockRepo.ConvertPaymentRequestFunc = func(order *entities.PaymentRequest, requestId string) (*string, error) {
		userID := "user_id"
		return &userID, nil
	}

	err := usecase.ConvertPaymentRequest(invalidBody, requestID)

	assert.Error(t, err, expectedError)
}
