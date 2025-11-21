package sqlite

import (
	"database/sql"

	"KeepInventory/internal/domain"
)

type ProjektRepositorySQLite struct {
	db *sql.DB
}

func NewProjektRepositorySQLite(db *sql.DB) *ProjektRepositorySQLite {
	return &ProjektRepositorySQLite{db: db}
}

func (r *ProjektRepositorySQLite) Create(p *domain.Projekt) (*domain.Projekt, error) {
	res, err := r.db.Exec(
		`INSERT INTO projekte (name, kunde_id) VALUES (?, ?)`,
		p.Name, p.Kunde,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	p.ID = id
	return p, nil
}

func (r *ProjektRepositorySQLite) FindByID(id int64) (*domain.Projekt, error) {
	row := r.db.QueryRow(
		`SELECT id, name, kunde_id FROM projekte WHERE id = ?`,
		id,
	)
	var p domain.Projekt
	if err := row.Scan(&p.ID, &p.Name, &p.Kunde); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProjektRepositorySQLite) FindAll() ([]*domain.Projekt, error) {
	rows, err := r.db.Query(
		`SELECT id, name, kunde FROM projekte ORDER BY name ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projekte := make([]*domain.Projekt, 0)
	for rows.Next() {
		var p domain.Projekt
		if err := rows.Scan(&p.ID, &p.Name, &p.Kunde); err != nil {
			return nil, err
		}
		projekte = append(projekte, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projekte, nil
}

func (r *ProjektRepositorySQLite) FindByKunde(kunde string) ([]*domain.Projekt, error) {
	rows, err := r.db.Query(
		`SELECT id, name, kunde FROM projekte WHERE kunde = ? ORDER name id DESC`,
		kunde,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projekte := make([]*domain.Projekt, 0)
	for rows.Next() {
		var p domain.Projekt
		if err := rows.Scan(&p.ID, &p.Name, &p.Kunde); err != nil {
			return nil, err
		}
		projekte = append(projekte, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projekte, nil
}
