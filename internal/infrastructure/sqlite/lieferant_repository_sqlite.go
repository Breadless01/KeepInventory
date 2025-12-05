package sqlite

import (
	"KeepInventory/internal/domain"
	"database/sql"
	"log"
)

type LieferantRepositorySQLite struct {
	db *sql.DB
}

func NewLieferantRepositorySQLite(db *sql.DB) *LieferantRepositorySQLite {
	return &LieferantRepositorySQLite{db: db}
}

func (r *LieferantRepositorySQLite) Create(l *domain.Lieferant) (*domain.Lieferant, error) {
	res, err := r.db.Exec(
		`INSERT INTO lieferanten (name, sitz) VALUES (?, ?)`,
		l.Name, l.Sitz,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	l.ID = id
	return l, nil
}

func (r *LieferantRepositorySQLite) Update(l *domain.Lieferant) (*domain.Lieferant, error) {
	res, err := r.db.Exec(`
		UPDATE lieferanten
		SET name = ?,
		    sitz = ?,
		WHERE id = ?		
	`, l.Name, l.Sitz, l.ID)
	log.Println(res, err)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (r *LieferantRepositorySQLite) FindById(id int64) (*domain.Lieferant, error) {
	row := r.db.QueryRow(
		`SELECT id, name, sitz FROM lieferanten WHERE id = ?`,
		id,
	)
	var l domain.Lieferant
	if err := row.Scan(&l.ID, &l.Name, &l.Sitz); err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *LieferantRepositorySQLite) FindAll() ([]*domain.Lieferant, error) {
	rows, err := r.db.Query(
		`SELECT id, name, sitz FROM lieferanten ORDER BY name ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lieferanten := make([]*domain.Lieferant, 0)
	for rows.Next() {
		var l domain.Lieferant
		if err := rows.Scan(&l.ID, &l.Name, &l.Sitz); err != nil {
			return nil, err
		}
		lieferanten = append(lieferanten, &l)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lieferanten, nil
}

func (r *LieferantRepositorySQLite) FindByFilter(filter domain.FilterState) ([]*domain.Lieferant, error) {
	base := `
        SELECT 
            id,
			name,
			sitz
		FROM lieferanten
    `
	where, args := buildWhereClause(filter.Filters, domain.ResourceLieferanten)
	query := base + " " + where + " ORDER BY name ASC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Lieferant

	for rows.Next() {
		var l domain.Lieferant
		if err := rows.Scan(
			&l.ID,
			&l.Name,
			&l.Sitz,
		); err != nil {
			return nil, err
		}

		result = append(result, &l)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
