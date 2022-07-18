package config

import "os"

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
}

func InitConfiguration() Config {

	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8080"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "admin"),
		// DB_USERNAME:    GetOrDefault("DB_USERNAME", "root"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "XmNP9Y5DQC2rUVsz4sDi"),
		// DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "123456"),
		DB_NAME:        GetOrDefault("DB_NAME", "vaccine"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "database-4.cn1x7ayh0czt.us-west-1.rds.amazonaws.com"),
		// DB_HOST:        GetOrDefault("DB_HOST", "127.0.0.1"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "ABC"),
	}
}

func GetOrDefault(envKey, defaultValue string) string {
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}
	return defaultValue
}
  