package magic

import "database/sql"

type MagicSqlite struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *MagicSqlite {
	return &MagicSqlite{
		db: db,
	}
}

func (m *MagicSqlite) Save(magic Magic) error {
	_, err := m.db.Exec("INSERT INTO magic (email, code) VALUES (?, ?)", magic.email, magic.code)
	if err != nil {
		return err
	}

	return nil
}

func (m *MagicSqlite) GetByEmail(email string) (Magic, error) {
	row := m.db.QueryRow("SELECT * FROM magic WHERE email = ?", email)
	var magic Magic
	err := row.Scan(&magic.email, &magic.code)
	if err != nil {
		return Magic{}, err
	}

	return magic, nil
}

func (m *MagicSqlite) DeleteByEmail(email string) error {
	_, err := m.db.Exec("DELETE FROM magic WHERE email = ?", email)
	if err != nil {
		return err
	}

	return nil
}
