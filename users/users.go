package users

import (
	"errors"
	"github.com/gin-gonic/gin"
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

var users = []user{
	{ID: "1", Name: "Mom", Email: "momma@yahoo.com", CreatedDate: "Today", PasswordHash: "password", Role: 0},
	{ID: "2", Name: "Dad", Email: "boobies@gmail.com", CreatedDate: "Today", PasswordHash: "boobies", Role: 1},
	{ID: "3", Name: "Dog", Email: "treats@gmail.com", CreatedDate: "Yesterday", PasswordHash: "goodboi", Role: 1},
	{ID: "4", Name: "Baby", Email: "worlddomination@gmail.com", CreatedDate: "July 4", PasswordHash: "P@S5w0rd", Role: 0},
}

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
