package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmy_library
#include <my_library.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("Invoking C library...\n")
	fmt.Println("Done", C.x(10))
}
