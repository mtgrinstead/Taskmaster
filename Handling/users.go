package Handling

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/libsql/libsql-client-go/libsql"
	"net/http"
)

type user struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	CreatedDate  string `json:"createdDate"`
	PasswordHash string `json:"password"`
	Role         int    `json:"role"`
}

// DATABASE FUNCTIONS

func GetAllUsers(c *gin.Context) {
	// Returns all users in database
	var dbUrl = "libsql://taskmaster-mtgrinstead.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDIzLTEwLTAyVDA1OjMwOjQxLjk3ODY1NjE1OVoiLCJpZCI6IjhhMzIzMDE4LTVkYWUtMTFlZS04YjVjLTMyNzE3OTI2MDEzYSJ9.cUDuRNAWL21Zf1kT0StQYCuP4FT0JQYaHYr8aCiCV9c-ghzTcvXJVxOoqoNY5HViAFEm7uPLF1N6jJ2YreCvBg"
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	query := "SELECT ID, Name, Email, CreatedDate, Role FROM users"
	rows, err := db.Query(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var userdb user
		err := rows.Scan(&userdb.ID, &userdb.Name, &userdb.Email, &userdb.CreatedDate, &userdb.Role)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, userdb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {}

func AddUser(c *gin.Context) {
	var newUser user
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the user name.
	if newUser.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User Name must not be empty"})
		return
	}

	// Validate the user email.
	if newUser.Email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User Email must not be empty"})
		return
	}

	// Validate the user password.
	for _, password := range newUser.PasswordHash {
		if password == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Password must not be empty"})
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

	stmt, err := db.Prepare("INSERT INTO users (Name, Email, PasswordHash, CreatedDate, Role) VALUES (?, ?, ?, DATE('now'), ?)")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Bind the user data to the prepared statement parameters.
	_, err = stmt.Exec(newUser.Name, newUser.Email, newUser.PasswordHash, newUser.Role)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the prepared statement.
	defer stmt.Close()

	// Return a success message to the user.
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}

var users = []user{
	{ID: "1", Name: "Mom", Email: "momma@yahoo.com", CreatedDate: "Today", PasswordHash: "password", Role: 0},
	{ID: "2", Name: "Dad", Email: "boobies@gmail.com", CreatedDate: "Today", PasswordHash: "boobies", Role: 1},
	{ID: "3", Name: "Dog", Email: "treats@gmail.com", CreatedDate: "Yesterday", PasswordHash: "goodboi", Role: 1},
	{ID: "4", Name: "Baby", Email: "worlddomination@gmail.com", CreatedDate: "July 4", PasswordHash: "P@S5w0rd", Role: 0},
}

// LOCALHOST FUNCTIONS
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func UserById(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func getUserById(id string) (*user, error) {
	for i, t := range users {
		if t.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func PromoteRole(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found."})
		return
	}

	user.Role = 0
	c.IndentedJSON(http.StatusOK, user)
}

func DemoteRole(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found."})
		return
	}

	user.Role = 1
	c.IndentedJSON(http.StatusOK, user)
}

func UpdatePassword(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
