package api

import (
	"errors"
	"poc/pkg/util"

	"github.com/aws/aws-lambda-go/events"
)

func UserId(event events.APIGatewayV2HTTPRequest) (userId string, err error) {
	defer util.ErrorWrapper("on request while getting userId", &err)

	ctx := event.RequestContext
	if ctx.Authorizer == nil {
		return "", errors.New("nil authorizer")
	}

	authorizer := ctx.Authorizer
	if authorizer.JWT == nil {
		return "", errors.New("nil JWT authorizer")
	}

	jwt := authorizer.JWT
	if jwt.Claims == nil {
		return "", errors.New("nil JWT claims")
	}

	userId, ok := event.RequestContext.Authorizer.JWT.Claims["sub"]
	if !ok {
		return "", errors.New("no sub attribute on JWT claims")
	}

	return userId, nil
}
