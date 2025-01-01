package env

import (
	"os"

	"github.com/mohdjishin/sportsphere/constants"
)

func GetEnv(key constants.EnvironmentVariable, fallback string) string {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		return fallback
	}
	return value
}
