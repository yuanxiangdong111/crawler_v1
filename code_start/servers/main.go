package main

import (
    "fmt"
)

func main() {

    i := A(22, 2)(11, 2)
    fmt.Println(i)
}

func A(a, b int) func(int, int) int {
    fmt.Println("a+b = ", a+b)
    return func(a, b int) int {
        return a + b
    }
}
