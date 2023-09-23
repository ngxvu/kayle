package conf

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// AppConfig presents app conf
type AppConfig struct {
	ServerEnv string `env:"SERVER_ENV" envDefault:"stg"`
	Port      string `env:"PORT"`
	// MongoDB
	MongoDBHost   string `env:"MONGODB_HOST"`
	MongoDBPort   string `env:"MONGODB_PORT"`
	MongoDBName   string `env:"MONGODB_DBNAME"`
	DBName        string `env:"DB_NAME"`
	EnableDB      string `env:"ENABLE_DB" envDefault:"true"`
	DbDebugEnable bool   `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

var config AppConfig

func SetEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
