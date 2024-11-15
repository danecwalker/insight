package users

type userService struct {
	r UserRepository
}

// CreateUser implements UserService.
func (u *userService) CreateUser(email string) (User, error) {
	panic("unimplemented")
}

// DeleteUserByEmail implements UserService.
func (u *userService) DeleteUserByEmail(email string) error {
	panic("unimplemented")
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (User, error) {
	panic("unimplemented")
}

func NewUserService(r UserRepository) UserService {
	return &userService{r}
}
