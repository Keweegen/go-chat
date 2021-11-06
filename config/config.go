package config

const (
    clientFile = iota
)

type Config struct {
    client int

    Service  configService  `json:"service"`
    Cache    configCache    `json:"cache"`
    Database configDatabase `json:"database"`
    HTTP     configHTTP     `json:"http"`
}

type configService struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Debug   bool   `json:"debug"`
}

type configCache struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    DB       int    `json:"db"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type configDatabase struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    DB       string `json:"db"`
    Username string `json:"username"`
    Password string `json:"password"`
    SSLMode  string `json:"ssl_mode"`
}

type configHTTP struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

var config *Config
