package application

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"KeepInventory/internal/domain"
)

type FilterConfigService struct {
	db       *sql.DB
	filePath string
}

type columnInfo struct {
	Name string
	Type string
}

func NewFilterConfigService(db *sql.DB, baseDir string) *FilterConfigService {
	return &FilterConfigService{
		db:       db,
		filePath: filepath.Join(baseDir, "filter_config.json"),
	}
}

func (s *FilterConfigService) Load() (domain.FilterConfig, error) {
	b, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			cfg, err := s.buildDefaultConfig()
			if err != nil {
				return domain.FilterConfig{}, err
			}
			_ = s.Save(cfg)
			return cfg, nil
		}
		return domain.FilterConfig{}, err
	}

	var cfg domain.FilterConfig
	if err := json.Unmarshal(b, &cfg); err != nil {
		return domain.FilterConfig{}, err
	}
	return cfg, nil
}

func (s *FilterConfigService) Save(cfg domain.FilterConfig) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}

// Default: für Bauteile, Kunden, Projekte Spalten vorbefüllen
func (s *FilterConfigService) buildDefaultConfig() (domain.FilterConfig, error) {
	var resources []domain.ResourceFilterConfig

	tables := []domain.FilterResource{
		domain.ResourceBauteile,
		domain.ResourceKunden,
		domain.ResourceProjekte,
		domain.ResourceLieferanten,
	}
	for _, ressource := range tables {
		tableName := string(ressource)
		cols, err := listColumns(s.db, tableName)
		if err != nil {
			return domain.FilterConfig{}, err
		}

		var fields []domain.FieldFilterConfig
		for _, c := range cols {
			if c.Name == "id" || c.Name == "erstelldatum" || c.Name == "sachnummer" {
				continue
			}
			fields = append(fields, domain.FieldFilterConfig{
				Field:   c.Name,
				Label:   labelFromColumn(c.Name),
				Enabled: false,
			})
		}

		resources = append(resources, domain.ResourceFilterConfig{
			Resource: ressource,
			Table:    capitalizeFirstLetter(tableName),
			Fields:   fields,
		})
	}

	return domain.FilterConfig{Resources: resources}, nil
}

func listColumns(db *sql.DB, table string) ([]columnInfo, error) {
	rows, err := db.Query(`PRAGMA table_info(` + table + `);`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cols []columnInfo
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dflt sql.NullString

		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk); err != nil {
			return nil, err
		}
		cols = append(cols, columnInfo{
			Name: name,
			Type: ctype,
		})
	}
	return cols, rows.Err()
}

func labelFromColumn(col string) string {
	col = strings.TrimSuffix(col, "_id")
	col = strings.ReplaceAll(col, "_", " ")
	if len(col) == 0 {
		return col
	}
	return strings.ToUpper(col[:1]) + col[1:]
}

func capitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}

	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
