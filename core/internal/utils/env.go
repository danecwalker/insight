package utils

import (
	"os"
	"strconv"
)

func GetEnvString(key string, fallback string) string {
	if value, ok := GetEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	if value, ok := GetEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	if value, ok := GetEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return b
	}
	return fallback
}

func GetEnvFloat64(key string, fallback float64) float64 {
	if value, ok := GetEnv(key); ok {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fallback
		}
		return f
	}
	return fallback
}

func GetEnvInt64(key string, fallback int64) int64 {
	if value, ok := GetEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

func GetEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
