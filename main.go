package main

import (
	"fmt"
	"github.com/baiv84/personio/handler"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}
	host := os.Getenv("host")
	dbname := os.Getenv("dbname")
	user := os.Getenv("user")
	password := os.Getenv("password")
	port := os.Getenv("port")
	const formatStr = "host=%s dbname=%s user=%s password=%s port=%s sslmode=disable"
	dsn := fmt.Sprintf(formatStr, host, dbname, user, password, port)

	crud := new(Crud)
	crud.InitDBEngine(dsn)
	defer crud.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	registerRoutes(r, crud)

	err = http.ListenAndServe(":3000", r)
	if err != nil {

		return
	}

}
