package test

import "fmt"

func Foo() {

	arr := []int{12, -13, 14, 4, 2, -1, -18}
	_ = arr
	arr2 := []int{}
	result1, result2 := maxNegMinPos(arr2)

	fmt.Println("maxNeg = \t", result1)
	fmt.Println("minPos = \t", result2)
}

// MaxNeg is -1
// MinPos is 2
func maxNegMinPos(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Not Value in array")
	}

	numberLobMax := 0
	numberBuakMin := 0

	for _, v := range arr {
		if v < numberLobMax {
			numberLobMax = v
		}

		if v > numberBuakMin {
			numberBuakMin = v
		}
	}

	for _, vv := range arr {
		if vv >= numberLobMax && vv < 0 {
			numberLobMax = vv
		}

		if vv <= numberBuakMin && vv > 0 {
			numberBuakMin = vv
		}
	}

	return numberLobMax, numberBuakMin
}
