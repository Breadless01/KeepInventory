// internal/infrastructure/sqlite/stammdaten_repositories_sqlite.go
package sqlite

import (
	"database/sql"

	"KeepInventory/internal/domain"
)

/* ----------------------------------------------------------
   Typ
---------------------------------------------------------- */

type TypRepositorySQLite struct {
	db *sql.DB
}

func NewTypRepositorySQLite(db *sql.DB) *TypRepositorySQLite {
	return &TypRepositorySQLite{db: db}
}

func (r *TypRepositorySQLite) FindByID(id int64) (*domain.Typ, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM typ WHERE id = ?`, id)

	var t domain.Typ
	if err := row.Scan(&t.ID, &t.Name, &t.Symbol); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TypRepositorySQLite) FindAll() ([]*domain.Typ, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM typ ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Typ
	for rows.Next() {
		var t domain.Typ
		if err := rows.Scan(&t.ID, &t.Name, &t.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &t)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Herstellungsart
---------------------------------------------------------- */

type HerstellungsartRepositorySQLite struct {
	db *sql.DB
}

func NewHerstellungsartRepositorySQLite(db *sql.DB) *HerstellungsartRepositorySQLite {
	return &HerstellungsartRepositorySQLite{db: db}
}

func (r *HerstellungsartRepositorySQLite) FindByID(id int64) (*domain.Herstellungsart, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM herstellungsart WHERE id = ?`, id)

	var h domain.Herstellungsart
	if err := row.Scan(&h.ID, &h.Name, &h.Symbol); err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *HerstellungsartRepositorySQLite) FindAll() ([]*domain.Herstellungsart, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM herstellungsart ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Herstellungsart
	for rows.Next() {
		var h domain.Herstellungsart
		if err := rows.Scan(&h.ID, &h.Name, &h.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &h)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Verschleissteil
---------------------------------------------------------- */

type VerschleissteilRepositorySQLite struct {
	db *sql.DB
}

func NewVerschleissteilRepositorySQLite(db *sql.DB) *VerschleissteilRepositorySQLite {
	return &VerschleissteilRepositorySQLite{db: db}
}

func (r *VerschleissteilRepositorySQLite) FindByID(id int64) (*domain.Verschleissteil, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM verschleissteil WHERE id = ?`, id)

	var v domain.Verschleissteil
	if err := row.Scan(&v.ID, &v.Name, &v.Symbol); err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *VerschleissteilRepositorySQLite) FindAll() ([]*domain.Verschleissteil, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM verschleissteil ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Verschleissteil
	for rows.Next() {
		var v domain.Verschleissteil
		if err := rows.Scan(&v.ID, &v.Name, &v.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &v)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Funktion
---------------------------------------------------------- */

type FunktionRepositorySQLite struct {
	db *sql.DB
}

func NewFunktionRepositorySQLite(db *sql.DB) *FunktionRepositorySQLite {
	return &FunktionRepositorySQLite{db: db}
}

func (r *FunktionRepositorySQLite) FindByID(id int64) (*domain.Funktion, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM funktion WHERE id = ?`, id)

	var f domain.Funktion
	if err := row.Scan(&f.ID, &f.Name, &f.Symbol); err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FunktionRepositorySQLite) FindAll() ([]*domain.Funktion, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM funktion ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Funktion
	for rows.Next() {
		var f domain.Funktion
		if err := rows.Scan(&f.ID, &f.Name, &f.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &f)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Material
---------------------------------------------------------- */

type MaterialRepositorySQLite struct {
	db *sql.DB
}

func NewMaterialRepositorySQLite(db *sql.DB) *MaterialRepositorySQLite {
	return &MaterialRepositorySQLite{db: db}
}

func (r *MaterialRepositorySQLite) FindByID(id int64) (*domain.Material, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM material WHERE id = ?`, id)

	var m domain.Material
	if err := row.Scan(&m.ID, &m.Name, &m.Symbol); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MaterialRepositorySQLite) FindAll() ([]*domain.Material, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM material ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Material
	for rows.Next() {
		var m domain.Material
		if err := rows.Scan(&m.ID, &m.Name, &m.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &m)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Oberflächenbehandlung
---------------------------------------------------------- */

type OberflaechenbehandlungRepositorySQLite struct {
	db *sql.DB
}

func NewOberflaechenbehandlungRepositorySQLite(db *sql.DB) *OberflaechenbehandlungRepositorySQLite {
	return &OberflaechenbehandlungRepositorySQLite{db: db}
}

func (r *OberflaechenbehandlungRepositorySQLite) FindByID(id int64) (*domain.Oberflaechenbehandlung, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM oberflaechenbehandlung WHERE id = ?`, id)

	var o domain.Oberflaechenbehandlung
	if err := row.Scan(&o.ID, &o.Name, &o.Symbol); err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *OberflaechenbehandlungRepositorySQLite) FindAll() ([]*domain.Oberflaechenbehandlung, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM oberflaechenbehandlung ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Oberflaechenbehandlung
	for rows.Next() {
		var o domain.Oberflaechenbehandlung
		if err := rows.Scan(&o.ID, &o.Name, &o.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &o)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Farbe
---------------------------------------------------------- */

type FarbeRepositorySQLite struct {
	db *sql.DB
}

func NewFarbeRepositorySQLite(db *sql.DB) *FarbeRepositorySQLite {
	return &FarbeRepositorySQLite{db: db}
}

func (r *FarbeRepositorySQLite) FindByID(id int64) (*domain.Farbe, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM farbe WHERE id = ?`, id)

	var f domain.Farbe
	if err := row.Scan(&f.ID, &f.Name, &f.Symbol); err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FarbeRepositorySQLite) FindAll() ([]*domain.Farbe, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM farbe ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Farbe
	for rows.Next() {
		var f domain.Farbe
		if err := rows.Scan(&f.ID, &f.Name, &f.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &f)
	}
	return list, rows.Err()
}

/* ----------------------------------------------------------
   Reserve
---------------------------------------------------------- */

type ReserveRepositorySQLite struct {
	db *sql.DB
}

func NewReserveRepositorySQLite(db *sql.DB) *ReserveRepositorySQLite {
	return &ReserveRepositorySQLite{db: db}
}

func (r *ReserveRepositorySQLite) FindByID(id int64) (*domain.Reserve, error) {
	row := r.db.QueryRow(`SELECT id, name, symbol FROM reserve WHERE id = ?`, id)

	var res domain.Reserve
	if err := row.Scan(&res.ID, &res.Name, &res.Symbol); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ReserveRepositorySQLite) FindAll() ([]*domain.Reserve, error) {
	rows, err := r.db.Query(`SELECT id, name, symbol FROM reserve ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*domain.Reserve
	for rows.Next() {
		var res domain.Reserve
		if err := rows.Scan(&res.ID, &res.Name, &res.Symbol); err != nil {
			return nil, err
		}
		list = append(list, &res)
	}
	return list, rows.Err()
}
