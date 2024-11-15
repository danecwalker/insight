package auth

type AuthError string

const (
	ErrInvalidEmail AuthError = "invalid email"
)

func (e AuthError) Error() string {
	return string(e)
}
