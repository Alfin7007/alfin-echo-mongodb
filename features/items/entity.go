package items

type Core struct {
	ID          string
	ChecklistID string
	Item        string
	Status      string
}

type Data interface {
	InsertData(Core) error
	GetData(string) ([]Core, error)
	GetOneData(checklist, id string) (Core, error)
	DeleteOneData(checklist, id string) error
	UpdateOneData(Core) error
}

type Bussiness interface {
	InsertItem(Core) error
	GetItems(string) ([]Core, error)
	GetOneItem(checklist, id string) (Core, error)
	DeleteOneItem(checklist, id string) error
	UpdateOneItem(Core) error
}
