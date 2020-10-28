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

// *users.User means im returning a variable that stoed an addres of another variable
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Go automatically handles conversion between values and pointers for method calls
	// so no need to do &user
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := users.User{
		ID: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
