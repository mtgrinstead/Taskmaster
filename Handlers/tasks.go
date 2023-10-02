package Handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
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

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func TaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func UpdateStatusInProgress(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	task, err := GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}

	task.Status = "In-Progress"
	c.IndentedJSON(http.StatusOK, task)
}

func UpdateStatusCompleted(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	task, err := GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}

	task.Status = "Completed"
	c.IndentedJSON(http.StatusOK, task)
}

func UpdateStatusTesting(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	task, err := GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
		return
	}

	task.Status = "Testing"
	c.IndentedJSON(http.StatusOK, task)
}

func GetTaskById(id string) (*task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
