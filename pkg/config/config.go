package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type IConfig interface {
	ApplicationName() string
	Mode() string
	HttpConfig() IHttpConfig
}

func NewConfig() IConfig {
	config := &Config{}
	config.LoadConfig()

	return config
}

type Config struct {
}

func GetEnv(key string, fallback string) string {
	return getEnv(key, fallback)
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func (*Config) ApplicationName() string {
	return GetEnv("APP_NAME", "miniapi")
}

func (*Config) Mode() string {
	return GetEnv("MODE", "development")
}

func (cfg *Config) LoadConfig() {

	mode := cfg.Mode()

	if mode != "test" {
		godotenv.Load(".env.local")
	}

	loadEnvFileName := ".env." + mode + ".local"
	if mode == "test" {
		workspaceDir := os.Getenv("WORKSPACE_DIR")
		if workspaceDir == "" {
			cwd, err := filepath.Abs(".")
			if err == nil {
				workspaceDir = filepath.Dir(cwd) + "/"
			}
		}
		loadEnvFileName = workspaceDir + ".env." + mode + ".local"
	}

	godotenv.Load(loadEnvFileName)
	if mode != "test" {
		godotenv.Load(".env.local")
	}

	godotenv.Load(".env." + mode)
	godotenv.Load()
}

func (*Config) HttpConfig() IHttpConfig {
	return &HttpConfig{}
}
