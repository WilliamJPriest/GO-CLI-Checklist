package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/WilliamJPriest/checklist/storage"
)

func Annihilate(){
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
		if eachrecord[0] ==eachrecord[0]{
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
		fmt.Println("\nAnnihilated Everything")	
	}else{
		fmt.Println("\nThere is Nothing Left to Annihilate")	
	}
	

}