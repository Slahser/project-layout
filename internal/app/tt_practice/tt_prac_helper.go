package tt_practice

import "fmt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else { // a = b or a < b
		min = b
		max = a
	}
	return
}

func Trace(s string) {
	fmt.Println()
	fmt.Println("entering:", s)
}

func Untrace(s string) {
	fmt.Println()
	fmt.Println("leaving:", s)
}

func Function1() {
	fmt.Println()
	fmt.Printf("In function1 at the top\n")
	defer Function2()
	fmt.Printf("In function1 at the bottom!\n")
	return
}

func Function2() {
	fmt.Println("Function2: Deferred until the end of the calling function!")
}

// this function changes reply:
func Multiply(a, b int16, reply *int16) {
	*reply = a * b
}

func Min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}















