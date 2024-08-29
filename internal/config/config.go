package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBHost     string `mapstructure:"DB_HOST"`
    DBPort     string `mapstructure:"DB_PORT"`
    DBName     string `mapstructure:"DB_NAME"`
    RedisAddr  string `mapstructure:"REDIS_ADDR"`
    JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (Config, error) {
    var config Config
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        return config, err
    }
    if err := viper.Unmarshal(&config); err != nil {
        return config, err
    }
    return config, nil
}
