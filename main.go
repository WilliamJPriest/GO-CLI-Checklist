package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/WilliamJPriest/GO-CLI-Checklist/cmd"
)


var checklistPath string = "checklist.csv"


func main() {
	var selectedAction string
	var createF string
	var readF string
	var updateF string
	var deleteF string

	flag.StringVar(&createF, "create", "", "skips to the create func")
	flag.StringVar(&readF, "read", "", "skips to the read func")
	flag.StringVar(&updateF, "update", "", "skips to the update func")
	flag.StringVar(&deleteF, "delete", "", "skips to the delete func")
	flag.Parse()

	if createF == "a"{
		cmd.create()
	}
	if readF == "a"{
		read()
	}
	if updateF == "a"{
		update()
	}
	if deleteF == "a"{
		delete()
	}
	
	fmt.Println(`
	
	Select Action: 

	Create
	Read
	Update
	Delete
	Millionaire

	`)
	fmt.Scanln(&selectedAction)
	if selectedAction == "Create" {
		cmd.create()
	};
	if selectedAction == "Read" {
		read()
	}
	if selectedAction == "Update" {
		update()
	}
	if selectedAction == "Delete" {
		delete()
	}
	if selectedAction == "Millionaire" {
		money()
	}
}


// func create() {
	// var todoItem string
	// t := Todos{}
	// fmt.Println("write custom id: Example: mom1")
	// fmt.Scanln(&t.id)
	// fmt.Println("write your todo:")
    // reader := bufio.NewReader(os.Stdin)
    // todoItem, _ = reader.ReadString('\n')
    // t.item = strings.TrimSpace(todoItem)
    // t.checked = false
    // todoList = append(todoList, t)

	// fmt.Println(todoList)


	// csvFile, err := os.OpenFile("checklist.csv", os.O_CREATE | os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }
	// defer csvFile.Close()

	// csvwriter := csv.NewWriter(csvFile)
	// defer csvwriter.Flush()

	// for _, eachrecord := range todoList {
	// 	recordSlice := []string{eachrecord.id, eachrecord.item, strconv.FormatBool(eachrecord.checked)}
	// 	if err := csvwriter.Write(recordSlice); err != nil {
	// 		log.Fatalf("error writing to CSV: %s", err)
	// 	}
	// }
// }

func read(){
	csvFile,err := os.Open(checklistPath)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile) 

	records, err := reader.ReadAll() 
  
    if err != nil  { 
        fmt.Println("Error reading records") 
    } 
      
    for _, eachrecord := range records  { 
        fmt.Println(eachrecord) 
    } 

}

func update() {
	var updateID string
	fmt.Println("Write id of the task you wish to update as complete")
	fmt.Scanln(&updateID)

    csvFile, err := os.Open(checklistPath)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer csvFile.Close()

    reader := csv.NewReader(csvFile)

    records, err := reader.ReadAll()

    if err != nil {
        fmt.Println("Error reading records")
    }

    newCSVFile, err := os.Create("checklist_new.csv")
    if err != nil {
        log.Fatalf("failed creating new file: %s", err)
    }
    defer newCSVFile.Close()

    csvwriter := csv.NewWriter(newCSVFile)

    for _, eachrecord := range records {
        if eachrecord[0] == updateID {
            recordSlice := []string{eachrecord[0], eachrecord[1], strconv.FormatBool(true)}

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

    if err := os.Rename("checklist_new.csv", checklistPath); err != nil {
        log.Fatalf("failed renaming file: %s", err)
    }
}

func delete(){
	var deleteID string
	fmt.Println("Write id of the task you wish to delete")
	fmt.Scanln(&deleteID)
	csvFile, err := os.Open(checklistPath)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer csvFile.Close()

    reader := csv.NewReader(csvFile)

    records, err := reader.ReadAll()

    if err != nil {
        fmt.Println("Error reading records")
    }

    newCSVFile, err := os.Create("checklist_new.csv")
    if err != nil {
        log.Fatalf("failed creating new file: %s", err)
    }
    defer newCSVFile.Close()

    csvwriter := csv.NewWriter(newCSVFile)

    for _, eachrecord := range records {
        if eachrecord[0] == deleteID {
            fmt.Println("Deleted")

        } else {
            if err := csvwriter.Write(eachrecord); err != nil {
                log.Fatalf("error writing to CSV: %s", err)
            }
        }
    }

    csvwriter.Flush()


    csvFile.Close()
    newCSVFile.Close()


    if err := os.Rename("checklist_new.csv", checklistPath); err != nil {
        log.Fatalf("failed renaming file: %s", err)
    }
}

func money(){
	fmt.Println(`

	⠀⠀⠀⠀⣀⣀⣀⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣀⡾⠧⠀⠀⠥⢀⡀⠀⠀
⠀⠀⠀⢀⣴⠋⠁⠀⠀⠀⠀⠀⠀⠀⠑⡄
⠀⠀⢠⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠁
⠀⢀⣻⠁⠀⠀⠀⣰⢿⠀⠸⣽⣗⠖⠃⠀
⠀⠸⢼⠀⠀⠀⠀⣗⢽⠀⠄⠀⠁⠀⠀⠀
⠀⢸⠝⡆⠀⠀⠀⠈⠛⠃⠰⠤⢀⠀⠀⠀
⠀⠀⢯⠜⠦⡀⠀⠀⠀⠀⠀⠀⠀⠉⢂⠀
⠀⠀⠀⠓⢎⣝⠕⣲⡆⠀⡀⠀⠀⠀⠀⠆
⠀⠀⠀⠀⣄⠈⢙⢕⡇⠀⣿⡆⠀⠀⠀⢸
⠀⣠⠔⠉⠈⠑⠴⢬⡇⠀⡷⠃⠀⠀⠀⡈
⠸⡡⢓⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⠁
⠀⠈⠫⣎⡝⡢⢤⣀⠀⠀⣀⣀⡤⡾⠃⠀
⠀⠀⠀⠀⠉⠚⣔⣿⣤⣤⡽⠓⠉⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠘⠛⠛⠋⠀⠀⠀`)
}