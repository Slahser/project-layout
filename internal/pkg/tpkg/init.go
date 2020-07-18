package tpkg

import (
	"flag"
	"math"
	"runtime"
)

var (
	Pi float64
)

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi

	var numCores = flag.Int("n", 2, "number of CPU cores to use")

	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
}
