package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	codedData, error := s.Encode("data", data)
	if error != nil {
		return error
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    codedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, error := r.Cookie("data")
	if error != nil {
		return nil, error
	}

	values := make(map[string]string)
	if error = s.Decode("data", cookie.Value, &values); error != nil {
		return nil, error
	}

	return values, nil

}
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
