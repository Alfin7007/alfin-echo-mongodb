package request

import checklist "explore/mongodb/features/checklists"

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
