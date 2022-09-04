package data

import (
	"explore/mongodb/features/checklist"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Checklist struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string             `bson:"user_id"`
	Name   string             `bson:"name"`
}

func fromCore(core checklist.Core) Checklist {
	return Checklist{
		ID:     primitive.NewObjectID(),
		UserID: core.UserID,
		Name:   core.Name,
	}
}

func (c *Checklist) toCore() checklist.Core {
	return checklist.Core{
		ID:     c.ID.Hex(),
		UserID: c.UserID,
		Name:   c.Name,
	}
}

func toCoreList(data []Checklist) []checklist.Core {
	result := []checklist.Core{}
	for ind := range data {
		result = append(result, data[ind].toCore())
	}
	return result
}
