package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    ServerPort   string
    DatabaseURL  string
    STUNServers  []string
}

// LoadConfig загружает конфигурацию из .env файла и переменных окружения
func LoadConfig() *Config {
    // Попытка загрузить .env файл (если он существует)
    if err := godotenv.Load(); err != nil {
        log.Println("Не удалось загрузить .env файл, используются переменные окружения")
    }

    config := &Config{
        ServerPort:   getEnv("SERVER_PORT", "8080"),
        DatabaseURL:  getEnv("DATABASE_URL", "postgres://user:password@localhost/dbname?sslmode=disable"),
        STUNServers:  getEnvAsSlice("STUN_SERVERS", []string{"stun:stun.l.google.com:19302"}),
    }

    return config
}

// getEnv получает значение переменной окружения или возвращает дефолтное значение
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

// getEnvAsSlice получает значение переменной окружения как массив строк или возвращает дефолтное значение
func getEnvAsSlice(key string, defaultValue []string) []string {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return splitAndTrim(value, ",")
    }
    return defaultValue
}

// splitAndTrim разбивает строку по разделителю и убирает пробелы
func splitAndTrim(str, sep string) []string {
    var result []string
    for _, item := range split(str, sep) {
        trimmed := trimSpaces(item)
        if len(trimmed) > 0 {
            result = append(result, trimmed)
        }
    }
    return result
}

// Дополнительные вспомогательные функции
func split(str, sep string) []string {
    return strings.Split(str, sep)
}

func trimSpaces(str string) string {
    return strings.TrimSpace(str)
}