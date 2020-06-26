package tt_practice

import (
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/pkg/tpkg"
	"os"
	"runtime"
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

	t1 ,t2:= 1,"here t2"
	t1Next := t1

	fmt.Println(&t1)
	fmt.Println(&t1Next == &t1) //值类型
	fmt.Println(&t2)



}

