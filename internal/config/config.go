package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type TaskManagerDB struct {
	PostgresHost     string `envconfig:"TASK_MANAGER_POSTGRES_HOST"`
	PostgresPort     string `envconfig:"TASK_MANAGER_POSTGRES_PORT"`
	PostgresUser     string `envconfig:"TASK_MANAGER_POSTGRES_SERVICE_USERNAME"`
	PostgresPassword string `envconfig:"TASK_MANAGER_POSTGRES_SERVICE_PASSWORD"`
	PostgresDatabase string `envconfig:"TASK_MANAGER_POSTGRES_SERVICE_DATABASE"`
	PostgresParams   string `envconfig:"TASK_MANAGER_POSTGRES_PARAMS"`
	MaxConnection    int    `envconfig:"TASK_MANAGER_POSTGRES_MAX_CONNECTION" default:"10"`
	MinConnection    int    `envconfig:"TASK_MANAGER_POSTGRES_MIN_CONNECTION" default:"0"`
}
type Config struct {
	TaskManagerDB TaskManagerDB
	ServiceName   string `envconfig:"TASK_MANAGER_SERVER_SERVICE_NAME"`
	Debug         bool   `envconfig:"TASK_MANAGER_DEBUG"`
	Secret        string `envconfig:"TASK_MANAGER_SECRET" default:"secret"`
}

func FromEnv() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("error while parse env config | %w", err)
	}

	return cfg, nil
}

func (c *Config) TaskManagerDBURL() string {
	pgURL := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		c.TaskManagerDB.PostgresUser,
		c.TaskManagerDB.PostgresPassword,
		c.TaskManagerDB.PostgresHost,
		c.TaskManagerDB.PostgresPort,
		c.TaskManagerDB.PostgresDatabase,
	)
	if c.TaskManagerDB.PostgresParams != "" {
		pgURL = fmt.Sprintf("%v?%v", pgURL, c.TaskManagerDB.PostgresParams)
	}

	return pgURL
}
