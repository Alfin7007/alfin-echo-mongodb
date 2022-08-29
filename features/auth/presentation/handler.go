package presentation

import (
	"explore/mongodb/features/auth"
	"explore/mongodb/features/auth/presentation/request"
	"explore/mongodb/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authBussiness auth.Bussiness
}

func NewAuthHandler(auth auth.Bussiness) *AuthHandler {
	return &AuthHandler{
		authBussiness: auth,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	authReq := request.Auth{}
	bindErr := c.Bind(&authReq)
	if bindErr != nil {
		return c.JSON(helper.BadRequest())
	}
	id, token, authErr := h.authBussiness.Login(authReq.ToCore())
	if authErr != nil {
		return c.JSON(helper.BadRequestWithMSG(authErr.Error()))
	}
	return c.JSON(helper.SuccessLogin(id, token))
}
