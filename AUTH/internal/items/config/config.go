package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server        ServerConfig
		Database      DatabaseConfig
		Redis         RedisConfig
		JWT           JWTConfig
		RabbitMQ      RabbitMQConfig
		TableName     string
		BookId        string
		Title         string
		Author        string
		PublisherYear string
	}
	JWTConfig struct {
		SecretKey string
	}

	ServerConfig struct {
		Port string
	}
	DatabaseConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	RedisConfig struct {
		Host string
		Port string
	}
	RabbitMQConfig struct {
		RabbitMQ string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.Server.Port = ":" + os.Getenv("SERVER_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")
	c.Redis.Host = os.Getenv("REDIS_HOST")
	c.Redis.Port = os.Getenv("REDIS_PORT")
	c.TableName = os.Getenv("TABLE_NAME")
	c.BookId = os.Getenv("BOOK_ID")
	c.Title = os.Getenv("TITLE")
	c.Author = os.Getenv("AUTHOR")
	c.PublisherYear = os.Getenv("PUB_YEAR")
	c.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")
	c.RabbitMQ.RabbitMQ = os.Getenv("RABBITMQ_URI")

	return nil
}

func New() (*Config, error) {
	var config Config
	if err := config.Load(); err != nil {
		return nil, err
	}
	return &config, nil
}
