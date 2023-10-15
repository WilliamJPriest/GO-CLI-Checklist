package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var todoList = []Todos{}

type Todos struct {
	id      int
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
	t := Todos{}

	fmt.Println(`Select Action: 

	Create
	Read
	Update
	Delete

	`)
	fmt.Scanln(&selectedAction)
	if selectedAction == "Create" {
		create(t)
	};
	if selectedAction == "Read" {
		read(selectedAction)
	}
	if selectedAction == "Update" {
		update(selectedAction)
	}
	if selectedAction == "Delete" {
		delete(selectedAction)
	}else{
		println("Error, Ctrl C to reset program")
	}
}

func create(t Todos) {
	t.id++
	t.item = "tacos"
	t.checked = false
	todoList = append(todoList, t)

	fmt.Println(todoList)


	csvFile, err := os.Create("checklist.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	for _, record := range todoList {
		recordSlice := []string{strconv.Itoa(record.id), record.item, strconv.FormatBool(record.checked)}
		if err := csvwriter.Write(recordSlice); err != nil {
			log.Fatalf("error writing to CSV: %s", err)
		}
	}
}

func read(selectedAction string){
	fmt.Println(selectedAction)
}

func update(selectedAction string){
	fmt.Println(selectedAction)
}

func delete(selectedAction string){
	fmt.Println(selectedAction)
}