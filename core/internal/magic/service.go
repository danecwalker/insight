package magic

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type magicService struct {
	r MagicRepository
}

// CreateMagic implements MagicService.
func (m *magicService) CreateMagic(email string) (string, error) {
	code := randomCode()
	magic := Magic{
		email: email,
		code:  code,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": magic.email,
		"code":  magic.code,
		"exp":   jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	b64Token := base64.URLEncoding.EncodeToString([]byte(tokenString))
	return b64Token, nil
}

// DeleteMagicByEmail implements MagicService.
func (m *magicService) DeleteMagicByEmail(email string) error {
	panic("unimplemented")
}

// GetMagicByEmail implements MagicService.
func (m *magicService) GetMagicByEmail(email string) (Magic, error) {
	panic("unimplemented")
}

func randomCode() string {
	digits := "0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = digits[rand.Intn(len(digits))]
	}

	return string(code)
}

func NewMagicService(r MagicRepository) MagicService {
	return &magicService{r}
}
