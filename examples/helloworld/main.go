package main

import (
	_ "github.com/arvaliullin/go-restful-service/examples/helloworld/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// @title Hello World API
// @version 1.0
// @description This is a simple Hello World API.
// @host localhost:1323
// @BasePath /

func main() {
	e := echo.New()

	// Приветственный маршрут
	e.GET("/", HelloWorld)

	// Маршрут для Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Start(":1323")
}

// HelloWorld godoc
// @Summary Echo Hello World
// @Description Responds with 'Hello, World!'
// @Tags Hello
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
