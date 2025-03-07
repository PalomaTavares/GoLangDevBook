package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.ExecuteTemplate(w, "login.html", nil)
}
func LoadNewAccountPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "create-user.html", nil)
}
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	var posts []models.Post
	if error = json.NewDecoder(response.Body).Decode(&posts); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}
func LoadUpdatePostPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postID)
	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	var post models.Post
	if error = json.NewDecoder(response.Body).Decode(&post); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	utils.ExecuteTemplate(w, "update-post.html", post)
}
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameORnick := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?user=%s", config.ApiUrl, nameORnick)

	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.DealWithErrorStatusCode(w, response)
		return
	}

	var users []models.User

	if error = json.NewDecoder(response.Body).Decode(&users); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}
	utils.ExecuteTemplate(w, "users.html", users)
}
func LoadUsersProfilePage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	UserLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == UserLoggedID {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, error := models.GetFullUser(userID, r)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedID int64
	}{
		User:         user,
		UserLoggedID: int64(UserLoggedID),
	})
}
func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	UserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, error := models.GetFullUser(UserID, r)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: error.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)

}
func LoadEditUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	UserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.GetUserData(channel, UserID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{ErrorAPI: "Error getting user"})
		return
	}
	utils.ExecuteTemplate(w, "edit-user.html", user)
}
func LoadUpdatePassword(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "update-password.html", nil)
}
