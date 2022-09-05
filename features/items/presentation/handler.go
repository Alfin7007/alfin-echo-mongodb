package presentation

import (
	"explore/mongodb/features/items"
	"explore/mongodb/features/items/presentation/request"
	"explore/mongodb/helper"
	"explore/mongodb/middlewares"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	bussiness items.Bussiness
}

func NewItemHandler(itemBussiness items.Bussiness) ItemHandler {
	return ItemHandler{
		bussiness: itemBussiness,
	}
}

func (h *ItemHandler) InsertItem(c echo.Context) error {
	_, tokenErr := middlewares.ExtracToken(c)
	if tokenErr != nil {
		return c.JSON(helper.Forbidden())
	}
	reqBody := request.Item{}
	bindErr := c.Bind(&reqBody)
	if bindErr != nil {
		return c.JSON(helper.BadRequest())
	}
	checklistID := c.Param("id")
	reqBody.ChecklistID = checklistID
	err := h.bussiness.InsertItem(reqBody.ToCore())
	if err != nil {
		return c.JSON(helper.BadRequest())
	}
	return c.JSON(helper.SuccessInsert())
}

func (h *ItemHandler) GetItems(c echo.Context) error {
	_, tokenErr := middlewares.ExtracToken(c)
	if tokenErr != nil {
		return c.JSON(helper.Forbidden())
	}
	checklistID := c.Param("id")
	result, err := h.bussiness.GetItems(checklistID)
	if err != nil {
		return c.JSON(helper.BadRequest())
	}
	return c.JSON(helper.SuccessGetData(result))
}
