package handlers

import (
	"strconv"
	"time"

	"github.com/aadejanovs/catalog/database"
	"github.com/aadejanovs/catalog/internal/app/dtos"
	"github.com/aadejanovs/catalog/internal/app/errors"
	"github.com/aadejanovs/catalog/internal/app/factories"
	repo "github.com/aadejanovs/catalog/internal/app/repositories/blueprint"
	"github.com/aadejanovs/catalog/internal/app/repositories/blueprint_dto"
	"github.com/aadejanovs/catalog/internal/app/services/blueprint"
	validation "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetBlueprint(c *fiber.Ctx) error {
	logger := c.Locals("logger").(*zap.SugaredLogger)
	redisClient, ok := c.Locals("redis").(*database.Redis)
	if !ok {
		logger.Error("cannot_retrieve_redis_from_context")
	}

	redisRepo := blueprint_dto.NewRedisRepository(redisClient)
	bpDto, err := redisRepo.Get(c.Params("id"))
	if err == nil {
		return c.JSON(bpDto)
	}

	db := c.Locals("db").(*gorm.DB)
	service := blueprint.NewGetBlueprintService(repo.NewBlueprintRepository(db))

	start := time.Now()
	bp, err := service.Get(c.Params("id"))
	logger.Infow("time_elapsed_from_mysql", "type", "bp-view", "time", time.Since(start))

	if err != nil {
		return c.Status(404).JSON(errors.NewErrorResponse(404))
	}

	bpDto = factories.BlueprintResponseDtoFromBlueprint(bp)
	err = redisRepo.Set(bpDto)
	if err != nil {
		logger.Error("cannot_save_bp_to_redis", "message", err)
	}

	return c.JSON(bpDto)
}

func ListBlueprints(c *fiber.Ctx) error {
	logger := c.Locals("logger").(*zap.SugaredLogger)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "20"))

	db := c.Locals("db").(*gorm.DB)
	service := blueprint.NewListBlueprintsService(repo.NewBlueprintRepository(db))

	start := time.Now()
	response, err := service.List(page, limit)
	logger.Infow("time_elapsed_from_mysql", "type", "bp-list", "time", time.Since(start))
	if err != nil {
		return c.Status(400).JSON(errors.NewErrorResponse(400))
	}

	return c.JSON(response)
}

func CursorListBlueprints(c *fiber.Ctx) error {
	logger := c.Locals("logger").(*zap.SugaredLogger)
	cursor := c.Query("cursor")
	limit, _ := strconv.Atoi(c.Query("limit", "20"))

	db := c.Locals("db").(*gorm.DB)
	service := blueprint.NewListBlueprintsService(repo.NewBlueprintRepository(db))

	start := time.Now()
	response, err := service.CursorList(cursor, limit)
	logger.Infow("time_elapsed_from_mysql", "type", "bp-list", "time", time.Since(start))
	if err != nil {
		return c.Status(400).JSON(errors.NewErrorResponse(400))
	}

	return c.JSON(response)
}

func CreateBlueprint(c *fiber.Ctx) error {
	dto := &dtos.CreateBlueprintRequestDto{}

	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(errors.NewErrorResponse(400))
	}

	if err := validation.New().Struct(dto); err != nil {
		vErr, ok := err.(validation.ValidationErrors)
		if !ok {
			return err
		}

		return c.Status(400).JSON(errors.NewValidationErrorResponse(vErr))
	}

	db := c.Locals("db").(*gorm.DB)
	service := blueprint.NewCreateBlueprintService(repo.NewBlueprintRepository(db))
	bp, err := service.Create(dto)
	if err != nil {
		logger := c.Locals("logger").(*zap.SugaredLogger)
		logger.Errorw("error_while_creating_blueprint", "message", err)

		return c.Status(503).JSON(errors.NewErrorResponse(503))
	}

	responseDto := factories.BlueprintResponseDtoFromBlueprint(bp)

	return c.Status(201).JSON(responseDto)
}
