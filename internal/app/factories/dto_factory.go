package factories

import (
	"github.com/aadejanovs/catalog/internal/app/domain"
	"github.com/aadejanovs/catalog/internal/app/dtos"
)

func BlueprintResponseDtoFromBlueprint(blueprint *domain.Blueprint) *dtos.BlueprintResponseDto {
	return &dtos.BlueprintResponseDto{
		ID:        blueprint.ID,
		Version:   blueprint.Version,
		BrandName: blueprint.BrandName,
	}
}
