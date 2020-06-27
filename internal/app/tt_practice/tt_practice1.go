package tt_practice

import (
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/pkg/tpkg"
	structrueLog "github.com/sirupsen/logrus"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type AudioOutput int

const (
	OutMute   AudioOutput = iota // 0
	OutMono                      // 1
	OutStereo                    // 2
	_
	_
	OutSurround // 5
)

func Tt1() {

	Trace("tt_practice1#Tt1()")
	defer Untrace("tt_practice1#Tt1()")

	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println(os.Environ())

	fmt.Println(tpkg.Pi)

	fmt.Println("tt")
	fmt.Println(OutSurround)

	var goos = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)

	t1, t2 := 1, "here t2"
	t1Next := t1

	fmt.Println(&t1)
	fmt.Println(&t1Next == &t1) //值类型
	fmt.Println(&t2)

	timens := int64(time.Now().Nanosecond())
	fmt.Println(timens)
	rand.Seed(timens)
	for i := 0; i < 5; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}

	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)
	fmt.Println()
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)

	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v", c1)

	str := "This is an example of a string"
	fmt.Println()
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	orig := "Hey, how are you George?"
	fmt.Printf("The original string is: %s\n", orig)
	lower := strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper := strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)

	orig2 := "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig2) //利用多返回值的特性，这些函数会返回 2 个值，第 1 个是转换后的结果（如果转换成功），第 2 个是可能出现的错误
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	_, _ = time.Parse("2006-01-02 15:04:05", "2020-06-27 01:04:36") //忽略错误

	var i1 = 5555
	//这个地址可以存储在一个叫做指针的特殊数据类型中
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	//一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关
	var intP *int = &i1
	//var p *type
	fmt.Println(intP)
	//符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；这被称为反引用
	fmt.Println(*intP)
	*intP = 6666
	//指针转移 当然可以继续操作
	fmt.Println(*intP)
	fmt.Println(i1)
	//指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝

	fmt.Println(Abs(-233))

	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig is not an integer - error occur\n")
	}

	num1 := 100
	switch num1 {
	case 0: // 空分支，只有当 i == 0 时才会进入分支
	case 1:
		fmt.Println("It's equal to 1") // 当 i == 0 时函数不会被调用

	case 2:
		fallthrough
	case 3:
		fmt.Println("It's equal to 2 or 3") // 当 i == 0 时函数也会被调用

	case 98, 99:
		fmt.Println("It's equal to 98 or 99")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	//switch {
	//case num1 > 98 && num1 < 99:
	//	fmt.Println("It's equal to 98 or 99")
	//case num1 > 100:
	//	fmt.Println("It's equal to 100")
	//}
	//
	//switch a, b := x[i], y[j] {
	//case a < b:
	//	t = -1
	//case a == b:
	//	t = 0
	//case a > b:
	//	t = 1
	//}

	var i = 5

	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

	i1, _, f1 := ThreeValues()
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)

	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)

	//闭包演示.其实比较勉强,没有函数返回,只是定义
	n = 0
	reply := &n //n的指针
	Multiply(10, 5, reply)
	fmt.Println("Multiply *reply:", *reply) // Multiply: 50
	fmt.Println("Multiply n:", n)           // Multiply: 50

	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = Min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)

	//关键字 defer 允许我们进行一些函数执行完成后的收尾工作
	Function1()

	Dd1()
	Dd2()

	//记录出入参
	_, _ = Funttt("funtt")

	//callback
	Callback(1, Add)

	//匿名函数
	rangePlus := func(rr int) {
		sum := 0
		for i := 1; i <= rr; i++ {
			sum += i
		}
		fmt.Println(sum)
	}
	rangePlus(100)

	//直接调用匿名函数
	func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()

	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	fmt.Println()
}

/**
闭包,在函数作为返回值的情况,对上级函数作用域内变量的影响
*/
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func Dd1() {
	i := 0
	defer fmt.Println(i) //此处是只打印一次0
	i++
	return
}

func Dd2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("i will write next line and you will reverse")
}

func Funttt(s string) (n int, err error) {
	fmt.Println()
	defer func() {
		log.Printf("Funttt(%q) = %d, %v\n", s, n, err)
		fmt.Println()
		structrueLog.WithFields(
			structrueLog.Fields{
				"timestamps": time.Now().Format("2006-01-02 15:04:05"),
				"wtf":        "yeah",
				"defer":      "yes",
				"param":      s,
				"return":     n,
			}).Info("tt defer log")
	}()
	return 7, io.EOF
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func Callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}
