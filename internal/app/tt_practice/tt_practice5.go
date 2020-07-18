package tt_practice

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

//TODO variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

type T struct {
	A int
	B string
}

/**
reflect & io
*/
func Tt5() {
	Trace("tt_practice5#Tt5()")
	defer Untrace("tt_practice5#Tt5()")

	//reflect basic type
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())

	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)

	//reflect struct
	t := T{23, "skidoo"}
	//reflect.ValueOf(&t).Elem()
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)

	inputFile, inputError := os.Open("internal/README.md")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	fmt.Println(inputFile)
	//t_bufio(inputFile)
	//cat(inputFile)

}

func t_bufio(f *os.File) {
	defer f.Close()
	inputReader := bufio.NewReader(f)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

/**
read file by slice
*/
func cat(f *os.File) {
	defer f.Close()
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}
