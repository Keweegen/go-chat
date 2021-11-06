package handlers

import (
    "reflect"
    "strings"

    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
)

type validationErrors fiber.Map

func validateRequest(req interface{}) validationErrors {
    var errors validationErrors

    validate := validator.New()
    validate.RegisterTagNameFunc(validatorTagNameJSON)

    if err := validate.Struct(req); err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            errors[err.Field()] = err.Error()
        }
    }

    return errors
}

func validatorTagNameJSON(field reflect.StructField) string {
    name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
    if name == "-" {
        return ""
    }
    return name
}

func sendSuccess(ctx *fiber.Ctx, result fiber.Map) error {
    return ctx.Status(fiber.StatusOK).JSON(result)
}

func sendBadRequest(ctx *fiber.Ctx, result fiber.Map) error {
    return ctx.Status(fiber.StatusBadRequest).JSON(result)
}

func sendNotFound(ctx *fiber.Ctx, message string) error {
    return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "message": message,
    })
}

func sendValidationError(ctx *fiber.Ctx, errors validationErrors) error {
    return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
        "message": "validation error",
        "errors": errors,
    })
}

func sendInternalServerError(ctx *fiber.Ctx, message string) error {
    return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "message": message,
    })
}
