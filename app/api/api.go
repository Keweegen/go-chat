package api

import (
    "fmt"

    "github.com/gofiber/fiber/v2"
    "github.com/keweegen/go-chat/app"
    "github.com/keweegen/go-chat/config"
)

func ServeHTTP(app *app.Application) {
    api := fiber.New()

    routes(api)

    if err := api.Listen(getAddress(app.Config)); err != nil {
        app.Logger.Error("failed listen HTTP server", "error", err)
    }
}

func getAddress(cfg *config.Config) string {
    return fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
}
