package Handling

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Due       string `json:"due"`
	Status    string `json:"status"`
	Assignees string `json:"assignee"`
}

// Database functions
func GetAllTasks(c *gin.Context) {
	// Returns all tasks in database
	var dbUrl = "libsql://taskmaster-mtgrinstead.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTEwLTAyVDA1OjMwOjQxLjk3ODY1NjE1OVoiLCJpZCI6IjhhMzIzMDE4LTVkYWUtMTFlZS04YjVjLTMyNzE3OTI2MDEzYSJ9.cUDuRNAWL21Zf1kT0StQYCuP4FT0JQYaHYr8aCiCV9c-ghzTcvXJVxOoqoNY5HViAFEm7uPLF1N6jJ2YreCvBg"
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	query := "SELECT ID, Title, Status, Assignees FROM tasks"
	rows, err := db.Query(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var taskdb task
		err := rows.Scan(&taskdb.ID, &taskdb.Title, &taskdb.Status, &taskdb.Assignees)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, taskdb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByTitle(c *gin.Context) {}

func AddTask(c *gin.Context) {
	// Validate the task data.
	var newTask task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the task title.
	if newTask.Title == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Task title must not be empty"})
		return
	}

	// Validate the task assignees.
	for _, assignee := range newTask.Assignees {
		if assignee == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Task assignee must not be empty"})
			return
		}
	}

	// Create a prepared statement.
	var dbUrl = "libsql://taskmaster-mtgrinstead.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTEwLTAyVDA1OjMwOjQxLjk3ODY1NjE1OVoiLCJpZCI6IjhhMzIzMDE4LTVkYWUtMTFlZS04YjVjLTMyNzE3OTI2MDEzYSJ9.cUDuRNAWL21Zf1kT0StQYCuP4FT0JQYaHYr8aCiCV9c-ghzTcvXJVxOoqoNY5HViAFEm7uPLF1N6jJ2YreCvBg"
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute SQL query: " + err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO tasks (Title, Status, Assignees) VALUES (?, ?, ?)")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Bind the task data to the prepared statement parameters.
	_, err = stmt.Exec(newTask.Title, newTask.Status, newTask.Assignees)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the prepared statement.
	defer stmt.Close()

	// Return a success message to the user.
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task added successfully"})
}

var tasks = []task{
	{ID: "1", Title: "Test1", Due: "Christmas", Status: "Not Started", Assignees: "Mom"},
	{ID: "2", Title: "Test2", Due: "November", Status: "In-Progress", Assignees: "Dada"},
	{ID: "3", Title: "Test3", Due: "July 4th", Status: "Completed", Assignees: "Baby"},
	{ID: "4", Title: "Test4", Due: "Tomorrow", Status: "Testing", Assignees: "Toddler"},
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
