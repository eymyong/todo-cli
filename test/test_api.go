package test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	[main.exe --add "kuyhee"]   ==> {"id": 2, "text": "kuyhee"}
	main.exe --get 2          ==> {"id": 2, "text": "kuyhee"}
	main.exe --rm 2
	main.exe --update  2 "heekuy"  ==> {"id": 2, "text": "heekuy"}
	main.exe                  ==> []
*/

func TestApi() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println("Not expect")
		return
	}
	// 1
	// list
	if len(args) == 1 {
		b, err := os.ReadFile("todo")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
		return
	}
	// ==2 if not order
	if len(args) == 2 {
		if args[1] == "--add" {
			fmt.Println("Not information to add")
			return
		}
		if args[1] == "--get" {
			fmt.Println("Not information to get")
			return
		}
		if args[1] == "--rm" {
			fmt.Println("Not information to rm")
			return
		}
		if args[1] == "--update" {
			fmt.Println("Not information to update")
			return
		}
		fmt.Println("not enough args")
		return
	}

	if len(args) > 2 {
		if args[1] == "--add" {
			add("todo", args[2])
			return
		}
		if args[1] == "--get" {
			if args[2] == "all" {
				getAll := readFile("todo")
				fmt.Println("getAll: \n", getAll)
				return
			}
			id, val := get("todo", args[2])
			fmt.Printf("id:%s = %s", id, val)
			return
		}
		if args[1] == "--rm" {
			rm(args[2])
			return
		}

		if args[1] == "--update" {
			if len(args) == 3 {
				fmt.Println("No data to Update")
				return
			}
			update(args[2], args[3])
			return
		}

		fmt.Println("wrong command (need: --add,--get,--rm)")
		return
	}

}

func rm(id string) {
	dataStr := readFile("todo")
	if dataStr == "" {
		fmt.Println("Not data")
		return
	}

	lines := strings.Split(dataStr, "\n")
	var newdata string
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		words := strings.Split(line, ":")
		words[0] = strings.TrimSpace(words[0])
		words[1] = strings.TrimSpace(words[1])

		if id != words[0] {
			if len(newdata) == 0 {
				newdata = newdata + lines[i]
				continue
			}
			if len(newdata) > 0 {
				newdata = newdata + "\n" + lines[i]
				continue
			}
		}

		fmt.Println("You Delete: ", id)
	}

	err := os.WriteFile("todo", []byte(newdata), 0664)
	if err != nil {
		panic("connot WriteFile")
	}
	fmt.Println("Completely deleted")

}

func update(id, newValue string) {
	dataStr := readFile("todo")
	if dataStr == "" {
		fmt.Println("Not Data")
		return
	}
	lines := getLines(dataStr)
	var newdata string
	for i := range lines {
		words := strings.Split(lines[i], ":")
		words[0] = strings.TrimSpace(words[0])
		words[1] = strings.TrimSpace(words[1])
		if id != words[0] {
			if len(newdata) == 0 {
				newdata = newdata + lines[i]
				continue
			}
			if len(newdata) > 0 {
				newdata = newdata + "\n" + lines[i]
				continue
			}
		}

		if id == words[0] {
			if len(newdata) == 0 {
				words[1] = newValue
				newdata = newdata + id + ": " + words[1]
				fmt.Println("Update ", id)
				continue
			}
			if len(newdata) > 0 {
				words[1] = newValue
				newdata = newdata + "\n" + id + ": " + words[1]
				fmt.Println("Update ", id)
				continue
			}
		}
	}
	err := os.WriteFile("todo", []byte(newdata), 0664)
	if err != nil {
		panic(err)
	}
	fmt.Println("Completely Update")

}

func get(fileName, id string) (string, string) {
	dataStr := readFile(fileName)

	lines := getLines(dataStr)
	_, val := getValue(id, lines)

	return id, val
}

// func removeStr(s string) string {
// 	return ""
// }

func getLines(data string) []string {
	lines := strings.Split(data, "\n")
	//fmt.Printf("get lines: %s\n", lines)
	return lines
}

func getValue(id string, lines []string) (string, string) {

	for i := range lines {
		words := strings.Split(lines[i], ":")
		words[0] = strings.TrimSpace(words[0])
		words[1] = strings.TrimSpace(words[1])
		if id == words[0] {
			return words[0], words[1]
		}
	}
	return "", "Not ID"
}

func add(fileName, input string) {
	dataStr := readFile(fileName)
	// if err != nil {
	// 	if !errors.Is(err, errNotFound) {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// write new data
	// }
	//var data string

	// if dataStr == "" {
	// 	dataStr = input + ": " + args[3]
	// }

	lines := strings.Split(dataStr, "\n")
	var numberStr string
	for i := range lines {
		words := strings.Split(lines[i], ":")
		words[0] = strings.TrimSpace(words[0])
		numberStr = words[0]
	}
	if len(numberStr) == 0 {
		numberStr = "1"
	}

	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		panic("not number")
	}

	newID := fmt.Sprintf("%d", numberInt+1)
	line := newID + ": " + input

	var addData string
	if len(dataStr) > 0 {
		addData = dataStr + "\n" + line
	}

	if len(dataStr) == 0 {
		dataStr = dataStr + line
		addData = dataStr
	}

	err = os.WriteFile(fileName, []byte(addData), 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("add %s succes", input)
}

// var (
// 	errNotFound = errors.New("empty data")
// )

func readFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	dataStr := string(b)
	return dataStr
}
