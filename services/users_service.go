package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
)

// entry business logic in services

// functions to handle a get user rqeuest
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
	user := users.User{ID: userID}

	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}


func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
		if err != nil {
			return nil, err
		}

			if isPartial {
				if user.Firstname != "" {
					current.Firstname = user.Firstname
				}
				if user.Lastname != "" {
					current.Lastname = user.Lastname
				}
				if user.Email != "" {
					current.Email = user.Email
				}

			} else {
				current.Firstname = user.Firstname
				current.Lastname = user.Lastname
				current.Email = user.Email
			}


		if err := current.Update(); err != nil {
			return nil, err
		}
		return current, nil
}


func DeleteUser(userID int64) (*users.User, *errors.RestErr) {
		user := users.User{ID: userID}
		if err := user.Delete(); err != nil {
			return nil ,err
		}
		return nil, nil
}
