package blueprint

import (
	"math"

	"github.com/aadejanovs/catalog/internal/app/domain"
	"github.com/aadejanovs/catalog/internal/app/dtos"
	"github.com/aadejanovs/catalog/internal/utils"
	"gorm.io/gorm"
)

type BlueprintRepository struct {
	db *gorm.DB
}

func NewBlueprintRepository(db *gorm.DB) *BlueprintRepository {
	return &BlueprintRepository{db: db}
}

func (repo *BlueprintRepository) Save(blueprint *domain.Blueprint) error {
	result := repo.db.Save(blueprint)
	return result.Error
}

func (repo *BlueprintRepository) OfId(id string) (*domain.Blueprint, error) {
	var bp domain.Blueprint
	result := repo.db.Where("id = ?", id).First(&bp)
	return &bp, result.Error
}

func (repo *BlueprintRepository) List(page, limit int) (*utils.Pagination[dtos.BlueprintListItemResponseDto], error) {
	var blueprints []*domain.Blueprint
	var bpItemDtos []*dtos.BlueprintListItemResponseDto

	pagination := &utils.Pagination[dtos.BlueprintListItemResponseDto]{
		Limit: limit,
		Page:  page,
	}

	repo.db.Scopes(paginate(blueprints, pagination, repo.db)).Find(&blueprints)

	for _, bp := range blueprints {
		bpItemDtos = append(bpItemDtos, &dtos.BlueprintListItemResponseDto{
			ID:        bp.ID,
			Version:   bp.Version,
			BrandName: bp.BrandName,
		})
	}

	pagination.Rows = bpItemDtos

	return pagination, nil
}

// This function is for testing purposes and does not represent production ready
// cursor based pagination.
func (repo *BlueprintRepository) CursorList(cursor string, limit int) (*utils.CursorPagination[dtos.BlueprintListItemResponseDto], error) {
	var blueprints []*domain.Blueprint
	var bpItemDtos []*dtos.BlueprintListItemResponseDto

	pagination := &utils.CursorPagination[dtos.BlueprintListItemResponseDto]{
		Limit: limit,
		Rel:   cursor,
	}

	bp, err := repo.OfId(cursor)
	if err != nil {
		repo.db.Model(bp).Last(bp)
	}

	repo.db.Scopes(cursorPaginate(blueprints, bp, pagination, repo.db)).Find(&blueprints)

	for _, bp := range blueprints {
		bpItemDtos = append(bpItemDtos, &dtos.BlueprintListItemResponseDto{
			ID:        bp.ID,
			Version:   bp.Version,
			BrandName: bp.BrandName,
		})
	}

	pagination.Rows = bpItemDtos

	return pagination, nil
}

func paginate(value interface{}, pagination *utils.Pagination[dtos.BlueprintListItemResponseDto], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}

func cursorPaginate(value interface{}, bp *domain.Blueprint, pagination *utils.CursorPagination[dtos.BlueprintListItemResponseDto], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("incremental_id < ?", bp.IncrementalId).Limit(pagination.GetLimit()).Order("incremental_id desc")
	}
}
