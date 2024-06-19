package test

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	[main.exe --add "kuyhee"]   ==> {"id": 2, "text": "kuyhee"}
	main.exe --get 2          ==> {"id": 2, "text": "kuyhee"}
	main.exe --rm 2
	main.exe --update  2 "heekuy"  ==> {"id": 2, "text": "heekuy"}
	main.exe                  ==> []
*/

func Test2() {
	args := os.Args
	user := make(map[string]string)
	user[args[2]] = args[3]

	if len(args) == 1 {
		fmt.Println("Not Expect")
		return
	}

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
	}

	if len(args) > 2 {
		if args[1] == "--add" {
			b, err := json.Marshal(args[2])
			if err != nil {
				panic(err)
			}
			_ = b

			//os.WriteFile("todo")

		}

		if args[1] == "--get" {

		}

		if args[1] == "--rm" {

		}

		if args[1] == "--update" {

		}
	}

}
