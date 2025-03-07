package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Debugging: Log the received data
	log.Println("Received post data:", post)

	postJSON, error := json.Marshal(post)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	response, error := requests.RequestWAuthentication(r, http.MethodPost, url, bytes.NewBuffer(postJSON))
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

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	url := fmt.Sprintf("%s/posts/%d/like", config.ApiUrl, postID)
	response, error := requests.RequestWAuthentication(r, http.MethodPut, url, nil)
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

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	url := fmt.Sprintf("%s/posts/%d/unlike", config.ApiUrl, postID)
	response, error := requests.RequestWAuthentication(r, http.MethodPut, url, nil)
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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	var post map[string]string
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Debugging: Log the received data
	log.Println("Received post data:", post)

	postJSON, error := json.Marshal(post)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postID)

	response, error := requests.RequestWAuthentication(r, http.MethodPut, url, bytes.NewBufferString(string(postJSON)))
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
func DeletePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postID)

	response, error := requests.RequestWAuthentication(r, http.MethodDelete, url, nil)
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
