package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/duckyone2"
	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	router     = chi.NewRouter()
	controller = duckyone2.Controller{}
)

func init() {
	// middlewares
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	// default response for http methods (405 if not specified)
	router.Post("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
}

func getRoutes() *chi.Mux {
	router.Post("/notify", func(w http.ResponseWriter, r *http.Request) {
		body, ok := utils.ParseJSON(r.Body)
		if !ok {
			return
		}
		meta := duckyone2.NotificationMeta{
			Event: fmt.Sprint(body["event"]),
			Mode:  fmt.Sprint(body["mode"]),
			Data:  body["data"],
		}
		log.Print(meta)
		controller.Execute(meta)
	})
	return router
}
