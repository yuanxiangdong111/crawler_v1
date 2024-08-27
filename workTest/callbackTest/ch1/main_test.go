package main

import (
    "sync"
    "testing"
)

var mu sync.Mutex

// 普通 map 读写
func BenchmarkMapWrite(b *testing.B) {
    m := make(map[int]int)
    for i := 0; i < b.N; i++ {
        m[i] = i
    }
}

func BenchmarkMapRead(b *testing.B) {
    m := make(map[int]int)
    for i := 0; i < b.N; i++ {
        m[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        mu.Lock()
        _ = m[i]
        mu.Unlock()
    }
}

// sync.Map 读写
func BenchmarkSyncMapWrite(b *testing.B) {
    var m sync.Map
    for i := 0; i < b.N; i++ {
        m.Store(i, i)
    }
}

func BenchmarkSyncMapRead(b *testing.B) {
    var m sync.Map
    for i := 0; i < b.N; i++ {
        m.Store(i, i)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = m.Load(i)
    }
}

func main() {
    // 这是为了避免 "no non-test Go files" 错误。
    // 通常，基准测试是通过 `go test` 命令运行的。
}
