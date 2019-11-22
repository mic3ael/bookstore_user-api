package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mic3ael/bookstore_user-api/domain/users"
	"github.com/mic3ael/bookstore_user-api/services"
	"github.com/mic3ael/bookstore_user-api/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetUser ...
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if userErr != nil {
		fmt.Println("err: ", userErr)
		err := errors.NewBadRequestError("user id should be a positive number")

		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser ...
func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		//TODO: Handle error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// SearchUser ...
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
