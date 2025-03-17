package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	controller "github.com/zk1569/test-mongodb/src/controllers"
)

func main() {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)

	mux.Route("/v1", func(r chi.Router) {
		controller.GetBankInstance().Mount(r)
	})

	srv := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running at 0.0.0.0:8080")
	srv.ListenAndServe()
}
