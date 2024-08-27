package main

import (
    "io"
    "net/http"
    "sync"
)

type Memo struct {
    mu    sync.Mutex
    f     Func
    cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type entry struct {
    res   result
    ready chan struct{}
}

type result struct {
    val interface{}
    err error
}

func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

func (memo *Memo) Get(key string) (interface{}, error) {
    memo.mu.Lock()
    e := memo.cache[key]
    if e == nil {
        // 为空的话说明缓存中没数据
        // 需要去设置
        // 设置的时候其他的 协程必须阻塞 直到e 不等于nil再解锁
        e = &entry{ready: make(chan struct{})}
        memo.cache[key] = e
        memo.mu.Unlock()

        e.res.val, e.res.err = memo.f(key)
        close(e.ready) // 广播通知 数据已经准备好了

    } else {
        // 非空
        memo.mu.Unlock()
        // 可能有多个协程在等待读
        <-e.ready
    }
    return e.res.val, e.res.err
}
