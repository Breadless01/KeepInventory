package sqlite

import "database/sql"

type ColumnInfo struct {
	Name string
	Type string
}

func ListColumns(db *sql.DB, table string) ([]ColumnInfo, error) {
	rows, err := db.Query(`PRAGMA table_info(` + table + `);`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cols []ColumnInfo
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dflt sql.NullString

		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk); err != nil {
			return nil, err
		}
		cols = append(cols, ColumnInfo{
			Name: name,
			Type: ctype,
		})
	}
	return cols, rows.Err()
}
