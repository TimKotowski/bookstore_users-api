package users

import (
	"bookstore_users-api/utils/errors"
	"fmt"
)

// only databse goes in this file to interact with the database
// going to be teansfering from the persitence layer ot the application
// interact with the database
// data access layer to our database

var (
	usersDB = make(map[int64]*User)
)


// get by ID primary key
// 	// not going ot have udpated fields because the GET method is a copy
// 	// so the methods need to be a pointer to make sure were modifying the actual object the memory postion of that user
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return  errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.Firstname = result.Firstname
	user.Lastname = result.Lastname
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return  nil
}

// save this user in the database
// not going to have udpated fields because the GET method is a copy
// so the methods need to be a pointer to make sure were modifying the actual object the memory postion of that user
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewNotFoundError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	 usersDB[user.ID] = user
	 return nil
}
