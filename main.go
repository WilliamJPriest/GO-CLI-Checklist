package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"flag"
	"path/filepath"

)




var todoList = []Todos{}


type Todos struct {
	id      string
	item   string
	checked bool
}

func getUserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback to the current directory if user home directory cannot be determined
		return "."
	}
	return home
}

var checklistPath = filepath.Join(getUserHomeDir(), "checklists.csv")
var newCheckListPath = filepath.Join(getUserHomeDir(), "checklist_new.csv")

func main() {
	var selectedAction string
	var createF bool
	var readF bool
	var updateF bool
	var deleteF bool

	flag.BoolVar(&createF, "create", false, "skips to the create func")
	flag.BoolVar(&readF, "read", false, "skips to the read func")
	flag.BoolVar(&updateF, "update", false, "skips to the update func")
	flag.BoolVar(&deleteF, "delete", false, "skips to the delete func")
	flag.Parse()

	if createF{
		create()
	}
	if readF{
		read()
	}
	if updateF{
		update()
	}
	if deleteF{
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
	lowerCStr := strings.ToLower((selectedAction))

	switch lowerCStr{

	case "create":
		create()

	case "read":
		read()

	case "update":
		update()

	case "delete":
		delete()

	case "millionaire":
		money()

	}

}


func create() {
	var todoItem string
	t := Todos{}
	fmt.Println("write custom id: Example: mom1")
	fmt.Scanln(&t.id)
	fmt.Println("write your todo:")
    reader := bufio.NewReader(os.Stdin)
    todoItem, _ = reader.ReadString('\n')
    t.item = strings.TrimSpace(todoItem)
    t.checked = false
    todoList = append(todoList, t)

	fmt.Println(todoList)


	csvFile, err := os.OpenFile(checklistPath, os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	for _, eachrecord := range todoList {
		recordSlice := []string{eachrecord.id, eachrecord.item, strconv.FormatBool(eachrecord.checked)}
		if err := csvwriter.Write(recordSlice); err != nil {
			log.Fatalf("error writing to CSV: %s", err)
		}
	}
	read()
}

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
		var checkBox = '☐'
        if eachrecord[2] == "true" {
				checkBox= '☑'
		}
		fmt.Println("id: "+eachrecord[0]+" | " +string(checkBox)+"  "+ eachrecord[1])
	}
}

func update(){
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

    newCSVFile, err := os.Create(newCheckListPath)
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

    if err := os.Rename(newCheckListPath, checklistPath); err != nil {
        log.Fatalf("failed renaming file: %s", err)
    }
	
	read()
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

    newCSVFile, err := os.Create(newCheckListPath)
    if err != nil {
        log.Fatalf("failed creating new file: %s", err)
    }
    defer newCSVFile.Close()

    csvwriter := csv.NewWriter(newCSVFile)

    for _, eachrecord := range records {
        if eachrecord[0] == deleteID {
			fmt.Println("\nDeleted: "+ eachrecord[1]+ "\n")

        } else {
            if err := csvwriter.Write(eachrecord); err != nil {
                log.Fatalf("error writing to CSV: %s", err)
            }
        }
    }

    csvwriter.Flush()


    csvFile.Close()
    newCSVFile.Close()


    if err := os.Rename(newCheckListPath, checklistPath); err != nil {
        log.Fatalf("failed renaming file: %s", err)
    }
	read()

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