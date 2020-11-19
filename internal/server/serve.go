package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Serve(host string, port int) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", dir)
	r.Get("/favicon.ico", favicon)
	r.Get("/*", eval)

	origin := fmt.Sprintf("%s:%d", host, port)
	logrus.Infof("listening on http://%s", origin)
	logrus.WithError(http.ListenAndServe(origin, r)).Fatalf("serve stopped")
}
