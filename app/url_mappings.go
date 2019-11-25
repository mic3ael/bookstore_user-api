package app

import (
	"github.com/mic3ael/bookstore_user-api/controllers/users"
)

// MapUrls ...
func mapUrls() {
	router.GET("/users/search", users.Search)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
	router.PUT("/users/:id", users.UpdateUser)
	router.PATCH("/users/:id", users.UpdateUser)
	router.DELETE("/users/:id", users.DeleteUser)
}
