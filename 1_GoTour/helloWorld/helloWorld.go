package main

/**
1. Define the package for your program file.
 Go program execution starts with the main package.
 if your program is in main package then that is executable one.
 other packages can be treated as libraries that can be used in executable program.

2. Import helps developer to import other built-in or custom packages into your source code.
  you can import one by one or group import as shown in above.

3. Main function - Go program start execution from main function.
 func is the keyword which used to define the function.

4. We are printing a message on console using the Println() method
  of fmt package which is imported at line 2.
  Println() function exported to the outside package by declaring the with capital letter.
  function those start with small letters will not be exposed outside the package,
  those will be private.


**/
import (
	"fmt"
)

//Creating a Function
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println("Hello World. You are setting up the Go workspace")
	x, y := swap("Let's", "begin")
	fmt.Println(x, y)
}
