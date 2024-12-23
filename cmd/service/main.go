package main

import (
	"database/sql"
	_ "github.com/arvaliullin/go-restful-service/docs"
	"github.com/arvaliullin/go-restful-service/internal/glossary/delivery"
	"github.com/arvaliullin/go-restful-service/internal/glossary/repository"
	"github.com/arvaliullin/go-restful-service/internal/glossary/usecase"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	swagger "github.com/swaggo/echo-swagger"
	"log"
	"os"
)

func main() {
	databaseURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewGlossaryRepository(db)
	uc := usecase.NewGlossaryUsecase(repo)
	e := echo.New()

	e.GET("/swagger/*", swagger.WrapHandler)

	delivery.NewGlossaryHandler(e, uc)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
