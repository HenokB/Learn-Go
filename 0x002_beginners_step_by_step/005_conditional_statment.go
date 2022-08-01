package main

import (
	"fmt"
)

func main() {
	var number int = 5
	if number%2 == 0 {
		fmt.Println("This number is even")
	} else {
		fmt.Println("This number is odd")
	}
}
