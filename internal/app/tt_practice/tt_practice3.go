package tt_practice

import (
	"fmt"
)


type Pr Person

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field
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

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
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
	(*pers2).lastName = "Woodward"  // 这是合法的
	upPr(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris","Woodward"}
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
}
