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
}

type Bussiness interface {
	InsertItem(Core) error
	GetItems(string) ([]Core, error)
}
