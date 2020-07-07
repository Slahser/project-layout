package main

import (
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/app/tt_practice"
	"time"
)

func main() {
	//=====
	start := time.Now()

	//practice
	tt_practice.Tt1()
	tt_practice.Tt2()
	tt_practice.Tt3()

	//=====
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("total cost is %s\n", cost)
}
