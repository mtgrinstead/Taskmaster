package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/users"

	"net/http"
)

type task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Due      string `json:"due"`
	Status   string `json:"status"`
	Assignee string `json:"assignee"`
}

var tasks = []task{
	{ID: "1", Title: "Test1", Due: "Christmas", Status: "Not Started", Assignee: "Mom"},
	{ID: "2", Title: "Test2", Due: "November", Status: "In-Progress", Assignee: "Dada"},
	{ID: "3", Title: "Test3", Due: "July 4th", Status: "Completed", Assignee: "Baby"},
	{ID: "4", Title: "Test4", Due: "Tomorrow", Status: "Testing", Assignee: "Toddler"},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func taskById(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func updateStatusInProgress(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}

	task.Status = "In-Progress"
	c.IndentedJSON(http.StatusOK, task)
}

func updateStatusCompleted(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}

	task.Status = "Completed"
	c.IndentedJSON(http.StatusOK, task)
}

func updateStatusTesting(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}

	task.Status = "Testing"
	c.IndentedJSON(http.StatusOK, task)
}

func getTaskById(id string) (*task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func createTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", taskById)
	router.POST("/tasks", createTask)
	router.PATCH("/in-progress", updateStatusInProgress)
	router.PATCH("/testing", updateStatusTesting)
	router.PATCH("/completed", updateStatusCompleted)

	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.UserById)
	router.POST("/users", users.CreateUser)
	//router.PATCH("/password", updatePassword)
	router.PATCH("/promote", users.PromoteRole)
	router.PATCH("/demote", users.DemoteRole)
	router.Run("localhost:8080")
}

//     eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTA5LTI4VDAzOjI2OjEwLjY1NjUyNTcyWiIsImlkIjoiOGEzMjMwMTgtNWRhZS0xMWVlLThiNWMtMzI3MTc5MjYwMTNhIn0.-GAjnWlihe6H8aoRfD_0B6GQ00YfERJQS35NfL_3ri1J8PoBFQOAKVxl9tbV9j5Bh2d03s6k9pi3_Zr3jgTDBg
