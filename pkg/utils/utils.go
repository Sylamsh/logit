package utils

import "os"

func GetEnvWithDefault(key, value string) string {
	str := os.Getenv(key)
	if len(str) > 0 {
		return str
	}
	return value
}
