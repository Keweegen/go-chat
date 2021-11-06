package handlers

import "github.com/gofiber/fiber/v2"

type ProfileHandlers struct{}

func (ph *ProfileHandlers) GetProfile(ctx *fiber.Ctx) error {
    profileUUID, err := ph.handleProfileParam(ctx)
    if err != nil {
        return sendValidationError(ctx, err)
    }

    // Get profile

    return sendSuccess(ctx, fiber.Map{
        "profile": fiber.Map{
            "uuid":        profileUUID,
            "name":        "Vasya Pupkin",
            "email":       "vasya.pupkin228@tg.org",
            "description": "Donuts are my life!!1!",
        },
    })
}

type updateProfileRequest struct {
    Name        string `json:"name" validate:"omitempty,min=3,max=255"`
    Email       string `json:"email" validate:"omitempty,email,min=6,max=255"`
    Description string `json:"description" validate:"omitempty,max=500"`
}

func (ph *ProfileHandlers) UpdateProfile(ctx *fiber.Ctx) error {
    profileUUID, err := ph.handleProfileParam(ctx)
    if err != nil {
        return sendValidationError(ctx, err)
    }

    fields := new(updateProfileRequest)
    if err := ctx.BodyParser(fields); err != nil {
        return sendInternalServerError(ctx, err.Error())
    }
    if err := validateRequest(*fields); err != nil {
        return sendValidationError(ctx, err)
    }

    // Update profile

    return sendSuccess(ctx, fiber.Map{
        "message": "profile updated",
        "profile": fiber.Map{
            "uuid":        profileUUID,
            "name":        fields.Name,
            "email":       fields.Email,
            "description": fields.Description,
        },
    })
}

func (ph *ProfileHandlers) handleProfileParam(ctx *fiber.Ctx) (string, validationErrors) {
    profileUUID := ctx.Params("profile")
    if profileUUID == "" {
        return "", validationErrors{"profile": "required"}
    }
    return profileUUID, nil
}
