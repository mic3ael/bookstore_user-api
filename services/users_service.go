package services

import (
	"strings"

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

// UpdateUser ...
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)

	if err != nil {
		return nil, err
	}

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	if !isPartial {
		current.FirstName = user.FirstName
		current.LastName = user.LastName

		if emailErr := user.Validate(); emailErr != nil {
			return nil, emailErr
		}

		current.Email = user.Email
	} else {

		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			if emailErr := user.Validate(); emailErr != nil {
				return nil, emailErr
			}
			current.Email = user.Email
		}
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userID uint64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func FindByStatus(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
