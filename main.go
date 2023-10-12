package main

import (
	"github.com/gin-gonic/gin"
	"main/Handling"
)

func main() {
	//lambda.Start(handler)

	router := gin.Default()

	router.GET("/dbusers", Handling.GetAllUsers)
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
	//router.Run("74.208.209.25:8443")

}

//     eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTA5LTI4VDAzOjI2OjEwLjY1NjUyNTcyWiIsImlkIjoiOGEzMjMwMTgtNWRhZS0xMWVlLThiNWMtMzI3MTc5MjYwMTNhIn0.-GAjnWlihe6H8aoRfD_0B6GQ00YfERJQS35NfL_3ri1J8PoBFQOAKVxl9tbV9j5Bh2d03s6k9pi3_Zr3jgTDBg
