package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// add new post o database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	post.IDAuthor = userID

	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	repository := repositories.NewPostRepository(db)
	post.ID, error = repository.Create(post)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// get all posts on user feed
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, error := repository.Get(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)

}

// get one post
func GetPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)

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

	repository := repositories.NewPostRepository(db)
	posts, error := repository.GetByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// updates post data
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)

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
	repository := repositories.NewPostRepository(db)
	postDB, error := repository.GetByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	if postDB.IDAuthor != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot update someone elses post"))
		return
	}
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.Update(postID, post); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// deletes post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)

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

	repository := repositories.NewPostRepository(db)
	postDB, error := repository.GetByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	if postDB.IDAuthor != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot delete someone elses post"))
		return
	}

	if error = repository.Delete(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
func GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewPostRepository(db)
	posts, error := repository.GetByUserID(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}
func Like(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
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

	repository := repositories.NewPostRepository(db)
	if error := repository.Like(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
func Unlike(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
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

	repository := repositories.NewPostRepository(db)
	if error := repository.Unlike(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
