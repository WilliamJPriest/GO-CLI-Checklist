package create


import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/WilliamJPriest/GO-CLI-Checklist/todos"
	
)

func create() {
	var todoItem string
	t := &todos.Todos{}
	fmt.Println("write custom id: Example: mom1")
	fmt.Scanln(&t.id)
	fmt.Println("write your todo:")
	reader := bufio.NewReader(os.Stdin)
	todoItem, _ = reader.ReadString('\n')
	t.item = strings.TrimSpace(todoItem)
	t.checked = false
	todos.todoList = append(todos.todoList, t)

	fmt.Println(todos.todoList)

	csvFile, err := os.OpenFile("checklist.csv", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	for _, eachrecord := range todos.todoList {
		recordSlice := []string{eachrecord.id, eachrecord.item, strconv.FormatBool(eachrecord.checked)}
		if err := csvwriter.Write(recordSlice); err != nil {
			log.Fatalf("error writing to CSV: %s", err)
		}
	}
}