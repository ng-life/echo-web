package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ng-life/echo-web/api"
	"github.com/ng-life/echo-web/db"
	"github.com/ng-life/echo-web/entity"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Decompress())
	e.Use(middleware.RequestID())
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))

	// Routes
	e.GET("/", api.CreateUser)

	repo := db.GetDB()
	err := repo.AutoMigrate(&entity.User{})
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
