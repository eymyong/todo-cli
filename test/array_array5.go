package test

import "fmt"

func ArrAndArr() {

	arr1 := []string{"Alice", "Bob", "John", "Jane"}
	arr2 := []string{"John", "Foobar", "Barbaz", "Foobaz", "Bob"}
	result := duplicate(arr1, arr2)
	fmt.Println(result)

}

func duplicate(arr1, arr2 []string) []string {
	newArr := []string{}
	//m := make( map[string]bool)

	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1 == v2 {
				newArr = append(newArr, v1)
			}
		}
	}
	return newArr
}
