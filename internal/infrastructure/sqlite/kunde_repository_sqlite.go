package sqlite

import (
	"database/sql"

	"KeepInventory/internal/domain"
)

type KundeRepositorySQLite struct {
	db *sql.DB
}

func NewKundeRepositorySQLite(db *sql.DB) *KundeRepositorySQLite {
	return &KundeRepositorySQLite{db: db}
}

func (r *KundeRepositorySQLite) Create(k *domain.Kunde) (*domain.Kunde, error) {
	res, err := r.db.Exec(
		`INSERT INTO kunden (name, sitz) VALUES (?, ?)`,
		k.Name, k.Sitz,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	k.ID = id
	return k, nil
}

func (r *KundeRepositorySQLite) FindByID(id int64) (*domain.Kunde, error) {
	row := r.db.QueryRow(
		`SELECT id, name, sitz FROM kunden WHERE id = ?`,
		id,
	)
	var k domain.Kunde
	if err := row.Scan(&k.ID, &k.Name, &k.Sitz); err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *KundeRepositorySQLite) FindAll() ([]*domain.Kunde, error) {
	rows, err := r.db.Query(
		`SELECT id, name, sitz FROM kunden ORDER BY name ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kunden := make([]*domain.Kunde, 0)
	for rows.Next() {
		var k domain.Kunde
		if err := rows.Scan(&k.ID, &k.Name, &k.Sitz); err != nil {
			return nil, err
		}
		kunden = append(kunden, &k)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return kunden, nil
}
