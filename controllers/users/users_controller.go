package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
			w.WriteHeader(saveErr.Status)
			jsonData, _ := json.Marshal(saveErr)
			w.Write([]byte(jsonData))
			return
		}
		jsonResult, _ := json.Marshal(result)
		w.Write(jsonResult)
	}
}

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userID, userErr := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
		if userErr != nil {
			err := errors.NewBadRequestError("invalid user id")
			w.WriteHeader(err.Status)
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
		result, saveErr := services.GetUser(userID)
		fmt.Printf("\n1 %p", result)
		if saveErr != nil {
			w.WriteHeader(saveErr.Status)
			jsonData, _ := json.Marshal(saveErr)
			w.Write([]byte(jsonData))
			return
		}
		jsonData, _ := json.Marshal(result)
		w.Write(jsonData)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")


	}
}
