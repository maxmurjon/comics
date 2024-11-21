package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	TimeExpiredAt = time.Hour * 24
)

type Config struct {
	Environment string

	ServerHost string
	ServerPort string

	Redis Redis

	Postgres Postgres

	SekretKey string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DataBase string
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DataBase string
}

func Load() *Config {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("NO .env file not foundd")
	}

	cfg := Config{}
	cfg.ServerHost = cast.ToString(getOrDefaultValue("SERVER_HOST", "localhost"))
	cfg.Postgres = Postgres{
		Host:     cast.ToString(getOrDefaultValue("POSTGRES_HOST", "localhost")),
		Port:     cast.ToInt(getOrDefaultValue("POSTGRES_PORT", "5432")),
		User:     cast.ToString(getOrDefaultValue("POSTGRES_USER", "admin")),
		Password: cast.ToString(getOrDefaultValue("POSTGRES_PASSWORD", "password")),
		DataBase: cast.ToString(getOrDefaultValue("POSTGRES_DATABASE", "database"))}
	cfg.Redis = Redis{
		Host:     cast.ToString(getOrDefaultValue("REDIS_HOST", "localhost")),
		Port:     cast.ToInt(getOrDefaultValue("REDIS_PORT", "5432")),
		Password: cast.ToString(getOrDefaultValue("REDIS_PASSWORD", "password")),
		DataBase: cast.ToString(getOrDefaultValue("REDIS_DATABASE", "database"))}

	cfg.SekretKey = cast.ToString(getOrDefaultValue("SEKRET_KEY", "sekret"))
	return &cfg
}

func getOrDefaultValue(key string, defaultValue string) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
