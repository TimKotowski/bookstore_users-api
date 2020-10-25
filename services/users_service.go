package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
)

// entry business logic in services

// functions to handle a get user rqeuest
// attepmt to get a user from the database
// points to the memory location of the struct created
//func CreateUser(user *users.User)
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	return nil, nil
};
