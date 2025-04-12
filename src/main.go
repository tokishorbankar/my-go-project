package main

import (
	"fmt"
	"my-go-project/pkg/utils"
)

func calculate(firstArg, secondArg int) {
	fmt.Println("Performing calculations...")
	fmt.Printf("First Arg: %d\n", firstArg)
	fmt.Printf("Second Arg: %d\n", secondArg)

	fmt.Printf("Sum: %d\n", utils.Add(firstArg, secondArg))
	fmt.Printf("Difference: %d\n", utils.Subtract(firstArg, secondArg))
	fmt.Printf("Multi: %d\n", utils.Multiply(firstArg, secondArg))
	quotient, err := utils.Divide(firstArg, secondArg)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %d\n", quotient)
	}
	fmt.Printf("First Arg Is Even: %t\n", utils.IsEven(firstArg))
	fmt.Printf("First Arg Is Odd: %t\n", utils.IsOdd(firstArg))
	fmt.Printf("Second Arg Is Even: %t\n", utils.IsEven(secondArg))
	fmt.Printf("Second Arg Is Odd: %t\n", utils.IsOdd(secondArg))
	fmt.Printf("Factorial: %d\n", utils.Factorial(firstArg))
	fmt.Println("Calculations completed successfully.")
}

func main() {
	fmt.Println("Welcome to My Go Project!")
	calculate(5, 2)
	fmt.Println("Goodbye!")

}
