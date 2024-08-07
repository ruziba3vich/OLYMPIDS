package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	SecretKey      string          `mapstructure:"SECRET_KEY"`
	Host           string          `mapstructure:"HOST"`
	GinServerPort  string          `mapstructure:"GIN_SERVER_PORT"`
	GrpcServerPort string          `mapstructure:"GRPC_SERVER_PORT"`
	DatabaseConfig *DatabaseConfig `mapstructure:"DATABASE"`
	RedisConfig    *RedisConfigs   `mapstructure:"REDIS"`
}

type DatabaseConfig struct {
	Port     string `mapstructure:"PORT"`
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
}

type RedisConfigs struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	DB       string `mapstructure:"DB"`
	Password string `mapstructure:"PASSWORD"`
}

func InitConfig(path string) (*Config, error) {
	var config Config
	if err := LoadConfig(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadConfig(path string, config *Config) error {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	// Set default values
	viper.SetDefault("SECRET_KEY", "your_secret_key")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("GIN_SERVER_PORT", "8080")
	viper.SetDefault("GRPC_SERVER_PORT", "50051")
	viper.SetDefault("DATABASE.PORT", "5432")
	viper.SetDefault("DATABASE.HOST", "localhost")
	viper.SetDefault("DATABASE.USER", "postgres")
	viper.SetDefault("DATABASE.PASSWORD", "1702")
	viper.SetDefault("DATABASE.NAME", "medal_service")
	viper.SetDefault("REDIS.HOST", "localhost")
	viper.SetDefault("REDIS.PORT", "6379")
	viper.SetDefault("REDIS.DB", "0")
	viper.SetDefault("REDIS.PASSWORD", "")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("unable to decode into struct: %v", err)
	}

	return nil
}
