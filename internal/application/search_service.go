package application

import "KeepInventory/internal/domain"

type SearchService struct {
	repo SearchRepository
}

func NewSearchService(repo SearchRepository) *SearchService {
	return &SearchService{repo: repo}
}

func (s *SearchService) Search(req domain.SearchRequest) ([]domain.SearchResult, error) {
	if req.Limit <= 0 {
		req.Limit = 10
	}
	return s.repo.Search(req)
}
