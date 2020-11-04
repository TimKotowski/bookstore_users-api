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

func getUserID(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

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
		userID, idErr := getUserID(chi.URLParam(r, "user_id"))
		if idErr != nil {
			w.WriteHeader(idErr.Status)
			w.Write([]byte(fmt.Sprintf("%v", idErr)))
			return
		}
		result, saveErr := services.GetUser(userID)
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
		userID, idErr := getUserID(chi.URLParam(r, "user_id"))
		if idErr != nil {
			w.WriteHeader(idErr.Status)
			w.Write([]byte(fmt.Sprintf("%v", idErr)))
			return
		}

		var user users.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			restErr := errors.NewBadRequestError("invalid json")
			w.WriteHeader(restErr.Status)
			w.Write([]byte(fmt.Sprintf("%v", restErr)))
			return
		}

		user.ID = userID
		isPartial := r.Method == http.MethodPatch

		result, saveErr := services.UpdateUser(isPartial, user)
		if saveErr != nil {
			w.WriteHeader(saveErr.Status)
			jsonData, _ := json.Marshal(saveErr)
			w.Write(jsonData)
			return
		}
		jsonResult, _ := json.Marshal(result)
		w.Write(jsonResult)
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, idErr := getUserID(chi.URLParam(r, "user_id"))
		if idErr != nil {
			w.WriteHeader(idErr.Status)
			w.Write([]byte(fmt.Sprintf("%v", idErr)))
			return
		}
		_, saveErr := services.DeleteUser(userID)
		if saveErr != nil {
			w.WriteHeader(saveErr.Status)
			jsonErr, _ := json.Marshal(saveErr)
			w.Write(jsonErr)
		}
		fmt.Println(userID)
	}
}
