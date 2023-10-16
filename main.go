package main

import (
	"github.com/gin-gonic/gin"
	"main/Handling"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	//lambda.Start(handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		//w.Write([]byte('{"message": "Success"}'))
	})

	router := gin.Default()

	router.GET("/dbusers", Handling.GetAllUsers)
	router.GET("/dbtasks", Handling.GetAllTasks)
	router.POST("/newdbuser", Handling.AddUser)
	router.POST("/dbtasks", Handling.AddTask)
	router.DELETE("/deleteuser/:id", Handling.DeleteUser)
	router.GET("/tasks", Handling.GetTasks)
	router.GET("/tasks/:id", Handling.TaskById)
	router.POST("/tasks", Handling.CreateTask)
	router.PATCH("/in-progress", Handling.UpdateStatusInProgress)
	router.PATCH("/testing", Handling.UpdateStatusTesting)
	router.PATCH("/completed", Handling.UpdateStatusCompleted)

	router.GET("/users", Handling.GetUsers)
	router.GET("/users/:id", Handling.UserById)
	router.POST("/newuser", Handling.CreateUser)
	router.PATCH("/password", Handling.UpdatePassword)
	router.PATCH("/promote", Handling.PromoteRole)
	router.PATCH("/demote", Handling.DemoteRole)
	router.Run("localhost:8080")
	//router.Run("74.208.209.25:8443")

}

//     eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTA5LTI4VDAzOjI2OjEwLjY1NjUyNTcyWiIsImlkIjoiOGEzMjMwMTgtNWRhZS0xMWVlLThiNWMtMzI3MTc5MjYwMTNhIn0.-GAjnWlihe6H8aoRfD_0B6GQ00YfERJQS35NfL_3ri1J8PoBFQOAKVxl9tbV9j5Bh2d03s6k9pi3_Zr3jgTDBg
