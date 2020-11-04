package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	"bookstore_users-api/utils/mysql_utils"
	"fmt"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser  = "DELETE FROM users WHERE id=?"
)

// only databse goes in this file to interact with the database
// going to be teansfering from the persitence layer ot the application
// interact with the database
// data access layer to our database

// get by ID primary key
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

// save this user in the database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.Firstname, user.Lastname, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Firstname, user.Lastname, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err !=nil {
		return mysql_utils.ParseError(err)
	}
	fmt.Println("DELETED USER")
	return nil
}
