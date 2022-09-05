package checklist

type Core struct {
	ID     string
	UserID string
	Name   string
}

type Data interface {
	InsertData(Core) error
	FindData(userID string) ([]Core, error)
	DeleteData(userID, id string) error
}

type Bussiness interface {
	CreateData(Core) error
	GetData(userID string) ([]Core, error)
	DeleteChecklist(userID, id string) error
}
