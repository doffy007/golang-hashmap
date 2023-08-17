package routers

import (
	"net/http"

	"github.com/doffy007/golang-hashmap/internal/controller"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Register(r *mux.Router, lg *logrus.Logger) {

	controller := controller.NewShortenURLController()

	r.HandleFunc("/short-url", controller.ShortenURLHandler).Methods(http.MethodPost)
	r.HandleFunc("/long-url", controller.LongURLHandler).Methods(http.MethodPost)
}
