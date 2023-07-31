package main

import (
	"context"
	"log"
	"poc/pkg/api"
	"poc/pkg/dal"
	"poc/pkg/db"

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
	// 		Body:       string(`[{"id": "uuid-something", "itemName": "name", "itemImageUrl": "https://r.mtdv.me/OgM5EWc3KQ"}]`),
	// 	}, nil
	// }

	conn, err := db.Connect(ctx)

	if err != nil {
		log.Printf("Unable to connect to database: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to connect to database")
	}
	defer conn.Close(ctx)

	dao := dal.Db(conn)

	items, err := dao.Item.GetAll(ctx)
	if err != nil {
		log.Printf("Unable to get items from the database: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to get items from the database")
	}

	data, err := json.Marshal(items)
	if err != nil {
		log.Printf("Unable to marshal output: %s", err.Error())
		return api.InternalServerErrorResponse("Unable to marshal output")
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}
