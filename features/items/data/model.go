package data

import (
	"explore/mongodb/features/items"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID `bson:"_id"`
	ChecklistID string             `bson:"checklist_id"`
	Item        string             `bson:"item"`
	Status      string             `bson:"status"`
}

func fromCore(core items.Core) Item {
	return Item{
		ID:          primitive.NewObjectID(),
		ChecklistID: core.ChecklistID,
		Item:        core.Item,
		Status:      core.Status,
	}
}

func (i Item) toCore() items.Core {
	return items.Core{
		ID:          i.ID.Hex(),
		ChecklistID: i.ChecklistID,
		Item:        i.Item,
		Status:      i.Status,
	}
}

func toCOreList(data []Item) []items.Core {
	result := []items.Core{}
	for ind := range data {
		result = append(result, data[ind].toCore())
	}
	return result
}
