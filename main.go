package main

import (
    "log"

    "github.com/keweegen/go-chat/app"
)

func main() {
    chat, err := app.GetApp()
    if err != nil {
        log.Fatal("Failed get application. Error: ", err)
        return
    }

    chat.ServeHTTP()

    chat.Logger.Info("Application successfully started!")
}
