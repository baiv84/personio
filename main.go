package main

import (
	"fmt"
	"github.com/baiv84/personio/handler"
	"github.com/baiv84/personio/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var api *handler.DBCursor

func registerRoutes(r chi.Router) {
	corsOpt := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsOpt.Handler)

	r.Route("/api/v1/person", func(r chi.Router) {
		r.Post("/", api.Create)
		r.Get("/", api.Read)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", api.ReadOne)
			r.Patch("/", api.Update)
			r.Delete("/", api.Delete)
		})
	})

}

func init() {
	fmt.Println("Initializing...")
	var pgConn *gorm.DB
	const formatStr = "host=%s dbname=%s user=%s password=%s port=%s sslmode=disable"
	dsn := fmt.Sprintf(formatStr, "localhost", "citizens", "postgres", "12345", "5432")
	pgConn, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err := pgConn.AutoMigrate(&model.Person{})
	if err != nil {
		return
	}

	api = new(handler.DBCursor)
	api.Init(pgConn)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	registerRoutes(r)

	err := http.ListenAndServe(":3000", r)
	if err != nil {

		return
	}

}
