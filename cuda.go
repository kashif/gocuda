package cuda

// #include <cuda.h>
// #include <stdlib.h>
import "C"

/*********************************
** Initialization
*********************************/

func Init(flags uint32) int {
	return int(C.cuInit(C.uint(flags)))
}

/*********************************
** Driver Version Query
*********************************/

func DriverGetVersion() (driverVersion, err  int) {
  var cDriverVersion, cErr C.int
  cErr = C.int(C.cuDriverGetVersion(&cDriverVersion))
  driverVersion = int(cDriverVersion)
  err = int(cErr)
  return
}

/************************************
**    Device management
***********************************/

func DeviceGet(ordinal int) (device, err int) {
	var cDevice C.CUdevice
	var cErr C.int
	cErr = C.int(C.cuDeviceGet(&cDevice, C.int(ordinal)))
	device = int(cDevice)
	err = int(cErr)
	return
}

func DeviceGetCount() (count, err int) {
	var cCount, cErr C.int
	cErr = C.int(C.cuDeviceGetCount(&cCount))
	count = int(cCount)
	err = int(cErr)
	return
}

func DeviceComputeCapability(dev int) (major, minor, err int) {
	var cErr, cMajor, cMinor C.int
	cErr = C.int(C.cuDeviceComputeCapability(&cMajor, &cMinor, C.CUdevice(dev)))
	major = int(cMajor)
	minor = int(cMinor)
	err = int(cErr)
	return
}