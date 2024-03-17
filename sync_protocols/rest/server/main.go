package main

import (
	"main/controllers"

	migrations "main/pkg/migrations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Migrations
	migrations.DatabaseMigration()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/posts", controllers.CreatePost)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
