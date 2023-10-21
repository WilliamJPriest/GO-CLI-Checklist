package storage

import (
	"os"
	"path/filepath"
)

func GetUserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback to the current directory if user home directory cannot be determined
		return "."
	}
	return home
}

var ChecklistPath = filepath.Join(GetUserHomeDir(), "checklists.csv")
var NewCheckListPath = filepath.Join(GetUserHomeDir(), "checklist_new.csv")
