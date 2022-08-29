package presentation

import (
	"explore/mongodb/features/users"
	"explore/mongodb/features/users/presentation/request"
	"explore/mongodb/helper"

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
