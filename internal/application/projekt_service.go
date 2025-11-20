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

func (s *ProjektService) CreateProjekt(name string, kundeID int64) (*domain.Projekt, error) {
	if _, err := s.kundeRepo.FindByID(kundeID); err != nil {
		return nil, err
	}

	p := &domain.Projekt{
		Name:    name,
		KundeID: kundeID,
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

func (s *ProjektService) ListProjekteByKunde(kundeID int64) ([]*domain.Projekt, error) {
	list, err := s.projRepo.FindByKundeID(kundeID)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Projekt{}
	}
	return list, nil
}
