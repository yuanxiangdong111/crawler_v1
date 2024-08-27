package main

import (
    "fmt"
    "io"
    "net/http"
)

type Func func(key string) (interface{}, error)

type result struct {
    value interface{}
    err   error
}

type Memo struct {
    f     Func              // 缓存的函数
    cache map[string]result // 缓存的内容
}

func New(f Func) *Memo {
    return &Memo{
        f:     f,
        cache: make(map[string]result),
    }
}

func (memo *Memo) Get(key string) (interface{}, error) {
    res, ok := memo.cache[key]
    if !ok {
        res.value, res.err = memo.f(key)
        memo.cache[key] = res
    }

    return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

var status chan int

func main() {

    status = make(chan int)
    a := make(chan int)

    go inFunc(a)
    go outFunc(a)

    <-status

}

func inFunc(inChan chan<- int) {
    for i := 0; i < 10; i++ {
        inChan <- i
    }

    close(inChan)
}

func outFunc(outChan <-chan int) {
    for k := range outChan {
        fmt.Println(k)
    }
    status <- 1
}
