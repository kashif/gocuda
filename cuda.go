package cuda

// #include <cuda.h>
// #include <stdlib.h>
import "C"

/*********************************
** Initialization
*********************************/

func Init(flags C.uint) int {
	return int(C.cuInit(flags))
}

/*********************************
** Driver Version Query
*********************************/

func DriverGetVersion(driverVersion *int) int {
  var i C.int
  ret := int(C.cuDriverGetVersion(&i))
  *driverVersion = int(i)
  return ret
}
