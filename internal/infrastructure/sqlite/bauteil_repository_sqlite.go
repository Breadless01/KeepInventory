package sqlite

import (
	"database/sql"
	"time"

	"KeepInventory/internal/application"
	"KeepInventory/internal/domain"
)

// BauteilRepositorySQLite implementiert BauteilRepository mit SQLite.
type BauteilRepositorySQLite struct {
	db *sql.DB
}

func NewBauteilRepositorySQLite(db *sql.DB) application.BauteilRepository {
	return &BauteilRepositorySQLite{db: db}
}

func (r *BauteilRepositorySQLite) Create(b *domain.Bauteil) (*domain.Bauteil, error) {
	res, err := r.db.Exec(`
        INSERT INTO bauteile (
            teil_name, kunde_id, projekt_id, erstelldatum,
            typ_id, herstellungsart_id, verschleissteil_id,
            funktion_id, material_id, oberflaechenbehandlung_id,
            farbe_id, reserve_id, sachnummer
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		b.TeilName,
		b.KundeID,
		b.ProjektID,
		b.Erstelldatum.Format(time.DateTime),
		b.TypID,
		b.HerstellungsartID,
		b.VerschleissteilID,
		b.FunktionID,
		b.MaterialID,
		b.OberflaechenbehandlungID,
		b.FarbeID,
		b.ReserveID,
		b.Sachnummer,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	b.ID = id
	return b, nil
}

func (r *BauteilRepositorySQLite) FindAll() ([]*domain.Bauteil, error) {
	rows, err := r.db.Query(`
        SELECT
            id, teil_name, kunde_id, projekt_id, erstelldatum,
            typ_id, herstellungsart_id, verschleissteil_id,
            funktion_id, material_id, oberflaechenbehandlung_id,
            farbe_id, reserve_id, sachnummer
        FROM bauteile
        ORDER BY teil_name DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*domain.Bauteil, 0)

	for rows.Next() {
		var b domain.Bauteil
		if err := rows.Scan(
			&b.ID,
			&b.TeilName,
			&b.KundeID,
			&b.ProjektID,
			&b.Erstelldatum,
			&b.TypID,
			&b.HerstellungsartID,
			&b.VerschleissteilID,
			&b.FunktionID,
			&b.MaterialID,
			&b.OberflaechenbehandlungID,
			&b.FarbeID,
			&b.ReserveID,
			&b.Sachnummer,
		); err != nil {
			return nil, err
		}
		result = append(result, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BauteilRepositorySQLite) CountByAttributes(
	typID, herstellungsartID, verschleissteilID,
	funktionID, materialID, oberflaechenbehandlungID,
	farbeID, reserveID int64,
) (int64, error) {
	row := r.db.QueryRow(`
        SELECT COUNT(*) FROM bauteile
        WHERE typ_id = ?
          AND herstellungsart_id = ?
          AND verschleissteil_id = ?
          AND funktion_id = ?
          AND material_id = ?
          AND oberflaechenbehandlung_id = ?
          AND farbe_id = ?
          AND reserve_id = ?
    `,
		typID,
		herstellungsartID,
		verschleissteilID,
		funktionID,
		materialID,
		oberflaechenbehandlungID,
		farbeID,
		reserveID,
	)

	var count int64
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
