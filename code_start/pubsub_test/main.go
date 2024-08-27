package main

import (
    "fmt"
    "strings"
    "time"

    "go_code/code_start/pubsub"
)

func main() {
    p := pubsub.NewPublisher(100*time.Millisecond, 10)
    defer p.Close()
    // 所有的通知
    all := p.SubscribeAll()

    // 订阅包含golang的主题
    // 如果包含的话 返回true
    golang := p.SubscribeTopic(func(v interface{}) bool {
        if v, ok := v.(string); ok {
            return strings.Contains(v, "golang")
        }
        return false
    })

    p.Publish("hello world")
    p.Publish("hello golang")

    go func() {

        for k := range all {
            fmt.Println("all : ", k)
        }
    }()

    go func() {

        for k := range golang {
            fmt.Println("golang : ", k)
        }
    }()

    time.Sleep(3 * time.Second)

}
