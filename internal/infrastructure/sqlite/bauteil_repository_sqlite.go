package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"KeepInventory/internal/domain"
)

// BauteilRepositorySQLite implementiert BauteilRepository mit SQLite.
type BauteilRepositorySQLite struct {
	db *sql.DB
}

func NewBauteilRepositorySQLite(db *sql.DB) *BauteilRepositorySQLite {
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

func (r *BauteilRepositorySQLite) FindByFilter(filter domain.FilterState) ([]*domain.Bauteil, error) {
	base := `
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
    `
	where, args := buildWhereClause(filter.Filters)
	query := base + " " + where + " ORDER BY teil_name ASC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Bauteil

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
	return result, nil
}

type row struct {
	ID   int64
	Name string
}

func (r *BauteilRepositorySQLite) GetAttributeValuesById(facets map[string]map[int64]int) map[string]map[int64]string {
	valueMap := make(map[string]map[int64]string)

	for key, _ := range facets {
		tableName := strings.Split(key, "_")[0]
		if tableName == "projekt" {
			tableName = "projekte"
		} else if tableName == "kunde" {
			tableName = "kunden"
		}
		rows, err := r.db.Query(`
			SELECT id, name 
			FROM ` + tableName,
		)

		if err != nil {
			log.Panicln(err)
			return nil
		}
		defer rows.Close()

		values := make(map[int64]string)

		for rows.Next() {
			v := row{}

			if err := rows.Scan(&v.ID, &v.Name); err != nil {
				log.Panicln(err)
				return nil
			}
			values[v.ID] = v.Name
		}
		valueMap[key] = values
	}
	return valueMap
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

func buildWhereClause(filters map[string][]any) (string, []any) {
	if len(filters) == 0 {
		return "", nil
	}

	var parts []string
	var args []any

	for field, ids := range filters {
		if len(ids) == 0 {
			continue
		}
		placeholders := make([]string, len(ids))
		for i, id := range ids {
			placeholders[i] = "?"
			args = append(args, id)
		}
		if field == "id" {
			field = "b." + field
		}
		parts = append(parts, fmt.Sprintf("%s IN (%s)", field, strings.Join(placeholders, ",")))
	}

	if len(parts) == 0 {
		return "", args
	}

	return "WHERE " + strings.Join(parts, " AND "), args
}

func (r *BauteilRepositorySQLite) SearchSuggestions(prefix string, limit int) ([]domain.BauteilSuggestion, error) {
	q := strings.TrimSpace(prefix) + "*"

	if limit <= 0 {
		limit = 10
	}

	stmt := fmt.Sprintf("SELECT b.id, b.teil_name, b.sachnummer FROM bauteile_fts JOIN bauteile b ON b.id = bauteile_fts.rowid WHERE bauteile_fts MATCH '%s' ORDER BY rank", q)
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []domain.BauteilSuggestion
	for rows.Next() {
		var s domain.BauteilSuggestion
		if err := rows.Scan(&s.ID, &s.TeilName, &s.Sachnummer); err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return res, rows.Err()
}
