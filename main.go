package main

import "fmt"

var todoList = []Todos{}

type Todos struct {
	id      int
	items   string
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

func create(t Todos){
	fmt.Println(t)
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