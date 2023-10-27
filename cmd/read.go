package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/WilliamJPriest/checklist/storage"

)

func Read(){

	csvFile,err := os.Open(storage.ChecklistPath)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile) 

	records, err := reader.ReadAll() 
  
    if err != nil  { 
        fmt.Println("Error reading records") 
    } 
    fmt.Println("")

    for _, eachrecord := range records  { 
		var checkBox = '☐'
        if eachrecord[2] == "true" {
				checkBox= '☑'
		}
		fmt.Println("id: "+eachrecord[0]+" | " +string(checkBox)+"  "+ eachrecord[1])
	}
	
}