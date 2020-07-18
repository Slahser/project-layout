package tt_practice

import (
	"fmt"
	"strings"
	"time"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}

/**
Function
*/
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
func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func upPr(p *Pr) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Person struct {
	firstName string
	lastName  string
}

func LongWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e8) // sleep for 0.5 seconds
	fmt.Println("End of longWait()")
}

func ShortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e8) // sleep for 0.2 seconds
	fmt.Println("End of shortWait()")
}

func Pump1(ch chan int) {
	for i := 0; ; i++ {
		if i%4623127 == 0 {
			ch <- i * 2
		}
	}
}

func Pump2(ch chan int) {
	for i := 0; ; i++ {
		if i%4623127 == 0 {
			ch <- i + 5
		}
	}
}
