package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64
	if n != 0 {
		os.Exit(2)
	}
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	var sum float64
	cnt := 0
	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += n
			cnt++
		}
	}
	fmt.Println("Avg : ", sum/float64(cnt))
}
