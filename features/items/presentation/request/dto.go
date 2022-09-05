package request

import "explore/mongodb/features/items"

type Item struct {
	ChecklistID string `json:"checklist_id"`
	Item        string `json:"item"`
}

func (i Item) ToCore() items.Core {
	return items.Core{
		ChecklistID: i.ChecklistID,
		Item:        i.Item,
	}
}
