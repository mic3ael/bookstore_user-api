package users

import (
	"strings"

	"github.com/mic3ael/bookstore_user-api/utils/errors"
)

const (
	StatusActive = "active"
)

// User ...
type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedOn uint64 `json:"created_on"`
	Deleted   bool   `json:"-"`
	UpdatedOn uint64 `json:"updated_on"`
	Status    string `json:"status"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
