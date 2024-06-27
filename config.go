package demo

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Echo  *EchoConfig
	Mongo *MongoConfig
}

type EchoConfig struct {
	Host string `env:"HOST, required"`
	Port string `env:"PORT, required"`
}

type MongoConfig struct {
	ConnectionUrl string `env:"MONGO_CONNECTION_URL, required"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	root_filepath, defined := os.LookupEnv("PROJECT_ROOT_DIR")

	if !defined {
		return nil, errors.New("root directory not defined, set the PROJECT_ROOT_DIR env variable")
	}

	env, defined := os.LookupEnv("PROJECT_ENV")

	if !defined {
		return nil, errors.New("environment not defined, set the PROJECT_ENV env variable")
	}

	dotEnvPath := filepath.Join(root_filepath, ".env."+env)
	err := godotenv.Load(dotEnvPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := envconfig.Process(ctx, config); err != nil {
		return nil, errors.New("error loading config: " + err.Error())
	}

	return config, nil
}
