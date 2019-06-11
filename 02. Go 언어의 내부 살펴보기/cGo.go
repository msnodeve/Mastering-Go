package main

import "fmt"

// #include <stdio.h>
// void callC() {
// 	printf("Calling C code !\n");
// }
import "C"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}
