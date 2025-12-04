package sqlite

import (
	"KeepInventory/internal/domain"
	"database/sql"
	"fmt"
	"strings"
)

type SearchRepositorySQLite struct {
	db      *sql.DB
	configs map[string]searchConfig
}

type searchConfig struct {
	FTSTable  string
	TableName string
	JoinSQL   string
	SelectSQL string
	WhereTpl  string
	OrderSQL  string
}

func NewSearchRepositorySQLite(db *sql.DB) *SearchRepositorySQLite {
	return &SearchRepositorySQLite{
		db: db,
		configs: map[string]searchConfig{
			"bauteil": {
				FTSTable:  "bauteile_fts",
				TableName: "bauteile",
				JoinSQL:   "JOIN bauteile b ON b.id = bauteile_fts.rowid",
				SelectSQL: `
					b.id              AS id,
					'bauteil'         AS type,
					b.teil_name       AS label,
					b.sachnummer      AS subtitle
				`,
				WhereTpl: `bauteile_fts MATCH ?`,
				OrderSQL: `ORDER BY rank`,
			},
			"kunde": {
				FTSTable:  "kunden_fts",
				TableName: "kunden",
				JoinSQL:   "JOIN kunden k ON k.id = kunden_fts.rowid",
				SelectSQL: `
					k.id         AS id,
					'kunde'      AS type,
					k.name       AS label,
					k.sitz       AS subtitle
				`,
				WhereTpl: `kunden_fts MATCH ?`,
				OrderSQL: `ORDER BY rank`,
			},
			"projekt": {
				FTSTable:  "projekte_fts",
				TableName: "projekte",
				JoinSQL:   "JOIN projekte p ON p.id = projekte_fts.rowid",
				SelectSQL: `
					p.id         AS id,
					'projekt'      AS type,
					p.name       AS label,
					p.kunde       AS subtitle
				`,
				WhereTpl: `projekte_fts MATCH ?`,
				OrderSQL: `ORDER BY rank`,
			},
		},
	}
}

func (r *SearchRepositorySQLite) Search(req domain.SearchRequest) ([]domain.SearchResult, error) {
	q := strings.TrimSpace(req.Query)
	if q == "" {
		return []domain.SearchResult{}, nil
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	var cfgs []searchConfig

	if req.ObjectType == "" {
		for _, cfg := range r.configs {
			cfgs = append(cfgs, cfg)
		}
	} else {
		if cfg, ok := r.configs[req.ObjectType]; ok {
			cfgs = append(cfgs, cfg)
		} else {
			return []domain.SearchResult{}, nil
		}
	}

	var results []domain.SearchResult

	for _, cfg := range cfgs {
		rows, err := r.runSearchQuery(cfg, q, req.Limit)
		if err != nil {
			return nil, err
		}
		results = append(results, rows...)
	}

	// später könntest du hier noch auf globale Limitierung, Ranking usw. normalisieren
	return results, nil
}

func (r *SearchRepositorySQLite) runSearchQuery(cfg searchConfig, q string, limit int) ([]domain.SearchResult, error) {
	var baseFrom string
	var args []interface{}

	if cfg.FTSTable != "" {
		baseFrom = fmt.Sprintf("FROM %s ", cfg.FTSTable)
		if cfg.JoinSQL != "" {
			baseFrom += cfg.JoinSQL + " "
		}
		match := q + "*"
		args = append(args, match)
	} else {
		baseFrom = fmt.Sprintf("FROM %s k ", cfg.TableName)
		pattern := "%" + q + "%"
		args = append(args, pattern, pattern)
	}

	query := fmt.Sprintf(`
		SELECT %s
		%s
		WHERE %s
		%s
		LIMIT ?;
	`, cfg.SelectSQL, baseFrom, cfg.WhereTpl, cfg.OrderSQL)

	args = append(args, limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []domain.SearchResult
	for rows.Next() {
		var r1 domain.SearchResult
		if err := rows.Scan(&r1.ID, &r1.Type, &r1.Label, &r1.Subtitle); err != nil {
			return nil, err
		}
		res = append(res, r1)
	}
	return res, rows.Err()
}
