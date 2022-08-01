/*
   - Taking inputs
*/

package main

import "fmt"

func main() {
	var Name string
	var Age int
	fmt.Printf("What is your name? : ")
	fmt.Scanln(&Name)
	fmt.Printf("How old are you? : ")
	fmt.Scanln(&Age)
	fmt.Println("Hello, ", Name, "nice to see you")
	fmt.Println("Oh wow you were born on", 2022-Age)

}
