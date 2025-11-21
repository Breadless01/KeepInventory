package application

import (
	"KeepInventory/internal/domain"
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

func (s *ProjektService) ListProjekteByKunde(kunde string) ([]*domain.Projekt, error) {
	list, err := s.projRepo.FindByKunde(kunde)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Projekt{}
	}
	return list, nil
}
