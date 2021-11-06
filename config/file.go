package config

import "github.com/spf13/viper"

type File struct {
    Path string
    Name string
}

func GetFile(opts File) (*Config, error) {
    if config != nil && config.client == clientFile {
        return config, nil
    }

    config = &Config{}
    config.client = clientFile

    viper.SetConfigName(opts.Name)
    viper.AddConfigPath(opts.Path)

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return config, nil
}
