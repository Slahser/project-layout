package tt_practice

import (
	"fmt"
	"math"
	"strconv"
)

/**
Method
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
*/

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string {
	b.thing = 2
	return fmt.Sprint(b)
}

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

type Point struct {
	x, y float64
}

type NamedPoint struct {
	Point
	name string
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func Tt4() {
	Trace("tt_practice4#Tt4()")
	defer Untrace("tt_practice4#Tt4()")

	var ti *TwoInts = new(TwoInts)
	ti.a = 1
	ti.b = 2

	fmt.Printf("AddThem: %s ", strconv.Itoa(ti.AddThem()))
	fmt.Printf("AddToParam: %s ", strconv.Itoa(ti.AddToParam(3)))
	fmt.Printf("toString: %s ", ti.String())
	fmt.Println()

	//change()接受一个指向 B 的指针，并改变它内部的成员；write() 通过拷贝接受 B 的值并只输出 B 的内容。
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
	fmt.Println()

	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)
	fmt.Println()

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
	fmt.Println()

	//隐藏类型
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印5
}
