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
	e.DELETE("/checklist/:id", presenter.ChecklistPresenter.DeleteCheclist, middlewares.JWTMIddleware())

	e.POST("/checklist/:checklistID/items", presenter.ItemPresenter.InsertItem, middlewares.JWTMIddleware())
	e.GET("/checklist/:checklistID/items", presenter.ItemPresenter.GetItems, middlewares.JWTMIddleware())
	e.GET("/checklist/:checklistID/items/:id", presenter.ItemPresenter.GetOneItem, middlewares.JWTMIddleware())
	e.DELETE("/checklist/:checklistID/items/:id", presenter.ItemPresenter.DeleteOneItem, middlewares.JWTMIddleware())
	e.PUT("/checklist/:checklistID/items/:id/status/:status", presenter.ItemPresenter.UpdateOneItem, middlewares.JWTMIddleware())
	return e
}
