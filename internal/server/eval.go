package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/orlade/arrai-browse/internal"
	"github.com/sirupsen/logrus"
)

// eval loads and evaluates the path, writing the output to the response.
func eval(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "*")
	logrus.Infof("evaluating path %s", path)

	out, err := internal.Eval(path, r.URL.Query()["args"]...)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("no path"))
		return
	}
	_, err = w.Write([]byte(out))
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("write failed"))
		return
	}
}
