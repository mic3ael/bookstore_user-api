package app

import (
	"github.com/mic3ael/bookstore_user-api/controllers/users"
)

// MapUrls ...
func mapUrls() {
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
}
