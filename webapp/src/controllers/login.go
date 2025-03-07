package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user map[string]string
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	userJSON, error := json.Marshal(user)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)

	response, error := http.Post(url, "application/json", bytes.NewBuffer(userJSON))
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	var authenticationData models.AuthenticationData
	if error = json.NewDecoder(response.Body).Decode(&authenticationData); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	if error = cookies.Save(w, authenticationData.ID, authenticationData.Token); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
