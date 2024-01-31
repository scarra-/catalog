package utils

import (
	"errors"

	customErr "github.com/aadejanovs/catalog/internal/app/errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func HandleHTTPErrors(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		logger := ctx.Locals("logger").(*zap.SugaredLogger)
		logger.Infow("error",
			"message", err,
		)
	}

	ctx.Status(404).JSON(customErr.NewErrorResponse((code)))
	return nil
}
