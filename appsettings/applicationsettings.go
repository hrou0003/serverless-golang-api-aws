package appsettings

import "os"

func GetFromEnvironment(envVar string, fallback string) string {
	value := os.Getenv(envVar)

	if len(value) == 0 {
		return fallback
	}

	return value
}
