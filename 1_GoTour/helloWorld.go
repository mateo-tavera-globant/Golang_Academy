package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println("Hello World. You are setting up the Go workspace")
	x, y := swap("Let's", "Begin")
	fmt.Println(x, y)
}
