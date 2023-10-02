package main

import (
	"github.com/gin-gonic/gin"
	"main/Handlers"
)

func main() {

	router := gin.Default()
	router.GET("getme", Handlers.CheckMe)
	router.GET("/tasks", Handlers.GetTasks)
	router.GET("/tasks/:id", Handlers.TaskById)
	router.POST("/tasks", Handlers.CreateTask)
	router.PATCH("/in-progress", Handlers.UpdateStatusInProgress)
	router.PATCH("/testing", Handlers.UpdateStatusTesting)
	router.PATCH("/completed", Handlers.UpdateStatusCompleted)

	router.GET("/users", Handlers.GetUsers)
	router.GET("/users/:id", Handlers.UserById)
	router.POST("/users", Handlers.CreateUser)
	router.PATCH("/password", Handlers.UpdatePassword)
	router.PATCH("/promote", Handlers.PromoteRole)
	router.PATCH("/demote", Handlers.DemoteRole)
	router.Run("localhost:8080")

}

//     eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTA5LTI4VDAzOjI2OjEwLjY1NjUyNTcyWiIsImlkIjoiOGEzMjMwMTgtNWRhZS0xMWVlLThiNWMtMzI3MTc5MjYwMTNhIn0.-GAjnWlihe6H8aoRfD_0B6GQ00YfERJQS35NfL_3ri1J8PoBFQOAKVxl9tbV9j5Bh2d03s6k9pi3_Zr3jgTDBg
