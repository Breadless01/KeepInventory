package application

import (
	"KeepInventory/internal/domain"
	"math/rand"
)

type ProjektService struct {
	projRepo  ProjektRepository
	kundeRepo KundeRepository
}

func NewProjektService(projRepo ProjektRepository, kundeRepo KundeRepository) *ProjektService {
	return &ProjektService{
		projRepo:  projRepo,
		kundeRepo: kundeRepo,
	}
}

func (s *ProjektService) CreateProjekt(name string, kunde string) (*domain.Projekt, error) {
	p := &domain.Projekt{
		Name:  name,
		Kunde: kunde,
	}
	return s.projRepo.Create(p)
}

func (s *ProjektService) ListProjekte() ([]*domain.Projekt, error) {
	list, err := s.projRepo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Projekt{}
	}
	return list, nil
}

func (s *ProjektService) FacetFilter(req domain.FilterState) (domain.ProjektFilterResult, error) {
	projekte, err := s.projRepo.FindByFilter(req)
	if err != nil {
		return domain.ProjektFilterResult{}, err
	}

	facets := s.buildFacets(projekte, req)

	total := len(projekte)
	start := (req.Page - 1) * req.PageSize
	if start < 0 {
		start = 0
	}
	return domain.ProjektFilterResult{
		Items:  projekte,
		Total:  total,
		Facets: facets,
	}, nil
}

func (s *ProjektService) buildFacets(projekte []*domain.Projekt, req domain.FilterState) map[string][]domain.FacetOption {
	facets := make(map[string]map[string]int)

	for _, k := range projekte {
		incProjektAttr(facets, "name", k.Name)
		incProjektAttr(facets, "kunde", k.Kunde)
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

func incProjektAttr(m map[string]map[string]int, field string, value string) {
	if m[field] == nil {
		m[field] = make(map[string]int)
	}
	m[field][value]++
}
