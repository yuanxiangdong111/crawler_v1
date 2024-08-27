package main

import (
    "fmt"
    "sync"
    "time"
)

type Student struct {
    Name    string
    Age     int
    Address string
}

func init() {

}

func NewStudent(name string, age int, address string) *Student {
    return &Student{
        Name:    name,
        Age:     age,
        Address: address,
    }
}

func (s *Student) SetName(name string) {
    s.Name = name
}

func (s *Student) SetAge(age int) {
    s.Age = age
}

func worker(wg *sync.WaitGroup, cancel chan int) {
    defer wg.Done()
    for {
        select {
        case <-time.After(200 * time.Millisecond):
            fmt.Println("超时退出")
            return
        case <-cancel:
            fmt.Println("正常退出")
            return
        default:
            fmt.Println("hello")
        }
    }
}

func main() {

    ch := make(chan int)
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(&wg, ch)
    }
    time.Sleep(200 * time.Microsecond)
    close(ch)
    wg.Wait()
}
