package main

import (
	"fmt"
	"cuda"
)

func main() {
	res := cuda.Init(0)
	fmt.Printf("%d\n", res)
}
