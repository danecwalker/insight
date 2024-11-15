package magic

type MagicService interface {
	CreateMagic(email string) (string, error)
	GetMagicByEmail(email string) (Magic, error)
	DeleteMagicByEmail(email string) error
}

type MagicRepository interface {
	Save(magic Magic) error
	GetByEmail(email string) (Magic, error)
	DeleteByEmail(email string) error
}
