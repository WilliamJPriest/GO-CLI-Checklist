package todos

var todoList = []Todos{}


type Todos struct {
	id      string
	item   string
	checked bool
}