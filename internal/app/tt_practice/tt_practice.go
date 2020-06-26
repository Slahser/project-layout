package tt_practice

import (
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/pkg/tpkg"
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

func Tt() {

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
		fmt.Print("%2.2f / ", 100*rand.Float32())
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
	*intP=6666
	//指针转移 当然可以继续操作
	fmt.Println(*intP)
	fmt.Println(i1)
	//指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝

}
