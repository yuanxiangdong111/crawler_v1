package main

import (
    "io"
    "net/http"
)

type Memo struct {
    requests chan request
}

type request struct {
    key      string
    response chan<- result
}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
    memo := &Memo{requests: make(chan request)}
    go memo.Serve(f)
    return memo
}

type entry struct {
    res   result
    ready chan struct{}
}

type result struct {
    val interface{}
    err error
}

func (memo *Memo) Serve(f Func) {
    cache := make(map[string]*entry)

    for req := range memo.requests {
        e := cache[req.key]
        if e == nil {
            // 为空
            // 且第一个请求
            e = &entry{ready: make(chan struct{})}
            cache[req.key] = e
            go e.call(f, req.key)
        }
        go e.deliver(req.response)

    }
}

func (e *entry) call(f Func, key string) {
    e.res.val, e.res.err = f(key)
    close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
    // 等待准备好的信号
    <-e.ready
    response <- e.res
}

func main() {
    var m Memo
    m.requests <- request{
        key:      "baidu.com",
        response: nil,
    }

}

func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}
