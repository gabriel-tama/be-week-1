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
	S3Url    string
	S3Region string

	JWTSecret   string
	BCRYPT_Salt string
}

func Get() (*Config, error) {

	var Conf *Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Conf = &Config{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USERNAME"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),

		S3Bucket: os.Getenv("S3_BUCKET_NAME"),
		S3Secret: os.Getenv("S3_SECRET_KEY"),
		S3ID:     os.Getenv("S3_ID"),
		S3Url:    os.Getenv("S3_BASE_URL"),
		S3Region: os.Getenv("S3_REGION"),
	}

	return Conf, nil
}
