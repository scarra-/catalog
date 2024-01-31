package blueprint

import (
	"github.com/aadejanovs/catalog/internal/app/dtos"
	repo "github.com/aadejanovs/catalog/internal/app/repositories/blueprint"
	"github.com/aadejanovs/catalog/internal/utils"
)

type ListBlueprintsService struct {
	repo *repo.BlueprintRepository
}

func NewListBlueprintsService(repo *repo.BlueprintRepository) *ListBlueprintsService {
	return &ListBlueprintsService{repo: repo}
}

func (s *ListBlueprintsService) List(page, limit int) (*utils.Pagination[dtos.BlueprintListItemResponseDto], error) {
	return s.repo.List(page, limit)
}

func (s *ListBlueprintsService) CursorList(cursor string, limit int) (*utils.CursorPagination[dtos.BlueprintListItemResponseDto], error) {
	return s.repo.CursorList(cursor, limit)
}
