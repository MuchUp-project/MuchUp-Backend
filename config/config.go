package config

import "os"

type Config struct {
    SecretKey string
    DBHost    string
    DBUser    string
    DBPass    string
    DBName    string
    DBPort    string
    GRPCPort  string
    HTTPPort  string
}

func LoadConfig() *Config {
    return &Config{
        SecretKey: os.Getenv("JWT_SECRET_KEY"),
        DBHost:    os.Getenv("DB_HOST"),
        DBUser:    os.Getenv("DB_USER"),
        DBPass:    os.Getenv("DB_PASSWORD"),
        DBName:    os.Getenv("DB_NAME"),
        DBPort:    os.Getenv("DB_PORT"),
        GRPCPort:  getEnv("GRPC_PORT", "50051"),
        HTTPPort:  getEnv("SERVER_PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    val := os.Getenv(key)
    if val == "" {
        return fallback
    }
    return val
}