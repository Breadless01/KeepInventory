package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"KeepInventory/internal/domain"
)

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

	for id := range b.LieferantIds {
		_, err := r.db.Exec(`
			INSERT INTO lieferant_bauteil (
			   lieferant_id = ?,
			   bauteil_id = ?,
			)
			`, id, b.ID,
		)

		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

type LieferantBauteil struct {
	LieferantId int64
	BauteilId   int64
}

func (r *BauteilRepositorySQLite) Update(b *domain.Bauteil) (*domain.Bauteil, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// 1. Bauteil updaten
	_, err = tx.Exec(`
        UPDATE bauteile
        SET teil_name = ?,
            kunde_id = ?,
            projekt_id = ?
        WHERE id = ?
    `, b.TeilName, b.KundeId, b.ProjektId, b.ID)
	if err != nil {
		return nil, err
	}

	// 2. Bisherige Zuordnungen für dieses Bauteil holen
	if len(b.LieferantIds) == 0 {
		// Wenn keine Lieferanten mehr: alle bisherigen löschen
		_, err = tx.Exec(`DELETE FROM lieferant_bauteil WHERE bauteil_id = ?`, b.ID)
		if err != nil {
			return nil, err
		}
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		return b, nil
	}

	rows, err := tx.Query(`
        SELECT lieferant_id, bauteil_id
        FROM lieferant_bauteil
        WHERE bauteil_id = ?`, b.ID)
	if err != nil {
		return nil, err
	}

	var junctions []LieferantBauteil
	for rows.Next() {
		var lb LieferantBauteil
		if err := rows.Scan(&lb.LieferantId, &lb.BauteilId); err != nil {
			rows.Close()
			return nil, err
		}
		junctions = append(junctions, lb)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()

	// 3. Berechnen: was einfügen, was löschen
	var toInsert []int64
	remaining := make([]LieferantBauteil, len(junctions))
	copy(remaining, junctions)

	for _, lieferantID := range b.LieferantIds {
		idx := findIndex(remaining, LieferantBauteil{LieferantId: lieferantID})
		if idx == -1 {
			toInsert = append(toInsert, lieferantID)
		}
		remaining = slices.DeleteFunc(junctions, func(lieferantBauteil LieferantBauteil) bool {
			return lieferantBauteil.LieferantId == lieferantID
		})
	}

	// remaining = was gelöscht werden muss
	for _, l := range remaining {
		_, err = tx.Exec(
			`DELETE FROM lieferant_bauteil WHERE bauteil_id = ? AND lieferant_id = ?`,
			b.ID, l.LieferantId,
		)
		if err != nil {
			return nil, err
		}
	}

	for _, lieferantID := range toInsert {
		_, err = tx.Exec(`
            INSERT INTO lieferant_bauteil (lieferant_id, bauteil_id)
            VALUES (?, ?)
        `, lieferantID, b.ID)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

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
			json_group_array(l.id) AS lieferanten_ids,
			json_group_array(l.name) AS lieferanten_namen,
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
        	LEFT JOIN
        	lieferant_bauteil lb ON lb.bauteil_id = b.id
        	LEFT JOIN
			lieferanten l ON l.id = lb.lieferant_id
        GROUP BY b.id
    `
	where, args := buildWhereClause(filter.Filters, domain.ResourceBauteile)
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
		var lieferantenJSON sql.NullString
		var lieferantenIDsJSON sql.NullString
		if err := rows.Scan(
			&b.ID,
			&b.TeilName,
			&kundeID,
			&kundeName,
			&projektID,
			&projektName,
			&lieferantenIDsJSON,
			&lieferantenJSON,
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

		if lieferantenJSON.Valid && lieferantenJSON.String != "" {
			if err := json.Unmarshal([]byte(lieferantenJSON.String), &b.Lieferanten); err != nil {
				return nil, fmt.Errorf("parse lieferanten json: %w", err)
			}
		}

		// JSON-Array der Lieferanten-IDs parsen
		if lieferantenIDsJSON.Valid && lieferantenIDsJSON.String != "" {
			// z.B. '[1,2,3]'
			if err := json.Unmarshal([]byte(lieferantenIDsJSON.String), &b.LieferantIds); err != nil {
				return nil, fmt.Errorf("parse lieferanten_ids json: %w", err)
			}
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

func buildWhereClause(filters map[string][]any, objType domain.FilterResource) (string, []any) {
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
			switch objType {
			case domain.ResourceBauteile:
				field = "b." + field
			}
		}
		parts = append(parts, fmt.Sprintf("%s IN (%s)", field, strings.Join(placeholders, ",")))
	}

	if len(parts) == 0 {
		return "", args
	}

	return "WHERE " + strings.Join(parts, " AND "), args
}

func findIndex(slice []LieferantBauteil, target LieferantBauteil) int {
	for i, v := range slice {
		if v.LieferantId == target.LieferantId {
			return i
		}
	}
	return -1
}
