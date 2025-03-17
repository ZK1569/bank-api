package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type envVariable struct {
	DatabaseUri string
}

var EnvVariable *envVariable

func init() {
	EnvVariable = &envVariable{
		DatabaseUri: getString("MONGODB_URI", ""),
	}
}

func getString(key, fallback string) string {
	godotenv.Load()

	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}
