package blueprint

import (
	"github.com/aadejanovs/catalog/internal/app/domain"
	"github.com/aadejanovs/catalog/internal/app/dtos"
	repo "github.com/aadejanovs/catalog/internal/app/repositories/blueprint"
)

type CreateBlueprintService struct {
	repo *repo.BlueprintRepository
}

func NewCreateBlueprintService(repo *repo.BlueprintRepository) *CreateBlueprintService {
	return &CreateBlueprintService{repo: repo}
}

func (s *CreateBlueprintService) Create(dto *dtos.CreateBlueprintRequestDto) (*domain.Blueprint, error) {
	bp := domain.NewBlueprintFromDto(dto)
	err := s.repo.Save(bp)

	return bp, err
}
