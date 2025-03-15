package utils

import (
	"github.com/caarlos0/env/v11"
)

type DatabaseConfig struct {
	Host     string            `env:"DB_HOST,required,notEmpty"`
	Port     uint16            `env:"DB_PORT,required,notEmpty"`
	DBName   string            `env:"DB_NAME,required,notEmpty"`
	Username string            `env:"DB_USERNAME,required,notEmpty"`
	Password string            `env:"DB_PASSWORD,required,notEmpty"`
	Params   map[string]string `env:"DB_PARAMS" envSeparator:"&" envKeyValSeparator:"="`
}

var internalDatabaseConfig *DatabaseConfig

func GetDatabaseConfig() *DatabaseConfig {
	if internalDatabaseConfig == nil {
		internalDatabaseConfig = &DatabaseConfig{}
		err := env.Parse(internalDatabaseConfig)
		if err != nil {
			panic(err)
		}
	}

	return internalDatabaseConfig
}
