package handlers

import (
	// local libraries
	"aws-lambda-in-golang/pkg/users"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrMethodNotAllowed = "405"

type ErrBody struct{
	// if string is empty return json is empty[https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html#:~:text=A%20string%20in%20Go%20is,a%20string%20cannot%20be%20nil%20.&text=However%2C%20a%20pointer%20to%20a,*string%20)%20can%20be%20nil%20.&text=A%20good%20rule%20of%20thumb,strings%20unless%20you%20need%20nil%20.]
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest,tableName string,dynoClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse,error){
	// get email from req
	email := req.QueryStringParameters["email"]

	if len(email) >0{
		// get single user
		result,errr := user.FetchUser(email,tableName,dynoClient)
	}

}