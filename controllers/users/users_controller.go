package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"encoding/json"
	"fmt"
	"net/http"
)

// every reqeust is handled by controller
// entry point of application
// prodive the funcitonaly or the endpoints to interact against the users api

func SearchUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("implement Me!"))
		return
	}
}

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// user := &users.User{}
		var user users.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			w.WriteHeader(restErr.Status)
			w.Write([]byte(fmt.Sprintf("%v", restErr)))
			return
		}

		// pass the populated struct to the service funciton
		result, saveErr := services.CreateUser(user)
		if saveErr != nil {
			w.Write([]byte(fmt.Sprintf("%v %v", saveErr.Status, saveErr)))
		}
		fmt.Println(result)
	}
}

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("implement Me!"))
		return
	}
}
