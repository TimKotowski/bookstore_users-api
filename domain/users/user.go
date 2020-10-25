package users

import (
	"bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func Validate(user *User) *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email request")
	}
	return nil
}
