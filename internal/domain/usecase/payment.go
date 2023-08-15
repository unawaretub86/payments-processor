package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/unawaretub86/payments-processor/internal/domain/entities"
)

func ConvertPaymentRequest(body, requestId string) error {
	var paymentRequest entities.PaymentRequest
	err := json.Unmarshal([]byte(body), &paymentRequest)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", err, requestId)
		return err
	}

	return sendSQS(paymentRequest, requestId)
}

func sendSQS(payment entities.PaymentRequest, requestId string) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsClient := sqs.New(sess)

	queueURL := os.Getenv("SQS_URL")

	paymentJSON, err := json.Marshal(payment)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling payment request: %v]", err, requestId)
		return err
	}

	_, err = sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(paymentJSON)),
		QueueUrl:    &queueURL,
	})

	if err != nil {
		fmt.Printf("[RequestId: %s][Error sending message to SQS: %v]", err, requestId)
		return err
	}

	return nil
}
