package main

import (
	"fmt"
	"strings"
	"flag"
	"github.com/WilliamJPriest/checklist/cmd"

)

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
		cmd.Create()
	}
	if readF{
		cmd.Read()
	}
	if updateF{
		cmd.Update()
	}
	if deleteF{
		cmd.Delete()
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
		cmd.Create()

	case "read":
		cmd.Read()

	case "update":
		cmd.Update()

	case "delete":
		cmd.Delete()

	case "millionaire":
		cmd.Money()

	}

}

