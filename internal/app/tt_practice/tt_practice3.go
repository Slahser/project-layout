package tt_practice

import (
	"fmt"
	"math"
)

type Pr Person

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

type Buffer struct {
	tt int
}

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

/**
接口嵌套
*/
type File interface {
	ReadWrite
	Lock
	Close()
}

type Any interface{}

type TopologicalGenus interface {
	Rank() int
}

func (sq *Square) Rank() int {
	return 1
}

func (r Rectangle) Rank() int {
	return 2
}

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	// ...
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

func Tt3() {
	Trace("tt_practice3#Tt3()")
	defer Untrace("tt_practice3#Tt3()")

	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	ms2 := &struct1{10, 15.5, "Chris2"}

	fmt.Println(ms)  //结构体
	fmt.Println(ms2) //结构体
	fmt.Println(&ms) //内存地址
	fmt.Println(*ms) //结构体内存实际存储数据

	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Pr)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	upPr(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)

	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

	//类型断言
	if v, ok := o.(valuable); ok { // checked type assertion
		fmt.Printf("类型断言 %s", v)
	}

	fmt.Println()
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if _, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is Circle")
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	//类型断言
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	fmt.Println()

	//接口提取

	topgen := []TopologicalGenus{r, q}

	//topoMap := make(map[int]TopologicalGenus, len(topgen))
	//keys := make([]int, len(topgen))
	//sort.Ints(keys)

	fmt.Println("Looping through topgen for rank ...")

	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}

	//正交与动态类型
	b := new(Bird)
	DuckDance(b)

}
