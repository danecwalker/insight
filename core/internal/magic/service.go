package magic

type magicService struct {
	r MagicRepository
}

// CreateMagic implements MagicService.
func (m *magicService) CreateMagic(email string) (string, error) {
	panic("unimplemented")
}

// DeleteMagicByEmail implements MagicService.
func (m *magicService) DeleteMagicByEmail(email string) error {
	panic("unimplemented")
}

// GetMagicByEmail implements MagicService.
func (m *magicService) GetMagicByEmail(email string) (Magic, error) {
	panic("unimplemented")
}

func NewMagicService(r MagicRepository) MagicService {
	return &magicService{r}
}
