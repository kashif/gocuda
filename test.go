package main

import (
	"fmt"
	"cuda"
)

func main() {
	res := cuda.Init(0)
	fmt.Printf("%d\n", res)
	
	driverVersion := 0
	cuda.DriverGetVersion(&driverVersion)
	fmt.Printf("%d\n", driverVersion)
}
