package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/WilliamJPriest/checklist/storage"
)

var todoList = []Todos{}


type Todos struct {
	id      string
	item   string
	checked bool
}

func Create() {
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

	fmt.Println("\n" + t.id + " | ☐  " + t.item + "\n")

	csvFile, err := os.OpenFile(storage.ChecklistPath, os.O_CREATE|os.O_APPEND, 0644)
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
	Read()
}