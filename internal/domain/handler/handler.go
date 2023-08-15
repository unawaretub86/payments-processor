package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/unawaretub86/payments-processor/internal/domain/usecase"
)

func HttpHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	lc, _ := lambdacontext.FromContext(ctx)

	requestId := lc.AwsRequestID

	body := request.Body

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"message\": \"Payment processed\"}",
	}

	return response, usecase.ConvertPaymentRequest(body, requestId)
}
