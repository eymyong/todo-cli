package foo

import (
	"fmt"
	"os"
	"strings"
)

func Foo() {
	b, err := os.ReadFile("todo")
	if err != nil {
		panic(err)
	}

	dataStr := string(b)
	//fmt.Println(dataStr)
	lines := strings.Split(dataStr, "\n")
	fmt.Println(lines[0])
	fmt.Println(lines[1])

	for i := range lines {
		line := strings.TrimSpace(lines[i])

		fmt.Println("before", line)
		words := strings.Split(line, ":")
		words[0] = strings.TrimSpace(words[0])
		words[1] = strings.TrimSpace(words[1])
		fmt.Println("after", words, "len", len(words))

		fmt.Println("++++++")

		// lines[i] = line
		// fmt.Println(line)
		// fmt.Println(words)
	}

}
