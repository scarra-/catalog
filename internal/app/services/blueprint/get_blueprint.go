package blueprint

import (
	"github.com/aadejanovs/catalog/internal/app/domain"
	repo "github.com/aadejanovs/catalog/internal/app/repositories/blueprint"
)

type GetBlueprintService struct {
	repo *repo.BlueprintRepository
}

func NewGetBlueprintService(repo *repo.BlueprintRepository) *GetBlueprintService {
	return &GetBlueprintService{repo: repo}
}

func (s *GetBlueprintService) Get(id string) (*domain.Blueprint, error) {
	return s.repo.OfId(id)
}
