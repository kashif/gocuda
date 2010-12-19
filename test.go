package main

import (
	"fmt"
	"cuda"
)

func main() {
	err := cuda.Init(0)
	fmt.Printf("Initialized: %d\n", err)
	
	driverVersion, err := cuda.DriverGetVersion()
	if (err == 0) {
		fmt.Printf("Driver Version: %d\n", driverVersion)
	}
	
	device, err := cuda.DeviceGet(0)
	if (err == 0) {
		fmt.Printf("Device: %d\n", device)
	}
	
	count, err := cuda.DeviceGetCount()
	if (err == 0) {
		fmt.Printf("Device count: %d\n", count)
	}
}
