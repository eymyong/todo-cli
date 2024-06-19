package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

/*
[main.exe --add "kuyhee"]   ==> {"id": 2, "text": "kuyhee"}
main.exe --get 2          ==> {"id": 2, "text": "kuyhee"}
main.exe --rm 2
main.exe --update  2 "heekuy"  ==> {"id": 2, "text": "heekuy"}
main.exe                  ==> []
*/
type Mode string

const (
	ModeAdd    Mode = "--add"
	ModeGetAll Mode = "--get all"
	ModeGet    Mode = "--get"
	ModeUpdate Mode = "--update"
	ModeRemove Mode = "--rm"
)

type job struct {
	id   string
	data string
	mode Mode
}

type todo struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func Test() {
	args := os.Args
	job, err := parse(args)
	if err != nil {
		panic(err)
	}
	fileName := "todo.json"
	switch job.mode {
	case ModeAdd:
		add(fileName, job.data)
	case ModeGetAll:
		get(fileName, job.id)
	case ModeGet:
		get(fileName, job.id)
	case ModeRemove:
		rm(fileName, job.id)
	case ModeUpdate:
		update(fileName, job.id, job.data)
	}
	if err != nil {
		panic("Incorrect Mode")
	}

}

func parse(args []string) (job, error) {
	if len(args) == 1 {
		return job{}, errors.New("Not order")
	}

	if len(args) == 2 {
		if args[1] == "--add" {
			return job{}, errors.New("There is no information to add")
		}
		if args[1] == "--get" {
			return job{}, errors.New("There is no information to get")
		}
		if args[1] == "--update" {
			return job{}, errors.New("There is no information to update")
		}
		if args[1] == "--rm" {
			return job{}, errors.New("There is no information to rm")
		}
	}

	if len(args) == 3 {
		if args[1] == "--add" {
			return job{mode: ModeAdd, data: args[2]}, nil
		}
		if args[1] == "--get" {
			if args[2] == "all" {
				return job{mode: ModeGetAll, id: args[2]}, nil
			}
			return job{mode: ModeGet, id: args[2]}, nil
		}
		if args[1] == "--rm" {
			return job{mode: ModeRemove, id: args[2]}, nil
		}
		if args[1] == "--update" {
			fmt.Println("Not data to update")
			return job{}, nil
		}

	}

	if len(args) == 4 {
		if args[1] == "--update" {
			return job{mode: ModeUpdate, id: args[2], data: args[3]}, nil
		}
	}

	return job{}, errors.New("Input incorrect")

}

func update(fileName, id, data string) {
	b, err := readFile(fileName)
	if err != nil {
		panic(err)
	}

	if len(b) == 0 {
		fmt.Println("No Data in File")
		return
	}

	listTodo, err := unmarshal(b)
	if err != nil {
		panic(err)
	}
	newListTodo := []todo{}
	var checkID bool
	for _, todo := range listTodo {
		if id == todo.Id {
			checkID = true
			todo.Data = data
			newListTodo = append(newListTodo, todo)
			continue
		}
		newListTodo = append(newListTodo, todo)
	}

	if checkID != true {
		fmt.Printf("Not found ID: %s", id)
		return
	}

	newByte, err := marshal(newListTodo)
	if err != nil {
		panic(err)
	}

	err = writeFile(fileName, newByte)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID to Update = ID: %s\nData = %s\nSucceed", id, data)
}

func rm(fileName, id string) {
	b, err := readFile(fileName)
	if err != nil {
		panic(err)
	}

	if len(b) == 0 {
		fmt.Println("No Data in File")
		return
	}

	listTodo, err := unmarshal(b)
	if err != nil {
		panic(err)
	}

	var newListTodo []todo
	var checkID bool
	for _, todo := range listTodo {
		if todo.Id == id {
			checkID = true
			continue
		}
		newListTodo = append(newListTodo, todo)
	}

	if checkID != true {
		fmt.Printf("Not found ID: %s", id)
		return
	}

	fmt.Printf("rm to ID = ID: %s\nSucceed", id)
	newByte, err := marshal(newListTodo)
	if err != nil {
		panic(err)
	}

	err = writeFile(fileName, newByte)
	if err != nil {
		panic(err)
	}
}

