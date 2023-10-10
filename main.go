package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	"main/Handling"
)

type RequestBody struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Context string `json:"context"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	requestBody := RequestBody{}
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return nil, err
	}

	//if requestBody.Payload.Context == "production" {
	//	mediumautopost.Do("")
	//} else {
	//	fmt.Println("context" + requestBody.Payload.Context + " detected, skipping")
	//}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success",
	}, nil
}

func main() {
	lambda.Start(handler)

	router := gin.Default()
	router.GET("/getme", Handling.CheckMe)
	router.GET("/tasks", Handling.GetTasks)
	router.GET("/tasks/:id", Handling.TaskById)
	router.POST("/tasks", Handling.CreateTask)
	router.PATCH("/in-progress", Handling.UpdateStatusInProgress)
	router.PATCH("/testing", Handling.UpdateStatusTesting)
	router.PATCH("/completed", Handling.UpdateStatusCompleted)

	router.GET("/users", Handling.GetUsers)
	router.GET("/users/:id", Handling.UserById)
	router.POST("/users", Handling.CreateUser)
	router.PATCH("/password", Handling.UpdatePassword)
	router.PATCH("/promote", Handling.PromoteRole)
	router.PATCH("/demote", Handling.DemoteRole)
	router.Run("localhost:8080")

}

//     eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTA5LTI4VDAzOjI2OjEwLjY1NjUyNTcyWiIsImlkIjoiOGEzMjMwMTgtNWRhZS0xMWVlLThiNWMtMzI3MTc5MjYwMTNhIn0.-GAjnWlihe6H8aoRfD_0B6GQ00YfERJQS35NfL_3ri1J8PoBFQOAKVxl9tbV9j5Bh2d03s6k9pi3_Zr3jgTDBg
