package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sheeiavellie/avito040424/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	port := os.Getenv("PORT")

	router := http.NewServeMux()

	router.HandleFunc("GET /user_banner", handlers.HandleTODO)
	router.HandleFunc("GET /banner", handlers.HandleTODO)
	router.HandleFunc("POST /banner", handlers.HandleTODO)
	router.HandleFunc("PATCH /banner/{id}", handlers.HandleTODO)
	router.HandleFunc("DELETE /banner/{id}", handlers.HandleTODO)

	log.Printf("server is listening on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
