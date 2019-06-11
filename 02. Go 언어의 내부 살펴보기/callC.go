package main

import (
	"fmt"
	"unsafe"

	// #cgo CFLAGS: -I${SRCDIR}/callClib
	// #cgo LDFLAGS: ${SRCDIR}/callC.a
	// #include <stdlib.h>
	// #include <callC.h>
	"C"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Mihalis!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)
	fmt.Println("All perfectly done!")
}
