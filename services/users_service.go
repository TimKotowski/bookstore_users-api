package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
	"fmt"
)

// entry business logic in services

// functions to handle a get user rqeuest
// attepmt to get a user from the database
// points to the memory location of the struct created
//func CreateUser(user *users.User)

// *users.User means im returning a variable that stoed an addres of another variable
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	fmt.Printf("\n1 %p", &user)

	if err := user.Validate(); err != nil {
		return nil, err
	}
	fmt.Printf("\n3 %p", &user)
	if err := user.Save(); err != nil {
		return nil, err
	}
	fmt.Printf("\n4 %p", &user)
	return &user, nil
};

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := users.User{
		ID: userID,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
