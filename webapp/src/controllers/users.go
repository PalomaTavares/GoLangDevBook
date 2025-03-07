package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user map[string]string
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Debugging: Log the received data
	log.Println("Received user data:", user)

	userJSON, error := json.Marshal(user)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.ApiUrl)
	response, error := http.Post(url, "application/json", bytes.NewBuffer(userJSON))
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.ApiUrl, userID)

	response, error := requests.RequestWAuthentication(r, http.MethodPost, url, nil)

	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		fmt.Println(error)
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiUrl, userID)

	response, error := requests.RequestWAuthentication(r, http.MethodPost, url, nil)

	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		fmt.Println(error)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a map
	var user map[string]string
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding JSON:", err)
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: "Invalid JSON request"})
		return
	}

	// Debugging: Log the received data
	log.Println("Received user data:", user)

	// Marshal the user map back into JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling user data:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: err.Error()})
		return
	}

	// Read the user ID from the cookie
	cookie, err := cookies.Read(r)
	if err != nil {
		log.Println("Error reading cookie:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Failed to read cookie"})
		return
	}

	usuarioID, err := strconv.ParseUint(cookie["id"], 10, 64)
	if err != nil {
		log.Println("Error parsing user ID from cookie:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Failed to parse user ID"})
		return
	}

	// Construct the URL for the API request
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, usuarioID)

	// Make the PUT request with authentication
	response, err := requests.RequestWAuthentication(r, http.MethodPut, url, bytes.NewBuffer(userJSON))
	if err != nil {
		log.Println("Error making API request:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: err.Error()})
		return
	}
	defer response.Body.Close()

	// Handle error status codes
	if response.StatusCode >= 400 {
		log.Println("API returned an error status code:", response.StatusCode)
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	// Return the response to the client
	responses.JSON(w, response.StatusCode, nil)
}
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var passwords map[string]string
	if err := json.NewDecoder(r.Body).Decode(&passwords); err != nil {
		log.Println("Error decoding JSON:", err)
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: "Invalid JSON request"})
		return
	}

	// Debugging: Log the received data
	log.Println("Received user data:", passwords)

	// Marshal the user map back into JSON
	passwordsJSON, err := json.Marshal(passwords)
	if err != nil {
		log.Println("Error marshaling user data:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: err.Error()})
		return
	}

	// Read the user ID from the cookie
	cookie, err := cookies.Read(r)
	if err != nil {
		log.Println("Error reading cookie:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Failed to read cookie"})
		return
	}

	userID, err := strconv.ParseUint(cookie["id"], 10, 64)
	if err != nil {
		log.Println("Error parsing user ID from cookie:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Failed to parse user ID"})
		return
	}

	// Construct the URL for the API request
	url := fmt.Sprintf("%s/users/%d/updatePassword", config.ApiUrl, userID)

	// Make the PUT request with authentication
	response, err := requests.RequestWAuthentication(r, http.MethodPost, url, bytes.NewBuffer(passwordsJSON))
	if err != nil {
		log.Println("Error making API request:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: err.Error()})
		return
	}
	defer response.Body.Close()

	// Handle error status codes
	if response.StatusCode >= 400 {
		log.Println("API returned an error status code:", response.StatusCode)
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	// Return the response to the client
	responses.JSON(w, response.StatusCode, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		log.Println("Error reading cookie:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Failed to read cookie"})
		return
	}

	userID, err := strconv.ParseUint(cookie["id"], 10, 64)
	if err != nil {
		log.Println("Error parsing user ID from cookie:", err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Failed to parse user ID"})
		return
	}

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userID)
	response, error := requests.RequestWAuthentication(r, http.MethodDelete, url, nil)

	if error != nil {
		log.Println("Error making API request:", error)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	defer response.Body.Close()

	// Handle error status codes
	if response.StatusCode >= 400 {
		log.Println("API returned an error status code:", response.StatusCode)
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	// Return the response to the client
	responses.JSON(w, response.StatusCode, nil)
}
