package test

import "fmt"

func TestDraw() {

	// drow1(5)
	draw2(5)

	// result := draw3(5)
	// fmt.Println("draw2 = ", result)
	// fmt.Println("----")
	// drow3(5)

}

func draw1(n int) {
	s := ""
	for i := 1; i <= n; i++ {
		s += "*"
		fmt.Println("draw1 = ", s)
	}
}

func draw2(n int) { // ยังทำไม่ได้ ต้องให้ * เรียกจจากมากไปน้อย
	s := ""
	for i := 1; i <= n; i++ {
		s = "*"
	}
	fmt.Println(s)
	fmt.Println("---")

	// s2 := s
	for i := len(s); i >= 1; i-- {

		fmt.Println(i)
		//fmt.Println(s2)

	}

}

func draw3(n int) string {
	s := ""
	for i := 1; i <= n; i++ {
		s += "*"
	}
	return s
}

func draw4(n int) {
	s := ""
	for i := 1; i <= n; i++ {
		s += "*"
	}

	for i := 1; i <= n; i++ {
		fmt.Println("draw3 = ", s)
	}
}
