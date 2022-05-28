package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd_recursive(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd_recursive(b, a%b)

}
func main() {
	first, err := strconv.Atoi(os.Args[1]) // if it is a number it parse it to base 10 otherwise it returns an error

	if err != nil {
		fmt.Println("First input parameter must be integer")
		os.Exit(1)
	}

	second, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Second input parameter must be integer")
		os.Exit(1)
	}
	fmt.Println("The Greatest common deviser of ", first, " and ", second, " is :", gcd_recursive(first, second))
}
