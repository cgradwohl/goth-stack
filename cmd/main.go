package main

import (
	"fmt"

	"github.com/cgradwohl/goth-stack/handler"
	"github.com/labstack/echo/v4"
)

// mock

func main() {
	app := echo.New()

	// note: here is how we can pass a DB instance to the handler as dependency injection
	// type DB struct{}
	// var dB DB
	// userHandler := handler.UserHandler{db}

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")

	fmt.Println("hello creature ...")
}
