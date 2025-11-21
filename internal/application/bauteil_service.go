package application

import (
	"KeepInventory/internal/domain"
	"time"
)

// BauteilService kapselt Anwendungslogik rund um Bauteile.
type BauteilService struct {
	repo            BauteilRepository
	typRepo         TypRepository
	artRepo         HerstellungsartRepository
	verschRepo      VerschleissteilRepository
	funktionRepo    FunktionRepository
	materialRepo    MaterialRepository
	oberflaecheRepo OberflaechenbehandlungRepository
	farbeRepo       FarbeRepository
	reserveRepo     ReserveRepository
}

func NewBauteilService(repo BauteilRepository,
	typRepo TypRepository,
	artRepo HerstellungsartRepository,
	verschRepo VerschleissteilRepository,
	funktionRepo FunktionRepository,
	materialRepo MaterialRepository,
	oberflaecheRepo OberflaechenbehandlungRepository,
	farbeRepo FarbeRepository,
	reserveRepo ReserveRepository,
) *BauteilService {
	return &BauteilService{
		repo:            repo,
		typRepo:         typRepo,
		artRepo:         artRepo,
		verschRepo:      verschRepo,
		funktionRepo:    funktionRepo,
		materialRepo:    materialRepo,
		oberflaecheRepo: oberflaecheRepo,
		farbeRepo:       farbeRepo,
		reserveRepo:     reserveRepo,
	}
}

type CreateBauteilInput struct {
	TeilName  string
	KundeId   int64
	ProjektId int64

	TypID                    int64
	HerstellungsartID        int64
	VerschleissteilID        int64
	FunktionID               int64
	MaterialID               int64
	OberflaechenbehandlungID int64
	FarbeID                  int64
	ReserveID                int64
}

func (s *BauteilService) CreateBauteil(in CreateBauteilInput) (*domain.Bauteil, error) {
	b := &domain.Bauteil{
		TeilName:     in.TeilName,
		KundeId:      in.KundeId,
		ProjektId:    in.ProjektId,
		Erstelldatum: time.Now().Local().String(),

		TypID:                    in.TypID,
		HerstellungsartID:        in.HerstellungsartID,
		VerschleissteilID:        in.VerschleissteilID,
		FunktionID:               in.FunktionID,
		MaterialID:               in.MaterialID,
		OberflaechenbehandlungID: in.OberflaechenbehandlungID,
		FarbeID:                  in.FarbeID,
		ReserveID:                in.ReserveID,
	}

	typ, err := s.typRepo.FindByID(b.TypID)
	if err != nil {
		return nil, err
	}
	art, err := s.artRepo.FindByID(b.HerstellungsartID)
	if err != nil {
		return nil, err
	}
	versch, err := s.verschRepo.FindByID(b.VerschleissteilID)
	if err != nil {
		return nil, err
	}
	fun, err := s.funktionRepo.FindByID(b.FunktionID)
	if err != nil {
		return nil, err
	}
	mat, err := s.materialRepo.FindByID(b.MaterialID)
	if err != nil {
		return nil, err
	}
	oberf, err := s.oberflaecheRepo.FindByID(b.OberflaechenbehandlungID)
	if err != nil {
		return nil, err
	}
	farbe, err := s.farbeRepo.FindByID(b.FarbeID)
	if err != nil {
		return nil, err
	}
	res, err := s.reserveRepo.FindByID(b.ReserveID)
	if err != nil {
		return nil, err
	}

	sachnummerKey := domain.SachnummerKey{
		TypSymbol:             typ.Symbol,
		HerstellungsartSymbol: art.Symbol,
		VerschleissteilSymbol: versch.Symbol,
		FunktionSymbol:        fun.Symbol,
		MaterialSymbol:        mat.Symbol,
		OberflaecheSymbol:     oberf.Symbol,
		FarbeSymbol:           farbe.Symbol,
		ReserveSymbol:         res.Symbol,
	}

	count, err := s.repo.CountByAttributes(
		b.TypID,
		b.HerstellungsartID,
		b.VerschleissteilID,
		b.FunktionID,
		b.MaterialID,
		b.OberflaechenbehandlungID,
		b.FarbeID,
		b.ReserveID,
	)
	if err != nil {
		return nil, err
	}

	suffix := domain.GenerateHexSuffix(count)
	b.Sachnummer = domain.BuildSachnummer(sachnummerKey, suffix)

	return s.repo.Create(b)
}

// ListBauteile gibt alle Bauteile aus dem Repository zurück.
func (s *BauteilService) ListBauteile() ([]*domain.Bauteil, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*domain.Bauteil{}
	}
	return list, nil
}
