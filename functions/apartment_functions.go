package functions

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func ApartmentList(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (resp events.APIGatewayProxyResponse) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Apartment List",
	}
}

func ApartmentCreate(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (resp events.APIGatewayProxyResponse) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Apartment Create",
	}
}

func ApartmentGet(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (resp events.APIGatewayProxyResponse) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Apartment Get",
	}
}

func ApartmentUpdate(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (resp events.APIGatewayProxyResponse) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Apartment Update",
	}
}

func ApartmentDelete(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (resp events.APIGatewayProxyResponse) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Apartment Delete",
	}
}
