package main

import (
	"fmt"
	"my-go-project/pkg/utils"
	"os"
	"strconv"
)


func calculatedByUser(firstArg, secondArg int) {
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

	if len(os.Args) < 3 {
		fmt.Println("Usage: ./my-go-app <firstArg> <secondArg>")
		return
	} 

	// Parse arguments
	firstArg, err1 := strconv.Atoi(os.Args[1])
	secondArg, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Both arguments must be integers.")
		return
	}

    // Perform calculations
    calculatedByUser(firstArg, secondArg)
    // Print the results
    
}
