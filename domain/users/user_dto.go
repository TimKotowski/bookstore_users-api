package users

import (
	"bookstore_users-api/utils/errors"
	"strings"
)

// data transfer obect
// going to be the data we are tansfering from the persistence layer to the aplciation and backwards
// working with a user then the user will be in the dto because the user is the onject we are going to be moving between the persitence layer and applicaiton
const (
	StatusActive = "active"
)

type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` // if we get a password field in the json do not take that json address of the json field
}

func (user *User) Validate() *errors.RestErr {
	// pointer revicer to allow the mothod to mutate the receiving struct
	user.Firstname = strings.TrimSpace(strings.ToLower(user.Firstname))
	user.Lastname = strings.TrimSpace(strings.ToLower(user.Lastname))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
