package magic

type MagicStore interface {
	Create(m *Magic) error
	GetByEmail(email string) (*Magic, error)
	DeleteByEmail(email string) error
}
