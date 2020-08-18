package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Mysql struct {
	Ip       string `json:"Ip" `
	Port     string `json:"Port" `
	User     string `json:"User" `
	Password string `json:"Password" `
	Db       string `json:"Db" `
}

type Redis struct {
	Ip   string `json:"Ip"`
	Port string `json:"Port"`
	Auth string `json:"Auth"`
}

type Jwt struct {
	Secret   []byte `json:"Secret"`
	LifeTime int    `json:"LifeTime" `
}

type MyConfig struct {
	Mysql
	Redis
	Jwt
	Telegram
}

type Telegram struct {
	Token   string
	Channel int
}

var Config = MyConfig{
	Mysql: Mysql{
		Ip:       GetStr("DB_HOST"),
		Port:     GetStr("DB_PORT"),
		User:     GetStr("DB_USERNAME"),
		Password: GetStr("DB_PASSWORD"),
		Db:       GetStr("DB_NAME"),
	},
	Redis: Redis{
		Ip:   GetStr("REDIS_ENDPOINT"),
		Port: GetStr("REDIS_PORT"),
		Auth: GetStr("REDIS_AUTH"),
	},
	Jwt: Jwt{
		Secret:   GetBytes("SECRET_KEY"),
		LifeTime: GetInt("TOKEN_LIFETIME"),
	},
	Telegram: Telegram{
		Token:   GetStr("BOT_TOKEN"),
		Channel: GetInt("CHANNEL_ID"),
	},
}

func GetStr(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	return value
}

func GetInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	output, err := strconv.Atoi(value)
	if err != nil {
		log.Panic(`Environmental variable [` + key + `] is not an integer`)
	}

	return output
}

func GetBytes(key string) []byte {
	value := os.Getenv(key)
	if value == "" {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	return []byte(value)
}
