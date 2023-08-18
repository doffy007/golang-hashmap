package routers

import (
	"net/http"

	"github.com/doffy007/golang-hashmap/internal/controller"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Register(r *mux.Router, lg *logrus.Logger) {

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	controller := controller.NewShortenURLController()

	r.HandleFunc("/short-url", controller.ShortenURLHandler).Methods(http.MethodPost)
	r.HandleFunc("/long-url", controller.LongURLHandler).Methods(http.MethodPost)
}
