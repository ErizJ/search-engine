package config

import "os"

// Config 全局配置结构
type Config struct {
	Server ServerConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port string
}

type MySQLConfig struct {
	DSN string
}

type RedisConfig struct {
	Addr     string
	Password string
}

// getEnv 读取环境变量，不存在则返回默认值
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Load 加载配置（优先读取环境变量，方便 Docker 部署）
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		MySQL: MySQLConfig{
			DSN: getEnv(
				"MYSQL_DSN",
				"root:password@tcp(127.0.0.1:3306)/searchengine?charset=utf8mb4&parseTime=True&loc=Local",
			),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "127.0.0.1:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
	}
}
