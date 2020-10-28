package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	"fmt"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = ("INSERT INTO users (first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)")
)

// only databse goes in this file to interact with the database
// going to be teansfering from the persitence layer ot the application
// interact with the database
// data access layer to our database

// get by ID primary key
//  not going ot have udpated fields because the GET method is a copy
//  so the methods need to be a pointer to make sure were modifying the actual object the memory postion of that user
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	return nil
}

// save this user in the database
// not going to have udpated fields because the GET method is a copy
// so the methods need to be a pointer to make sure were modifying the actual object the memory postion of that user
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.Firstname, user.Lastname, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.ID = userID
	return nil
}
