package bussiness

import (
	"explore/mongodb/features/items"
)

type itemBussiness struct {
	data items.Data
}

func NewItemBussiness(itemData items.Data) items.Bussiness {
	return &itemBussiness{
		data: itemData,
	}
}

func (uc *itemBussiness) InsertItem(core items.Core) error {
	core.Status = "uncheck"
	err := uc.data.InsertData(core)
	if err != nil {
		return err
	}
	return nil
}

func (uc *itemBussiness) GetItems(checklistID string) ([]items.Core, error) {
	result, err := uc.data.GetData(checklistID)
	if err != nil {
		return []items.Core{}, nil
	}
	return result, nil
}

func (uc *itemBussiness) GetOneItem(checklistID, id string) (items.Core, error) {
	result, err := uc.data.GetOneData(checklistID, id)
	if err != nil {
		return items.Core{}, err
	}
	return result, nil
}

func (uc *itemBussiness) DeleteOneItem(checklistID, id string) error {
	err := uc.data.DeleteOneData(checklistID, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *itemBussiness) UpdateOneItem(core items.Core) error {
	err := uc.data.UpdateOneData(core)
	if err != nil {
		return err
	}
	return nil
}
