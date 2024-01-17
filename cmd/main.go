package main

import (
	"context"
	"fmt"

	"github.com/cgradwohl/goth-stack/handler"
	"github.com/labstack/echo/v4"
)

// mock

func main() {
	app := echo.New()

	// the following will serve any file from the assets directory for path /*.
	// For example, a request to /css/pico.css will fetch and serve assets/css/pico.css file.
	app.Static("/", "assets")

	// note: here is how we can pass a DB instance to the handler as dependency injection
	// type DB struct{}
	// var dB DB
	// userHandler := handler.UserHandler{db}

	userHandler := handler.UserHandler{}
	app.Use(withUser)
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")

	fmt.Println("hello creature ...")
}

// custom middleware function in echo
// here I am mocking authorization by user email
func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get user from db
		// user := db.FindUser()
		ctx := context.WithValue(c.Request().Context(), "user", "c@gg.com from  middleware context")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
