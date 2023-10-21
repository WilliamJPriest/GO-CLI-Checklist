package storage

import (
	"os"
)

func GetUserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback to the current directory if user home directory cannot be determined
		return "."
	}
	return home
}
