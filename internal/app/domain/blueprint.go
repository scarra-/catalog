package domain

import (
	"time"

	"github.com/aadejanovs/catalog/internal/app/dtos"
	"github.com/aadejanovs/catalog/internal/utils"
	"gorm.io/gorm"
)

type Blueprint struct {
	IncrementalId string `gorm:"column:incremental_id"`
	ID            string `gorm:"primarykey"`

	Version   string `gorm:"column:version;not null"`
	BrandName string `gorm:"column:brand_name;not null"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func NewBlueprintFromDto(dto *dtos.CreateBlueprintRequestDto) *Blueprint {
	blueprintID := "bp-" + utils.RandomKey(29)

	return &Blueprint{
		ID:        blueprintID,
		Version:   dto.Version,
		BrandName: dto.BrandName,
	}
}
