package utils

import "fmt"

func Add(a int, b int) int {
    return a + b
}

func Subtract(a int, b int) int {
    return a - b
}

func Multiply(a int, b int) int {
    return a * b
}

func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func IsEven(n int) bool {
    return n%2 == 0
}

func IsOdd(n int) bool {
    return n%2 != 0
}

func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}