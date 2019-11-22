package services

import (
	"github.com/mic3ael/bookstore_user-api/domain/users"
	"github.com/mic3ael/bookstore_user-api/utils/errors"
)

// GetUser ...
func GetUser(userID uint64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUser ...
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUser ...
func FindUser() {

}
