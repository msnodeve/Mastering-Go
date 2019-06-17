package aPackage

import "fmt"

func A() {
	fmt.Println("This is function A!")
}

func B() {
	fmt.Println("PrivateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21
