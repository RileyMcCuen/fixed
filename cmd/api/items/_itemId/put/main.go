package main

import (
	"context"
	"log"
	"poc/pkg/api"
	"poc/pkg/dal"
	"poc/pkg/db"
	"poc/pkg/model"

	"github.com/go-json-experiment/json"

	"github.com/RileyMcCuen/llb"
	"github.com/RileyMcCuen/llb/pkg/handlerutil"
	"github.com/aws/aws-lambda-go/events"
)

func main() {
	llb.Start(handlerutil.InOutTypeHandler(Handler, nil))
}

func Handler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// if event.QueryStringParameters["mock"] == "true" {
	// 	return events.APIGatewayV2HTTPResponse{
	// 		StatusCode: http.StatusOK,
	// 		Body:       string(`{"id": "uuid-something", "itemName": "name", "itemImageUrl": "https://r.mtdv.me/OgM5EWc3KQ"}`),
	// 	}, nil
	// }

	item := model.Item{}
	if err := json.Unmarshal([]byte(event.Body), &item); err != nil {
		return api.BadRequestResponse(err.Error())
	}

	item.Id = event.PathParameters["itemId"]

	conn, err := db.Connect(ctx)

	if err != nil {
		log.Printf("Unable to connect to database: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to connect to database")
	}
	defer conn.Close(ctx)

	dao := dal.Db(conn)

	if err = dao.Item.Update(ctx, item); err != nil {
		log.Printf("Unable to update from the database: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to update from the database")
	}

	data, err := json.Marshal(item)
	if err != nil {
		log.Printf("Unable to marshal output: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to marshal output")
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}
