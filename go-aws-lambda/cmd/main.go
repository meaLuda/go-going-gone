package main

import (
	// local libraries
	"aws-lambda-in-golang/pkg/handlers"
	"os"

	// external
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// database client
var (dynaClient dynamodbiface.DynamoDBAPI)

func main(){
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(
		&aws.Config{Region:aws.String(region)},
	)
	if err != nil{
		return
	}
	dynaClient := dynamodb.NewSession(awsSession)
	lambda.Start(handler)
}

const tableName = "<my_table_name>"

func handler(req events.APIGatewayProxyRequest)(*events.APIGatewayProxyResponse,error){
	switch req.HTTPMethod{
		// map request header method to apropriate handler functions
		case "GET":
			return handler.GetUser(req,tableName,dynaClient)
		default:
			return 
		
	}
}