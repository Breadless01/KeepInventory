package application

import (
	"KeepInventory/internal/domain"
)

type KundeService struct {
	repo KundeRepository
}

func NewKundeService(repo KundeRepository) *KundeService {
	return &KundeService{repo: repo}
}

func (s *KundeService) CreateKunde(name, sitz string) (*domain.Kunde, error) {
	k := &domain.Kunde{
		Name: name,
		Sitz: sitz,
	}
	return s.repo.Create(k)
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
