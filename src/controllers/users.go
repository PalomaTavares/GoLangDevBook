package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// create and insert on data base
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("registration"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

// get all users from database according to name and nick filter
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	nameORNick := strings.ToLower(r.URL.Query().Get("user"))
	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, error := repository.Get(nameORNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// one specific user form database
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, error := repository.GetByID(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("is not possible to update this user"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("edit"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Update(userID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("is not possible to delete this user"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Delete(userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot follow yourself"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Follow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot unfollow yourself"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Unfollow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// gets all followers from a user
func GetAllFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	followers, error := repository.GetAllFollowers(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// get all users a user is following
func GetAllFollowing(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, error := repository.GetAllFollowing(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	responses.JSON(w, http.StatusOK, users)

}
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	idToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	if idToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("not possible to update the password of a different user"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)

	var password models.Password
	if error = json.Unmarshal(requestBody, &password); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()
	repository := repositories.NewUserRepository(db)
	passwordDB, error := repository.GetPassword(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(passwordDB, password.Current); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	passwordWHash, error := security.Hash(password.New)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.UpdatePassword(userID, string(passwordWHash)); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
