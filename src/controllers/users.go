package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// create and insert on data base
func CrateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}
	db, error := db.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUserRepository(db)
	repository.Create(user)
}

// get all users from database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

// one specific user form database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting user"))
}

// update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user"))
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user"))
}
