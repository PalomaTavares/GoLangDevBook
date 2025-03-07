package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedIn time.Time `json:"createdIn"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func GetFullUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go GetUserData(userChannel, userID, r)
	go GetFollowers(followersChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	go GetPosts(postsChannel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("Error getting user")
			}
			user = userLoaded

		case followersLoaded := <-followersChannel:
			if followersLoaded == nil {
				return User{}, errors.New("Error getting followers")
			}
			followers = followersLoaded

		case followingLoaded := <-followingChannel:
			if followingLoaded == nil {
				return User{}, errors.New("Error getting users following")
			}
			following = followingLoaded

		case postsLoaded := <-postsChannel:
			if postsLoaded == nil {
				return User{}, errors.New("Error loading posts")
			}
			posts = postsLoaded
		}
	}
	user.Followers = followers
	user.Following = following
	user.Posts = posts
	return user, nil
}

func GetUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userID)
	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if error = json.NewDecoder(response.Body).Decode(&user); error != nil {
		channel <- User{}
		return
	}
	channel <- user
}

func GetFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userID)
	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if error = json.NewDecoder(response.Body).Decode(&followers); error != nil {
		channel <- nil
		return
	}
	if followers == nil {
		channel <- make([]User, 0) //slice vazio
		return
	}
	channel <- followers
}

func GetFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiUrl, userID)
	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if error = json.NewDecoder(response.Body).Decode(&following); error != nil {
		channel <- nil
		return
	}
	if following == nil {
		channel <- make([]User, 0) //slice vazio
		return
	}
	channel <- following
}

func GetPosts(channel chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.ApiUrl, userID)
	response, error := requests.RequestWAuthentication(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post

	if error = json.NewDecoder(response.Body).Decode(&posts); error != nil {
		channel <- nil
		return
	}
	if posts == nil {
		channel <- make([]Post, 0) //slice vazio
		return
	}
	channel <- posts

}
