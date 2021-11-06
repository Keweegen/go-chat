package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/keweegen/go-chat/app/api/handlers"
)

func routes(f *fiber.App) {
    apiGroup := f.Group("/api/v1")

    profileRoutes(apiGroup)
    chatRoutes(apiGroup)
}

func profileRoutes(router fiber.Router) {
    ph := new(handlers.ProfileHandlers)
    group := router.Group("/profile")
    group.Get("/:profile", ph.GetProfile)
    group.Put("/:profile", ph.UpdateProfile)
}

func chatRoutes(router fiber.Router) {
    ch := new(handlers.ChatHandlers)
    group := router.Group("/chat")
    group.Get("/list", ch.List)
    group.Post("", ch.CreateChat)
    group.Get("/:chat", ch.GetChat)
    group.Put("/:chat", ch.UpdateChat)
    group.Delete("/:chat", ch.DeleteChat)

    messageRoutes(group.Group("/:chat"))
}

func messageRoutes(router fiber.Router) {
    mh := new(handlers.MessageHandlers)
    group := router.Group("/message")
    group.Get("/list", mh.GetMessage)
    group.Post("", mh.CreateMessage)
    group.Put("/:message", mh.UpdateMessage)
    group.Delete("/:message", mh.DeleteMessage)
}
