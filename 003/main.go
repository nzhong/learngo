package main

import (
	"fmt"
	"runtime"
	"strconv"
)

var test bool

func main() {
	if 3 > 2 {
		fmt.Println(runtime.NumCPU() + 10)
	}
	fmt.Println("main" + strconv.FormatBool(test))
	test2a()
}
