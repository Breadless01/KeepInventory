package application

import (
	"KeepInventory/internal/domain"
	"math/rand"
)

type LieferantService struct {
	repo LieferantRepository
}

type LieferantInput struct {
	ID   int64
	Name string
	Sitz string
}

func NewLieferantService(repo LieferantRepository) *LieferantService {
	return &LieferantService{repo: repo}
}

func (s *LieferantService) ListLieferanten() ([]*domain.Lieferant, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Lieferant{}
	}
	return list, nil
}

func (s *LieferantService) FacetFilter(req domain.FilterState) (domain.LieferantFilterResult, error) {
	lieferanten, err := s.repo.FindByFilter(req)
	if err != nil {
		return domain.LieferantFilterResult{}, err
	}

	facets := s.buildFacets(lieferanten, req)

	total := len(lieferanten)
	start := (req.Page - 1) * req.PageSize
	if start < 0 {
		start = 0
	}

	return domain.LieferantFilterResult{
		Items:  lieferanten,
		Total:  total,
		Facets: facets,
	}, nil
}

func (s *LieferantService) CreateLieferant(name, sitz string) (*domain.Lieferant, error) {
	l := &domain.Lieferant{
		Name: name,
		Sitz: sitz,
	}
	return s.repo.Create(l)
}

func (s *LieferantService) UpdateLieferant(in LieferantInput) (*domain.Lieferant, error) {
	l := &domain.Lieferant{
		ID:   in.ID,
		Name: in.Name,
		Sitz: in.Sitz,
	}
	return s.repo.Update(l)
}

func (s *LieferantService) buildFacets(lieferanten []*domain.Lieferant, req domain.FilterState) map[string][]domain.FacetOption {
	facets := make(map[string]map[string]int)

	for _, l := range lieferanten {
		incLieferantAttr(facets, "name", l.Name)
		incLieferantAttr(facets, "sitz", l.Sitz)
	}

	out := make(map[string][]domain.FacetOption)

	for field, m := range facets {
		var opts []domain.FacetOption
		for name, count := range m {
			if name != "" {
				opts = append(opts, domain.FacetOption{
					ID:    rand.Int63(),
					Name:  name,
					Count: count,
				})
			}

		}
		out[field] = opts
	}

	return out
}

func incLieferantAttr(m map[string]map[string]int, field string, value string) {
	if m[field] == nil {
		m[field] = make(map[string]int)
	}
	m[field][value]++
}
