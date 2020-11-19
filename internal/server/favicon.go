package server

import (
	"io/ioutil"
	"net/http"
)

// favicon writes the favicon bytes to the response.
func favicon(w http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadFile("assets/favicon.png")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(bs)
}
