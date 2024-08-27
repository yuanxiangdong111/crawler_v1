package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func Producer(factor int, out chan<- int) {
    i := 1
    for {
        time.Sleep(200 * time.Millisecond)
        out <- i * factor
        i++
    }
}

func Cosumer(in <-chan int) {
    for v := range in {
        time.Sleep(200 * time.Millisecond)
        fmt.Println(v)
    }
}

func main() {

    ch := make(chan int, 64)
    go Producer(3, ch)
    go Producer(5, ch)
    go Cosumer(ch)

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    fmt.Printf("quit (%v)\n", <-sig)
}
