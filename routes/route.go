package routes

import (
	"explore/mongodb/factory"

	"github.com/labstack/echo/v4"
)

func InitRoute(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.POST("/users", presenter.UserPresenter.RegistrasiUser)
	e.GET("/users/:id", presenter.UserPresenter.GetUser)
	e.POST("/login", presenter.AuthPresenter.Login)
	return e
}
