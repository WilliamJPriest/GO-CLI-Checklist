package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/WilliamJPriest/checklist/storage"

)

func Update(){
	var updateID string
	fmt.Println("\nWrite id of the task you wish to update as complete")
	fmt.Scanln(&updateID)

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

    for _, eachrecord := range records {
        if eachrecord[0] == updateID {
            recordSlice := []string{eachrecord[0], eachrecord[1], strconv.FormatBool(true)}
			fmt.Println("\nUpdated: "+ eachrecord[1]+ "\n")
				
            if err := csvwriter.Write(recordSlice); err != nil {
                log.Fatalf("error writing to CSV: %s", err)
            }
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
	Read()
}