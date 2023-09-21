package handlers

//  contains a function which takes status and body parameters and creates valid APIGatewayProxyResponse
import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiRsponse(status int,body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers:map[string]string{"Content-Type":"application/json"}}
	resp.StatusCode = status

	sBody, _ := json.Marshal(body)
	resp.body = string(sBody)
	return &resp, nil 
}