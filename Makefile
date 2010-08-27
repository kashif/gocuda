include $(GOROOT)/src/Make.inc

TARG=cuda

CGOFILES=\
	cuda.go

CGO_LDFLAGS=-lcuda
CGO_CFLAGS=-I/usr/local/cuda/include

CLEANFILES+=test

include $(GOROOT)/src/Make.pkg

# Simple test programs

test: install test.go
	$(GC) test.go
	$(LD) -o $@ test.$O
