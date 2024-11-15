package auth

import "regexp"

type registerRequest struct {
	Email string `json:"email"`
}

func (r *registerRequest) Validate() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(r.Email) {
		return ErrInvalidEmail
	}

	return nil
}
