package presentation

import (
	"explore/mongodb/features/users"
	"explore/mongodb/features/users/presentation/request"
	"explore/mongodb/features/users/presentation/response"
	"explore/mongodb/helper"
	"explore/mongodb/middlewares"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBussiness users.Bussiness
}

func NewUseHandler(bussiness users.Bussiness) *UserHandler {
	return &UserHandler{
		userBussiness: bussiness,
	}
}

func (h *UserHandler) RegistrasiUser(c echo.Context) error {
	reqBody := request.User{}
	bindErr := c.Bind(&reqBody)
	if bindErr != nil {
		return c.JSON(helper.BadRequest())
	}

	regErr := h.userBussiness.Register(reqBody.ToCore())
	if regErr != nil {
		return c.JSON(helper.BadRequest())
	}

	return c.JSON(helper.SuccessInsert())
}

func (h *UserHandler) GetUser(c echo.Context) error {
	userID, tokenErr := middlewares.ExtracToken(c)
	if tokenErr != nil {
		return c.JSON(helper.Forbidden())
	}
	userCore, getErr := h.userBussiness.GetUser(userID)
	if getErr != nil {
		return c.JSON(helper.BadRequest())
	}
	return c.JSON(helper.SuccessGetData(response.FromCore(userCore)))
}
