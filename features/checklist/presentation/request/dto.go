package request

import "explore/mongodb/features/checklist"

type Checklist struct {
	Name   string
	UserID string
	Status string
}

func (c Checklist) ToCore() checklist.Core {
	return checklist.Core{
		UserID: c.UserID,
		Name:   c.Name,
	}
}
