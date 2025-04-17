package tests_db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arly/arlyApi/utilities"
	"github.com/joho/godotenv"
)

func DbConfig() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}
	fmt.Println("Current directory:", dir)

	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" || dbPassword == "''" {
		dbPassword = "''"
		utilities.LogError("DB_PASSWORD is empty, this may cause error when connecting to the database", nil)
	}
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimezone := os.Getenv("DB_TIMEZONE")
	if dbTimezone == "" {
		dbTimezone = "Asia/Shanghai"
		utilities.LogInfo("DB_TIMEZONE is empty, defaulting to Asia/Shanghai")
	}

	// DSN (Data Source Name)
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode + " TimeZone=" + dbTimezone
	return dsn
}

func environmentType() string {
	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
		utilities.LogInfo("ENVIRONMENT is not set, defaulting to development")
	}
	return environment
}

func ServerPort() string {
	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
		utilities.LogInfo("PORT is not set, defaulting to 5050")
	}
	return port
}

func ServerUrl() string {
	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	env := environmentType()
	url := os.Getenv("SERVER_URL")

	if url == "" {
		if env == "development" {
			url = "http://0.0.0.0"
		} else if env == "production" {
			utilities.LogFatal("SERVER_URL is not set in production environment", nil)
		}
	}

	if ServerPort() == "80" {
		return url
	}

	return url + ":" + ServerPort()
}

type jwtInformation struct {
	secret        string
	expire        string
	refreshSecret string
	refreshExpire string
}

func JwtInformation() jwtInformation {
	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	env := environmentType()

	secret := os.Getenv("JWT_SECRET")
	expire := os.Getenv("JWT_EXPIRE")
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	refreshExpire := os.Getenv("JWT_REFRESH_EXPIRE")

	if env == "development" {
		if secret == "" {
			secret = "dev_secret"
			utilities.LogInfo("JWT_SECRET is not set, defaulting to dev_secret")
		}
		if expire == "" {
			expire = "24h"
			utilities.LogInfo("JWT_EXPIRE is not set, defaulting to 24h")
		}
		if refreshSecret == "" {
			refreshSecret = "dev_refresh_secret"
			utilities.LogInfo("JWT_REFRESH_SECRET is not set, defaulting to dev_refresh_secret")
		}
		if refreshExpire == "" {
			refreshExpire = "72h"
			utilities.LogInfo("JWT_REFRESH_EXPIRE is not set, defaulting to 72h")
		}
	}

	if secret == "" || expire == "" || refreshSecret == "" || refreshExpire == "" {
		utilities.LogFatal("JWT environment variables are not properly set", nil)
	}

	return jwtInformation{
		secret:        secret,
		expire:        expire,
		refreshSecret: refreshSecret,
		refreshExpire: refreshExpire,
	}
}

type OpenAiConfig struct {
	APIKey      string
	Model       string
	Temperature float64
}

func GetOpenAiConfig() OpenAiConfig {
	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		utilities.LogFatal("OPENAI_API_KEY is not set", nil)
	}

	return OpenAiConfig{
		APIKey:      apiKey,
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
	}
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
	SSL      bool
}

func GetSMTPConfig() SMTPConfig {

	if err := godotenv.Load("../.env.tests"); err != nil {
		utilities.LogFatal("Error loading .env file", err)
	}

	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	name := os.Getenv("SMTP_NAME")
	password := os.Getenv("SMTP_PASSWORD")
	sslStr := os.Getenv("SMTP_SSL")

	if host == "" || portStr == "" || user == "" || name == "" || password == "" || sslStr == "" {
		utilities.LogFatal("SMTP environment variables are not properly set", nil)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		utilities.LogFatal("SMTP_PORT must be a valid integer", err)
	}

	ssl, err := strconv.ParseBool(sslStr)
	if err != nil {
		utilities.LogFatal("SMTP_SSL must be a valid boolean (true/false)", err)
	}

	return SMTPConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Name:     name,
		Password: password,
		SSL:      ssl,
	}
}
