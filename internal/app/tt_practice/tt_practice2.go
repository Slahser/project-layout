package tt_practice

import (
	"fmt"
	"sort"
)

func Tt2() {

	Trace("tt_practice2#Tt1()")
	defer Untrace("tt_practice2#Tt1()")

	//数组
	a := [...]string{"a", "b", "c", "d"}
	//切片
	b := []string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i]+b[i])
	}

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) //capacity代表从起点开始,可以向右最大扩展多少

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range

	//只能向右操作 相当于0:2
	slice1 = slice1[:2]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//只能向右操作 相当于将左边界向右移动
	slice1 = slice1[2:]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//只能向右操作 相当于0:2
	slice1 = slice1[:2]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//返回一个类型为 T 的初始值?
	slace3 := make([]int, 5, 10)
	fmt.Println(slace3)
	//new(T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址
	//即数组的指针 &[]int
	slace4 := new([10]int)[0:5]
	fmt.Println(slace4)

	var slice5 []int = make([]int, 4)

	slice5[0] = 1
	slice5[1] = 2
	slice5[2] = 3
	slice5[3] = 4

	for idx, value := range slice5 {
		fmt.Printf("Slice at %d is: %d\n", idx, value)
	}

	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	if _, ok := mapLit["one"]; ok {
		fmt.Printf("mapLit contains key one")
	}

	//map切片
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = i
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	fmt.Println()

	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	//靠,强行切片
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
	fmt.Println()




}
