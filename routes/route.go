package routes

import (
	"explore/mongodb/factory"
	"explore/mongodb/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoute(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	middlewares.LogMiddleware(e)
	e.POST("/users", presenter.UserPresenter.RegistrasiUser)
	e.GET("/users", presenter.UserPresenter.GetUser, middlewares.JWTMIddleware())
	e.POST("/login", presenter.AuthPresenter.Login)

	e.POST("/checklist", presenter.ChecklistPresenter.CreateData, middlewares.JWTMIddleware())
	e.GET("/checklist", presenter.ChecklistPresenter.GetData, middlewares.JWTMIddleware())
	return e
}
