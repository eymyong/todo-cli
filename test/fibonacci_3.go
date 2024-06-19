package test

import "fmt"

//1, 2, 3, 5, 8, 13, 21, ...
// fib(2); // 2
// fib(3); // 3
// fib(4); // 5 ======
// fib(5); // 8
// fib(6); // 13

func TestFBN() {
	arr := []int{1, 2, 3, 5, 8, 13, 21}
	_ = arr
	result := fib(6) // 13
	fmt.Println(result)
}

func fib(n int) int { // ทำไม่ได้
	//aa := 0
	for i := 1; i <= 15; i += i {
		for j := 2; j <= 30; j++ {

		}

	}

	// fmt.Printf("%d ", i)
	return (n - 1) + (n - 2)
}
