package main

import (
    "fmt"
)

func main() {

    s := 10
    fmt.Printf("s = %b\n", s)
    s = s >> 1
    fmt.Println("s = ", s)

}
