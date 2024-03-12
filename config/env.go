package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string

	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string

	S3Bucket string
	S3Secret string
	S3ID     string
}

func Get() (*Config, error) {

	var Conf *Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Conf = &Config{
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		S3Bucket:   os.Getenv("S3_BUCKET"),
		S3Secret:   os.Getenv("SECRET_KEY"),
		S3ID:       os.Getenv("S3_ID"),
	}

	return Conf, nil
}
