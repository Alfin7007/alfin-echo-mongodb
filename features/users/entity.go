package users

type Core struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type Data interface {
	InsertData(Core) error
}

type Bussiness interface {
	Register(Core) error
}
