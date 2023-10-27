package storage

import (
	"os"
	"path/filepath"
)

func GetUserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		
		return "."
	}
	return home
}

var ChecklistPath = filepath.Join(GetUserHomeDir(), "checklists.csv")
var NewCheckListPath = filepath.Join(GetUserHomeDir(), "checklist_new.csv")
