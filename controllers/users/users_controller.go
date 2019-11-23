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

func getUserId(userIdParam string) (uint64, *errors.RestErr) {
	userID, userErr := strconv.ParseUint(userIdParam, 10, 64)
	if userErr != nil {
		fmt.Println("err: ", userErr)
		return 0, errors.NewBadRequestError("user id should be a positive number")
	}
	return userID, nil
}

// GetUser ...
func GetUser(c *gin.Context) {
	userID, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
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

func UpdateUser(c *gin.Context) {
	userID, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		//TODO: Handle error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userID, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
