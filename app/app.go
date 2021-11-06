package app

import (
    "database/sql"
    "errors"
    "fmt"

    "github.com/keweegen/go-chat/app/api"
    "github.com/keweegen/go-chat/cache"
    "github.com/keweegen/go-chat/config"
    "github.com/keweegen/go-chat/database"
    "github.com/keweegen/go-chat/logger"
)

type Application struct {
    Logger   *logger.Logger
    Config   *config.Config
    Cache    *cache.Cache
    Database *sql.DB
}

var App *Application

func GetApp() (*Application, error) {
    if App != nil {
        return App, nil
    }

    App = &Application{}

    if err := App.setConfig(); err != nil {
        return nil, err
    }
    if err := App.setCache(); err != nil {
        return nil, err
    }
    if err := App.setDatabase(); err != nil {
        return nil, err
    }

    App.setLogger()

    return App, nil
}

func (a *Application) ServeHTTP() {
    api.ServeHTTP(a)
}

func (a *Application) setLogger() {
    a.Logger = logger.GetLogger()

    if a.Config != nil {
        cfg := a.Config.Service

        message := fmt.Sprintf("%s - v%s", cfg.Name, cfg.Version)
        if cfg.Debug {
            message += " [DEBUG]"
        }

        a.Logger.With("service", message)
    }
}

func (a *Application) setConfig() error {
    opts := config.File{
        Path: ".",
        Name: "app.yml",
    }

    cfg, err := config.GetFile(opts)
    if err != nil {
        return err
    }

    a.Config = cfg
    return nil
}

func (a *Application) setCache() error {
    if a.Config == nil {
        return a.failConfigEmpty()
    }

    cfg := a.Config.Cache

    connection := cache.Connection{
        Client:   cache.ClientRedis,
        Host:     cfg.Host,
        Port:     cfg.Port,
        Database: cfg.DB,
        Username: cfg.Username,
        Password: cfg.Password,
    }

    cch, err := cache.GetCache(connection)
    if err != nil {
        return err
    }

    a.Cache = cch
    return nil
}

func (a *Application) setDatabase() error {
    if a.Config == nil {
        return a.failConfigEmpty()
    }

    cfg := a.Config.Database

    connection := database.Connection{
        Driver:   database.DriverPostgres,
        Host:     cfg.Host,
        Port:     cfg.Port,
        Database: cfg.DB,
        Username: cfg.Username,
        Password: cfg.Password,
        SSLMode:  cfg.SSLMode,
    }
    db, err := database.GetDatabase(connection)
    if err != nil {
        return err
    }

    a.Database = db
    return nil
}

func (a *Application) failConfigEmpty() error {
    return errors.New("first you need to initialize the config")
}
