package sqlite

import (
	"KeepInventory/internal/domain"
	"database/sql"

	"github.com/pkg/errors"
)

type LieferantBauteilRepositorySQLite struct {
	db *sql.DB
}

func NewLieferantBauteilRepositorySQLite(db *sql.DB) *LieferantBauteilRepositorySQLite {
	return &LieferantBauteilRepositorySQLite{db: db}
}

func (r *LieferantBauteilRepositorySQLite) Create(lb *domain.LieferantBauteil) error {
	_, err := r.db.Exec(
		`INSERT INTO lieferant_bauteil (lieferant_id, bauteil_id) VALUES (?, ?)`,
		lb.LiferantId, lb.BauteilId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *LieferantBauteilRepositorySQLite) Delete(bauteilId int64, lieferantId int64) error {
	if bauteilId != 0 && lieferantId != 0 {
		_, err := r.db.Exec(
			`		
					DELETE FROM lieferant_bauteil 
					WHERE 
					    lieferant_id = ?
					    and 
					    bauteil_id = ?
				`,
			lieferantId, bauteilId,
		)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("lieferant_id or bauteil_id is zero")
}

func (r *LieferantBauteilRepositorySQLite) FindByLieferantId(id int64) ([]*domain.LieferantBauteil, error) {
	var result []*domain.LieferantBauteil

	base := `
			SELECT 
				lb.id, 
				lb.bauteil_id, 
				lb.lieferant_id,
				l.name,
				b.teil_name
			FROM lieferant_bauteil
			LEFT JOIN
				lieferanten l ON lb.lieferant_id = l.id
			LEFT JOIN
				bauteile b ON lb.bauteil_id = b.id
			WHERE lb.lieferant_id = ?
		`
	rows, err := r.db.Query(base, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lb domain.LieferantBauteil
		if err := rows.Scan(
			&lb.ID,
			&lb.BauteilId,
			&lb.LiferantId,
			&lb.Lieferant,
			&lb.Bauteil,
		); err != nil {
			return nil, err
		}

		result = append(result, &lb)
	}

	return result, nil
}

func (r *LieferantBauteilRepositorySQLite) FindByBauteilId(id int64) ([]*domain.LieferantBauteil, error) {
	var result []*domain.LieferantBauteil

	base := `
			SELECT 
				lb.id, 
				lb.bauteil_id, 
				lb.lieferant_id,
				l.name,
				b.teil_name
			FROM lieferant_bauteil
			LEFT JOIN
				lieferanten l ON lb.lieferant_id = l.id
			LEFT JOIN
				bauteile b ON lb.bauteil_id = b.id
			WHERE lb.bauteil_id = ?
		`
	rows, err := r.db.Query(base, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lb domain.LieferantBauteil
		if err := rows.Scan(
			&lb.ID,
			&lb.BauteilId,
			&lb.LiferantId,
			&lb.Lieferant,
			&lb.Bauteil,
		); err != nil {
			return nil, err
		}

		result = append(result, &lb)
	}

	return result, nil
}
