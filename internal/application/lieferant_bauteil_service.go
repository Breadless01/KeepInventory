package application

import "KeepInventory/internal/domain"

type LieferantBauteilService struct {
	repo LieferantBauteilRepository
}

type LieferantBauteilInput struct {
	ID          int64
	LieferantId int64
	BauteilId   int64
}

func NewLieferantBauteilService(repo LieferantBauteilRepository) *LieferantBauteilService {
	return &LieferantBauteilService{repo: repo}
}

func (s *LieferantBauteilService) Create(in LieferantBauteilInput) error {
	lb := domain.LieferantBauteil{
		LiferantId: in.LieferantId,
		BauteilId:  in.BauteilId,
	}
	return s.repo.Create(&lb)
}

func (s *LieferantBauteilService) Delete(bauteilId int64, lieferantId int64) error {
	return s.repo.Delete(bauteilId, lieferantId)
}

func (s *LieferantBauteilService) FindByBauteilId(id int64) ([]*domain.LieferantBauteil, error) {
	res, err := s.repo.FindByBauteilId(id)
	if err != nil {
		return []*domain.LieferantBauteil{}, err
	}
	return res, nil
}

func (s *LieferantBauteilService) FindByLieferantId(id int64) ([]*domain.LieferantBauteil, error) {
	res, err := s.repo.FindByLieferantId(id)
	if err != nil {
		return []*domain.LieferantBauteil{}, err
	}
	return res, nil
}
