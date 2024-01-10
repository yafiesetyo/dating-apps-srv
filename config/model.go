package config

import "time"

type (
	Config struct {
		Port  string      `mapstructure:"port"`
		JWT   JWTConfig   `mapstructure:"jwt"`
		Redis RedisConfig `mapstructure:"redis"`
		Db    DbConfig    `mapstructure:"db"`
	}

	DbConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DbName   string `mapstructure:"dbname"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	}

	RedisConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Db   int    `mapstructure:"db"`
	}

	JWTConfig struct {
		TTL       time.Duration `mapstructure:"ttl"`
		SecretKey string        `mapstructure:"secret_key"`
	}
)
