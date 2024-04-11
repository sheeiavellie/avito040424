package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sheeiavellie/avito040424/api"
	"github.com/sheeiavellie/avito040424/handlers"
	"github.com/sheeiavellie/avito040424/middlewares"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	port := os.Getenv("PORT")

	router := http.NewServeMux()

	router.HandleFunc(
		"GET /user_banner",
		middlewares.AuthorizeToken(
			handlers.HandleGetUserBanner(
				"",
			),
			api.UserRole,
		),
	)
	router.HandleFunc("GET /banner", handlers.HandleGetBanner)
	router.HandleFunc("POST /banner", handlers.HandlePostBanner)
	router.HandleFunc("PATCH /banner/{id}", handlers.HandleTODO)
	router.HandleFunc("DELETE /banner/{id}", handlers.HandleTODO)

	log.Printf("server is listening on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
