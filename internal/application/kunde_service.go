package application

import (
	"KeepInventory/internal/domain"
	"math/rand"
)

type KundeService struct {
	repo KundeRepository
}

func NewKundeService(repo KundeRepository) *KundeService {
	return &KundeService{repo: repo}
}

func (s *KundeService) ListKunden() ([]*domain.Kunde, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Kunde{}
	}
	return list, nil
}

func (s *KundeService) FacetFilter(req domain.FilterState) (domain.KundeFilterResult, error) {
	kunden, err := s.repo.FindByFilter(req)
	if err != nil {
		return domain.KundeFilterResult{}, err
	}

	facets := s.buildFacets(kunden, req)

	total := len(kunden)
	start := (req.Page - 1) * req.PageSize
	if start < 0 {
		start = 0
	}
	end := start + req.PageSize
	if end > total {
		end = total
	}
	pageItems := kunden[start:end]

	return domain.KundeFilterResult{
		Items:  pageItems,
		Total:  total,
		Facets: facets,
	}, nil
}

func (s *KundeService) CreateKunde(name, sitz string) (*domain.Kunde, error) {
	k := &domain.Kunde{
		Name: name,
		Sitz: sitz,
	}
	return s.repo.Create(k)
}

func (s *KundeService) buildFacets(kunden []*domain.Kunde, req domain.FilterState) map[string][]domain.FacetOption {
	facets := make(map[string]map[string]int)

	for _, k := range kunden {
		incKundeAttr(facets, "name", k.Name)
		incKundeAttr(facets, "sitz", k.Sitz)
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

func incKundeAttr(m map[string]map[string]int, field string, value string) {
	if m[field] == nil {
		m[field] = make(map[string]int)
	}
	m[field][value]++
}
