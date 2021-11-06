package handlers

import "github.com/gofiber/fiber/v2"

type MessageHandlers struct{}

func (mh *MessageHandlers) List(ctx *fiber.Ctx) error {
    return nil
}

func (mh *MessageHandlers) CreateMessage(ctx *fiber.Ctx) error {
    return nil
}

func (mh *MessageHandlers) GetMessage(ctx *fiber.Ctx) error {
    return nil
}

func (mh *MessageHandlers) UpdateMessage(ctx *fiber.Ctx) error {
    return nil
}

func (mh *MessageHandlers) DeleteMessage(ctx *fiber.Ctx) error {
    return nil
}
