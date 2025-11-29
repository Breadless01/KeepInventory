package application

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"KeepInventory/internal/domain"
	"KeepInventory/internal/infrastructure/sqlite"
)

type FilterConfigService struct {
	db       *sql.DB
	filePath string
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
		// Datei existiert nicht -> Default bauen
		if os.IsNotExist(err) {
			cfg, err := s.buildDefaultConfig()
			if err != nil {
				return domain.FilterConfig{}, err
			}
			// direkt speichern
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

	// Beispiel: Bauteile
	cols, err := sqlite.ListColumns(s.db, "bauteile")
	if err != nil {
		return domain.FilterConfig{}, err
	}

	var fields []domain.FieldFilterConfig
	for _, c := range cols {
		// Zeug wie id, erstelldatum etc. optional rausfiltern
		if c.Name == "id" || c.Name == "erstelldatum" || c.Name == "sachnummer" {
			continue
		}
		fields = append(fields, domain.FieldFilterConfig{
			Field:   c.Name,
			Label:   labelFromColumn(c.Name), // z.B. "material_id" -> "Material"
			Enabled: false,                   // erst in Settings einschalten
		})
	}

	resources = append(resources, domain.ResourceFilterConfig{
		Resource: domain.ResourceBauteile,
		Table:    "bauteile",
		Fields:   fields,
	})

	return domain.FilterConfig{Resources: resources}, nil
}

func labelFromColumn(col string) string {
	// quick & dirty: typ_id -> Typ, material_id -> Material
	col = strings.TrimSuffix(col, "_id")
	col = strings.ReplaceAll(col, "_", " ")
	if len(col) == 0 {
		return col
	}
	return strings.ToUpper(col[:1]) + col[1:]
}
