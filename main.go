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
	var cleanF bool
	var annihilateF bool

	flag.BoolVar(&createF, "create", false, "skips to the create func")
	flag.BoolVar(&readF, "read", false, "skips to the read func")
	flag.BoolVar(&updateF, "update", false, "skips to the update func")
	flag.BoolVar(&deleteF, "delete", false, "skips to the delete func")
	flag.BoolVar(&cleanF, "clean", false, "skips to the delete func")
	flag.BoolVar(&annihilateF, "annihilate", false, "skips to the annihilate func")
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
	if cleanF{
		cmd.Clean()
	}
	if annihilateF{
		cmd.Annihilate()
	}
	
	fmt.Println(`
	
	Select Action: 

	Create
	Read
	Update
	Delete
	Clean
	Annihilate
	Millionaire
	Exit

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

	case "clean":
		cmd.Clean()

	case "annihilate":
		cmd.Annihilate()
		
	case "millionaire":
		cmd.Money()
	
	case "exit":
		cmd.Exit()
	}

}

