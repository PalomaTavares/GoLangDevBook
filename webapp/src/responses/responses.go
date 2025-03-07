package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	ErrorAPI string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if error := json.NewEncoder(w).Encode(data); error != nil {
			log.Fatal(error)
		}
	}

}

func DealWithErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var error ErrorAPI
	json.NewDecoder(r.Body).Decode(&error)
	JSON(w, r.StatusCode, error)
}
