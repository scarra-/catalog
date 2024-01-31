package dtos

type CreateBlueprintRequestDto struct {
	Version   string `json:"version" validate:"required"`
	BrandName string `json:"brand_name" validate:"required"`
}

type BlueprintListItemResponseDto struct {
	ID        string `json:"id"`
	Version   string `json:"version"`
	BrandName string `json:"brand_name"`
}

type BlueprintResponseDto struct {
	ID        string `json:"id"`
	Version   string `json:"version"`
	BrandName string `json:"brand_name"`
}
