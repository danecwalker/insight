package magic

type Magic struct {
	email string
	code  string
}

func NewMagic(email string, code string) *Magic {
	return &Magic{
		email: email,
		code:  code,
	}
}
