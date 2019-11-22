package users

import (
	"strings"

	"github.com/mic3ael/bookstore_user-api/utils/errors"
)

// User ...
type User struct {
	ID        uint64 `json:"id"`
	FirstNmae string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt uint64 `json:"created_at"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
