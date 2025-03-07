package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// puts token in the request
func RequestWAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, error := http.NewRequest(method, url, data)
	if error != nil {
		return nil, error
	}

	cookie, _ := cookies.Read(r)

	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return nil, error
	}
	return response, nil
}