func add(fileName string, data string) {
	b, err := readFile(fileName)
	if err != nil {
		panic(err)
	}
	newTodoList := []todo{}
	newTodo := todo{Id: uuid.NewString(), Data: data}

	if len(b) == 0 {
		newTodoList = append(newTodoList, newTodo)

		newByte, err := marshal(newTodoList)
		if err != nil {
			panic(err)
		}

		err = writeFile(fileName, newByte)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Add Data: %s", data)
		return
	}

	if len(b) >= 0 {
		todoList, err := unmarshal(b)
		if err != nil {
			panic(err)
		}

		todoList = append(todoList, newTodo)

		newByte, err := marshal(todoList)
		if err != nil {
			panic(err)
		}

		err = writeFile(fileName, newByte)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Data to add = %s", data)
	}

}

func get(fileName string, id string) {
	b, err := readFile(fileName)
	if err != nil {
		panic(err)
	}

	if len(b) == 0 || string(b) == "nil" {
		fmt.Println("No Data in File")
		return
	}

	listTodo, err := unmarshal(b)
	if err != nil {
		panic(err)
	}

	if id == "all" {
		fmt.Printf("All Data\n%s", listTodo)
		return
	}

	c := checkID(listTodo, id)
	if c != "" {
		fmt.Printf("ID: %s\nData: %s", id, c)
	}
}

func checkID(listTodo []todo, id string) string {
	for _, v := range listTodo {
		if v.Id == id {
			return v.Data
		}
	}
	fmt.Printf("Not found ID: %s", id)
	return ""
}

// ===============================================================================================

func readFile(fileName string) ([]byte, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return []byte{}, errors.New("Unable to Readfile")
	}

	return b, nil
}

func writeFile(fileName string, data []byte) error {
	err := os.WriteFile(fileName, data, 0664)
	if err != nil {
		return errors.New("Unable to Writefile")
	}
	return nil
}

func marshal(listTodo []todo) ([]byte, error) {
	b, err := json.Marshal(listTodo)
	if err != nil {
		return []byte{}, errors.New("Unable to Marshal")
	}
	return b, nil
}

func unmarshal(arrByte []byte) ([]todo, error) {
	t := []todo{}
	err := json.Unmarshal(arrByte, &t)
	if err != nil {
		return []todo{}, errors.New("Unable to Unmarshal")
	}

	return t, nil
}

// func toLine(todo todo) (string, error) {
// 	if todo.Id == "" {
// 		return "", errors.New("empty todo Id")
// 	}

// 	return fmt.Sprintf("%s: %s", todo.Id, todo.Data), nil
// }

// // TODO: change input type
// func toLines(todos []todo) (string, error) {
// 	s := ""
// 	for _, t := range todos {
// 		line, err := toLine(t)
// 		if err != nil {
// 			return "", err
// 		}

// 		s += line + "\n"
// 	}

// 	return s, nil
// }

// [{"name":"pak","id": 3,"Age": 10},{"name":"yong","Age": 15},{"name": "art","Age": 20}]
// func fromLine(line string) (todo, error) {
// 	words := strings.Split(line, ":")
// 	if len(words) != 2 {
// 		return todo{}, fmt.Errorf("invalid line format: %s", line)
// 	}

// 	t := todo{
// 		Id:   words[0],
// 		Data: words[1],
// 	}
// 	fmt.Println("Id =", t.Id)
// 	fmt.Println("Data = ", t.Data)
// 	fmt.Println("todo = ", t)

// 	return t, nil
// }

// func fromLines(linesStr string) (map[string]todo, error) {
// 	lines := strings.Split(linesStr, "\n")
// 	todos := make(map[string]todo)
// 	for _, line := range lines {
// 		t, err := fromLine(line)
// 		if err != nil {
// 			return nil, err
// 		}

// 		todos[t.Id] = t
// 	}

// 	return todos, nil
// }

// func fromMapToList(todos map[string]todo) []todo {
// 	result := []todo{}
// 	for k := range todos {
// 		result = append(result, todos[k])
// 	}

// 	return result
// }
