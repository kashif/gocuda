package cuda

// #include <cuda.h>
// #include <stdlib.h>
import "C"

func Init(flags C.uint) int {
	return int(C.cuInit(flags))
}
