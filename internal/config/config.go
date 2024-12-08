package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"time"
)

const (
	CONFIG_PATH = "./config/config.yml"
)

type Config struct {
	ServerCfg  ServerConfig  `yaml:"server"`
	StorageCfg StorageConfig `yaml:"storage"`
}

type ServerConfig struct {
	Env        string        `yaml:"env" env-default:"dev"`
	TokenTTL   time.Duration `yaml:"token_ttl"`
	GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	GRPCPort    string        `yaml:"grpc_port"`
	GRPCTimeout time.Duration `yaml:"grpc_timeout"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func New() *Config {
	var cfg Config

	err := cleanenv.ReadConfig(CONFIG_PATH, &cfg)
	if err != nil {
		slog.Warn("unable to read config")
		panic(err)
	}

	return &cfg
}
