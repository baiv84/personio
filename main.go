package main

import (
	"github.com/baiv84/personio/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Crud = handler.DBCursor

func registerRoutes(r chi.Router, crud *Crud) {
	corsOpt := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsOpt.Handler)

	r.Route("/api/v1/person", func(r chi.Router) {
		r.Post("/", crud.Create)
		r.Get("/", crud.Read)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", crud.ReadOne)
			r.Patch("/", crud.Update)
			r.Delete("/", crud.Delete)
		})
	})

}

func main() {
	crud := new(Crud)
	crud.InitDBEngine()
	defer crud.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	registerRoutes(r, crud)

	err := http.ListenAndServe(":3000", r)
	if err != nil {

		return
	}

}
