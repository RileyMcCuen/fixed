package api

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func NoContentResponse() (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}

func BadRequestResponse(reason string) (events.APIGatewayV2HTTPResponse, error) {
	log.Println("BadRequestResponse:", reason)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusBadRequest,
		Body:       reason,
	}, nil
}

func InternalServerErrorResponse(reason string) (events.APIGatewayV2HTTPResponse, error) {
	log.Println("InternalServerErrorResponse:", reason)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       reason,
	}, nil
}
