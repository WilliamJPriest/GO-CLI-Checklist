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
)

// const (
// 	Create string = "Create",
// 	Read   string = "Read",
// )

var todoList = []Todos{}


type Todos struct {
	id      string
	item   string
	checked bool
}

// type TodoListFuncs interface {
// 	create()
// 	read()
// 	update()
// 	delete()
// }

func main() {
	var selectedAction string
	var createF string
	var readF string
	flag.StringVar(&createF, "create", "", "skips to the create func")
	flag.StringVar(&readF, "read", "", "skips to the read func")
	flag.Parse()

	if createF == "a"{
		create()
	}
	if readF == "a"{
		read()
	}
	
	fmt.Println(`
	
	Select Action: 

	Create
	Read
	Update
	Delete
	Become a Millionaire

	`)
	fmt.Scanln(&selectedAction)
	if selectedAction == "Create" {
		create()
	};
	if selectedAction == "Read" {
		read()
	}
	if selectedAction == "Update" {
		update(selectedAction)
	}
	if selectedAction == "Delete" {
		delete(selectedAction)
	}
	if selectedAction == "Millionaire" {
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


	csvFile, err := os.OpenFile("checklist.csv", os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	for _, record := range todoList {
		recordSlice := []string{record.id, record.item, strconv.FormatBool(record.checked)}
		if err := csvwriter.Write(recordSlice); err != nil {
			log.Fatalf("error writing to CSV: %s", err)
		}
	}
}

func read(){
	csvFile,err := os.Open("checklist.csv")
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

func update(selectedAction string){
	fmt.Println(selectedAction)
}

func delete(selectedAction string){
	fmt.Println(selectedAction)
}

func money(){
	fmt.Println("ttt")
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