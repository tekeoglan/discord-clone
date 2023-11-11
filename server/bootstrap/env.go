package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ClientAddress  string `mapstructure:"CLIENT_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	CacheHost      string `mapstructure:"CACHE_HOST"`
	CachePort      string `mapstructure:"CACHE_PORT"`
	CacheUser      string `mapstructure:"CACHE_USER"`
	CachePass      string `mapstructure:"CACHE_PASSWORD"`
}

func NewEnv() *Env {
	file := "dev.env"
	if os.Getenv("ENV") == "production" {
		file = ".env"
	}

	env := Env{}

	viper.SetConfigFile("/etc/discord-clone/" + file)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the env file:", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Can't decode env file: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The app is running in development mode")
	}

	return &env
}
