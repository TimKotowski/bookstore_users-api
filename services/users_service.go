package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/crypto_utils"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
)

var (
	//A variable of an interface type can hold a value of a type that implements the interface
	UserService uersServiceInterface = &usersService{}
)

type uersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) (*users.User, *errors.RestErr)
	Search(string) (users.Users, *errors.RestErr)
}

type usersService struct {}

// entry business logic in services
// functions to handle a get user rqeuest
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Go automatically handles conversion between values and pointers for method calls
	// so no need to do &user
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService)  GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}

	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService)  UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := UserService.GetUser(user.ID)
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

func (s *usersService)  DeleteUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Delete(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *usersService)  Search(status string) (users.Users, *errors.RestErr) {
		dao := &users.User{}
		return dao.FindByStatus(status)
}
