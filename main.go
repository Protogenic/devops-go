package main

import "fmt"

func main() {
    sum := add(5, 3)
    fmt.Println("Сумма:", sum)
  
    difference := subtract(5, 3)
    fmt.Println("Разность:", difference)
}

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}
