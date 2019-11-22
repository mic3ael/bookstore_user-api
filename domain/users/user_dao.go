package users

import (
	"fmt"

	"github.com/mic3ael/bookstore_user-api/utils/errors"
)

var (
	usersDB = make(map[uint64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundRequestError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstNmae = result.FirstNmae
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {

		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	usersDB[user.ID] = user
	return nil
}
