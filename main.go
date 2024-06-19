package main

import "github.com/eymyong/yong2-go/test"

/*
	[main.exe --add "kuyhee"]   ==> {"id": 2, "text": "kuyhee"}
	main.exe --get 2          ==> {"id": 2, "text": "kuyhee"}
	main.exe --rm 2
	main.exe                  ==> []
*/

func main() {
	test.Test()
	//foo.Foo()
}
