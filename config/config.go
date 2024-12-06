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
	if err := godotenv.Load("config/.env"); err != nil {
		fmt.Println("NO .env file not bilmadim foundd")
	}

	cfg := Config{}
	cfg.ServerHost = cast.ToString(getOrDefaultValue("SERVER_HOST", "18.153.65.158"))
	cfg.Postgres = Postgres{
		Host:     cast.ToString(getOrDefaultValue("POSTGRES_HOST", "18.153.65.158")),
		Port:     cast.ToInt(getOrDefaultValue("POSTGRES_PORT", "5432")),
		User:     cast.ToString(getOrDefaultValue("POSTGRES_USER", "maxmurjon")),
		Password: cast.ToString(getOrDefaultValue("POSTGRES_PASSWORD", "max22012004")),
		DataBase: cast.ToString(getOrDefaultValue("POSTGRES_DATABASE", "comics"))}
	cfg.Redis = Redis{
		Host:     cast.ToString(getOrDefaultValue("REDIS_HOST", "18.153.65.158")),
		Port:     cast.ToInt(getOrDefaultValue("REDIS_PORT", "5432")),
		Password: cast.ToString(getOrDefaultValue("REDIS_PASSWORD", "max22012004")),
		DataBase: cast.ToString(getOrDefaultValue("REDIS_DATABASE", "comics"))}

	cfg.SekretKey = cast.ToString(getOrDefaultValue("SEKRET_KEY", "sekret"))
	fmt.Println(cfg)
	return &cfg
}

func getOrDefaultValue(key string, defaultValue string) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
