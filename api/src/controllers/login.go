package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	token, err := authentication.CreateToken(savedUSer.ID)
	if err != nil {
		fmt.Println("Erro ao criar token:", err)
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	userID := strconv.FormatUint(savedUSer.ID, 10)

	responses.JSON(w, http.StatusOK, models.AuthenticationData{ID: userID, Token: token})
}
