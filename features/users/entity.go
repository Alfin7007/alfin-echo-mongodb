package users

type Core struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type Data interface {
	InsertData(Core) error
	GetData(id string) (Core, error)
}

type Bussiness interface {
	Register(Core) error
	GetUser(id string) (Core, error)
}
