package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// autheticates user
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	savedUSer, error := repository.GetByEmail(user.Email)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	if error = security.VerifyPassword(savedUSer.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	w.Write([]byte("Login successfull"))
}
