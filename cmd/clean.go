package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/WilliamJPriest/checklist/storage"
)

func Clean() {
	Read()
	csvFile, err := os.Open(storage.ChecklistPath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	newCSVFile, err := os.Create(storage.NewCheckListPath)
	if err != nil {
		log.Fatalf("failed creating new file: %s", err)
	}
	defer newCSVFile.Close()

	csvwriter := csv.NewWriter(newCSVFile)

	var found bool
	for _, eachrecord := range records {
		if eachrecord[2] == "true" {
			found= true
		} else {
			if err := csvwriter.Write(eachrecord); err != nil {
				log.Fatalf("error writing to CSV: %s", err)
			}
		}
	}

	csvwriter.Flush()

	csvFile.Close()
	newCSVFile.Close()

	if err := os.Rename(storage.NewCheckListPath, storage.ChecklistPath); err != nil {
		log.Fatalf("failed renaming file: %s", err)
	}
	if found == true{
		fmt.Println("\nCleaned all completed tasks")	
	}else{
		fmt.Println("\nMust complete tasks before you clean them")	
	}
	

}