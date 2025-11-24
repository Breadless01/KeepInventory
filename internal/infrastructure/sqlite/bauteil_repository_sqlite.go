package sqlite

import (
	"database/sql"
	"log"
	"strconv"

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
	log.Println(b)
	res, err := r.db.Exec(`
        INSERT INTO bauteile (
            teil_name, kunde_id, projekt_id, erstelldatum,
            typ_id, herstellungsart_id, verschleissteil_id,
            funktion_id, material_id, oberflaechenbehandlung_id,
            farbe_id, reserve_id, sachnummer
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		b.TeilName,
		b.KundeId,
		b.ProjektId,
		b.Erstelldatum,
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
		log.Println(err)
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
			b.id,
			b.teil_name,
			b.kunde_id,
			k.name,
			b.projekt_id,
			p.name,
			b.erstelldatum,
			b.typ_id,
			b.herstellungsart_id,
			b.verschleissteil_id,
			b.funktion_id,
			b.material_id,
			b.oberflaechenbehandlung_id,
			b.farbe_id,
			b.reserve_id,
			b.sachnummer
		FROM bauteile b
			LEFT JOIN
			kunden k ON b.kunde_id = k.id
			LEFT JOIN
			projekte p ON b.projekt_id = p.id
		ORDER BY teil_name DESC;

    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*domain.Bauteil, 0)

	for rows.Next() {
		var b domain.Bauteil
		var kundeID sql.NullString
		var projektID sql.NullString
		var kundeName sql.NullString
		var projektName sql.NullString
		if err := rows.Scan(
			&b.ID,
			&b.TeilName,
			&kundeID,
			&kundeName,
			&projektID,
			&projektName,
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
		if kundeID.Valid && kundeID.String != "" {
			b.KundeId, err = strconv.ParseInt(kundeID.String, 10, 64)
			if err != nil {
				return nil, err
			}
		} else {
			b.KundeId = 0
		}

		if projektID.Valid && projektID.String != "" {
			b.ProjektId, err = strconv.ParseInt(projektID.String, 10, 64)
			if err != nil {
				return nil, err
			}
		} else {
			b.ProjektId = 0
		}

		if kundeName.Valid {
			b.Kunde = kundeName.String
		} else {
			b.Kunde = ""
		}

		if projektName.Valid {
			b.Projekt = projektName.String
		} else {
			b.Projekt = ""
		}
		result = append(result, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println(result)
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
