package main

import (
    "fmt"
)

var bbb = B{}

func main() {

    bbb.GetA("adsadasdsads")
}

type B struct{}

func (b B) GetA(s string) {
    fmt.Println(s)
}

type A interface {
    GetA(s string)
}
