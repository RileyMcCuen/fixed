package main

import (
	"context"
	"log"
	"net/http"
	"poc/pkg/api"
	"poc/pkg/dal"
	"poc/pkg/db"
	"poc/pkg/model"

	"github.com/go-json-experiment/json"

	"github.com/RileyMcCuen/llb"
	"github.com/RileyMcCuen/llb/pkg/handlerutil"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

func main() {
	llb.Start(handlerutil.InOutTypeHandler(Handler, nil))
}

func Handler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// if event.QueryStringParameters["mock"] == "true" {
	// 	return events.APIGatewayV2HTTPResponse{
	// 		StatusCode: http.StatusCreated,
	// 		Body:       string(""),
	// 	}, nil
	// }

	item := model.Item{}
	if err := json.Unmarshal([]byte(event.Body), &item); err != nil {
		return api.BadRequestResponse(err.Error())
	}

	item.Id = uuid.NewString()

	conn, err := db.Connect(ctx)

	if err != nil {
		log.Printf("Unable to connect to database: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to connect to database")
	}
	defer conn.Close(ctx)

	dao := dal.Db(conn)

	if err := dao.Item.Insert(ctx, item); err != nil {
		log.Printf("Unable to insert item into database: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to insert item into database")
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Body:       string(""),
	}, nil
}
