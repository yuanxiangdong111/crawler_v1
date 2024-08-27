package main

import (
    "fmt"

    "go_code/go_gin/ch2"
)

type A struct {
    Name *ch2.SSS
}

func init() {
    fmt.Println("A init")
}

func main() {
    fmt.Println("111")
    a := &A{}
    fmt.Println(a)
}
