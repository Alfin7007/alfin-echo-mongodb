package auth

type Core struct {
	ID       string
	Email    string
	Password string
}

type Data interface {
	FindData(email string) (Core, error)
}

type Bussiness interface {
	Login(Core) (id, token string, err error)
}
