package api

import (
	"github.com/labstack/echo/v4"
	"github.com/ng-life/echo-web/db"
	"github.com/ng-life/echo-web/entity"
	"net/http"
)

func CreateUser(c echo.Context) error {
	repo := db.GetDB()
	var user entity.User
	user.Name = "Lifeng: " + c.RealIP()
	repo.Create(&user)
	return c.JSON(http.StatusOK, &user)
}
