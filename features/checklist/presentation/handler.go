package presentation

import (
	"explore/mongodb/features/checklist"
	"explore/mongodb/features/checklist/presentation/request"
	"explore/mongodb/helper"
	"explore/mongodb/middlewares"
	"log"

	"github.com/labstack/echo/v4"
)

type ChecklistBussiness struct {
	bussiness checklist.Bussiness
}

func ChecklistHandler(checkBussiness checklist.Bussiness) *ChecklistBussiness {
	return &ChecklistBussiness{
		bussiness: checkBussiness,
	}
}

func (h *ChecklistBussiness) CreateData(c echo.Context) error {

	userID, tokenErr := middlewares.ExtracToken(c)
	if tokenErr != nil {
		return c.JSON(helper.Forbidden())
	}

	reqBody := request.Checklist{}
	bindErr := c.Bind(&reqBody)
	if bindErr != nil {
		return c.JSON(helper.BadRequest())
	}
	reqBody.UserID = userID
	err := h.bussiness.CreateData(reqBody.ToCore())
	if err != nil {
		return c.JSON(helper.BadRequest())
	}
	return c.JSON(helper.SuccessInsert())
}

func (h *ChecklistBussiness) GetData(c echo.Context) error {
	userID, tokenErr := middlewares.ExtracToken(c)
	if tokenErr != nil {
		return c.JSON(helper.Forbidden())
	}
	result, err := h.bussiness.GetData(userID)
	if err != nil {
		log.Print(err)
		return c.JSON(helper.BadRequest())
	}
	return c.JSON(helper.SuccessGetData(result))
}
