package config

import (
	"log"
	"os"
	"strconv"

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
	BCRYPT_Salt int
	JWTExp      int
}

func Get() (*Config, error) {

	var Conf *Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWTExp, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		log.Fatal("Error parsing JWT_EXPIRATION")
		JWTExp = 60
	}

	salt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		log.Fatal("Error parsing BCRYPT_SALT")
		salt = 10
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

		JWTSecret:   os.Getenv("JWT_SECRET"),
		BCRYPT_Salt: salt,
		JWTExp:      JWTExp,
	}

	return Conf, nil
}
