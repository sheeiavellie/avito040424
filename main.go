package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sheeiavellie/avito040424/api"
	"github.com/sheeiavellie/avito040424/handlers"
	"github.com/sheeiavellie/avito040424/middlewares"
	"github.com/sheeiavellie/avito040424/repository"
	"github.com/sheeiavellie/avito040424/storage"
)

func main() {
	// env stuff
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	port := os.Getenv("PORT")
	psConnStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("PS_USER"),
		os.Getenv("PS_PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PS_PORT"),
		os.Getenv("PS_DB_NAME"),
	)

	// create base context
	ctx := context.Background()

	// initialize dependencies
	postgresStorage, err := storage.NewPostgresStorage(ctx, psConnStr)
	if err != nil {
		log.Fatalf("postgres can't be created: %s", err)
	}
	defer postgresStorage.Close()

	lruCache := storage.NewLRUCacheStorage(100)

	bannerRepo := repository.NewBannerRepository(postgresStorage, lruCache)

	// routing
	router := http.NewServeMux()

	router.HandleFunc(
		"GET /user_banner",
		middlewares.AuthorizeToken(
			handlers.HandleGetUserBanner(
				ctx,
				bannerRepo,
			),
			api.UserRole,
		),
	)
	router.HandleFunc(
		"GET /banner",
		middlewares.AuthorizeToken(
			handlers.HandleGetBanners(
				ctx,
				bannerRepo,
			),
			api.AdminRole,
		),
	)
	router.HandleFunc(
		"POST /banner",
		middlewares.AuthorizeToken(
			handlers.HandlePostBanner(
				ctx,
				bannerRepo,
			),
			api.AdminRole,
		),
	)
	router.HandleFunc("PATCH /banner/{id}", handlers.HandleTODO)
	router.HandleFunc("DELETE /banner/{id}", handlers.HandleTODO)

	// run server
	log.Printf("server is listening on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
