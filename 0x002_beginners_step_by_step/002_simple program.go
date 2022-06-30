package main
import "fmt"

/*
	- A simple program to add two numbers
	- Learn how to declare variables
	- Learn how to take an input
*/

func main() {
    var firstNumber int=0
    var secondNumber int=0
    var result int=0
    
    fmt.Printf("Enter first number: ")
    fmt.Scanf("%d",&firstNumber)
    
    fmt.Printf("Enter number2: ")
    fmt.Scanf("%d",&secondNumber)
    
    result = firstNumber & secondNumber
    fmt.Printf("Result: %d",result)
}
