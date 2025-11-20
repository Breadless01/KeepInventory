package application

import "KeepInventory/internal/domain"

type TypService struct {
	repo TypRepository
}

func NewTypService(repo TypRepository) *TypService {
	return &TypService{repo: repo}
}

func (s *TypService) FindAll() ([]*domain.Typ, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Typ{}
	}
	return list, nil
}

func (s *TypService) FindByID(id int64) (*domain.Typ, error) {
	return s.repo.FindByID(id)
}

type HerstellungsartService struct {
	repo HerstellungsartRepository
}

func NewHerstellungsartService(repo HerstellungsartRepository) *HerstellungsartService {
	return &HerstellungsartService{repo: repo}
}

func (s *HerstellungsartService) FindAll() ([]*domain.Herstellungsart, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Herstellungsart{}
	}
	return list, nil
}

func (s *HerstellungsartService) FindByID(id int64) (*domain.Herstellungsart, error) {
	return s.repo.FindByID(id)
}

type VerschleissteilService struct {
	repo VerschleissteilRepository
}

func NewVerschleissteilService(repo VerschleissteilRepository) *VerschleissteilService {
	return &VerschleissteilService{repo: repo}
}

func (s *VerschleissteilService) FindAll() ([]*domain.Verschleissteil, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Verschleissteil{}
	}
	return list, nil
}

func (s *VerschleissteilService) FindByID(id int64) (*domain.Verschleissteil, error) {
	return s.repo.FindByID(id)
}

type FunktionService struct {
	repo FunktionRepository
}

func NewFunktionService(repo FunktionRepository) *FunktionService {
	return &FunktionService{repo: repo}
}

func (s *FunktionService) FindAll() ([]*domain.Funktion, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Funktion{}
	}
	return list, nil
}

func (s *FunktionService) FindByID(id int64) (*domain.Funktion, error) {
	return s.repo.FindByID(id)
}

type MaterialService struct {
	repo MaterialRepository
}

func NewMaterialService(repo MaterialRepository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) FindAll() ([]*domain.Material, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Material{}
	}
	return list, nil
}

func (s *MaterialService) FindByID(id int64) (*domain.Material, error) {
	return s.repo.FindByID(id)
}

type OberflaechenbehandlungService struct {
	repo OberflaechenbehandlungRepository
}

func NewOberflaechenbehandlungService(repo OberflaechenbehandlungRepository) *OberflaechenbehandlungService {
	return &OberflaechenbehandlungService{repo: repo}
}

func (s *OberflaechenbehandlungService) FindAll() ([]*domain.Oberflaechenbehandlung, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Oberflaechenbehandlung{}
	}
	return list, nil
}

func (s *OberflaechenbehandlungService) FindByID(id int64) (*domain.Oberflaechenbehandlung, error) {
	return s.repo.FindByID(id)
}

type FarbeService struct {
	repo FarbeRepository
}

func NewFarbeService(repo FarbeRepository) *FarbeService {
	return &FarbeService{repo: repo}
}

func (s *FarbeService) FindAll() ([]*domain.Farbe, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Farbe{}
	}
	return list, nil
}

func (s *FarbeService) FindByID(id int64) (*domain.Farbe, error) {
	return s.repo.FindByID(id)
}

type ReserveService struct {
	repo ReserveRepository
}

func NewReserveService(repo ReserveRepository) *ReserveService {
	return &ReserveService{repo: repo}
}

func (s *ReserveService) FindAll() ([]*domain.Reserve, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Reserve{}
	}
	return list, nil
}

func (s *ReserveService) FindByID(id int64) (*domain.Reserve, error) {
	return s.repo.FindByID(id)
}
