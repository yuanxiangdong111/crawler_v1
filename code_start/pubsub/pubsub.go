package pubsub

import (
    "sync"
    "time"
)

type (
    subscriber chan interface{}         // 订阅者为一个通道
    topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// Publisher 发布者
type Publisher struct {
    mu          sync.RWMutex             // 读写锁
    buffer      int                      // 订阅者的缓存
    timeout     time.Duration            // 发布超时时间
    subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建一个发布者对象，可以设置发布超时时间和缓存队列长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
    return &Publisher{
        buffer:      buffer,
        timeout:     publishTimeout,
        subscribers: make(map[subscriber]topicFunc),
    }
}

// 订阅所有的主题的订阅者
func (p *Publisher) SubscribeAll() chan interface{} {
    return p.SubscribeTopic(nil)
}

// 添加新的订阅者
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
    ch := make(chan interface{}, p.buffer)
    p.mu.Lock()
    defer p.mu.Unlock()
    p.subscribers[ch] = topic
    return ch
}

// 取消订阅
func (p *Publisher) Evict(sub chan interface{}) {
    p.mu.Lock()
    defer p.mu.Unlock()
    delete(p.subscribers, sub)
    close(sub)
}

// 发布主题
func (p *Publisher) Publish(v interface{}) {
    p.mu.RLock()
    defer p.mu.RUnlock()
    var wg sync.WaitGroup
    for sub, topic := range p.subscribers {
        wg.Add(1)
        go p.SendTopic(sub, topic, v, &wg)
    }
    wg.Wait()
}

// 关闭所有订阅者通道
func (p *Publisher) Close() {
    p.mu.Lock()
    defer p.mu.Unlock()

    for k := range p.subscribers {
        delete(p.subscribers, k)
        close(k)
    }
}

// 发送主题
func (p *Publisher) SendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
    defer wg.Done()

    if topic != nil && !topic(v) {
        return // 过滤掉不感兴趣的消息
    }

    select {
    case sub <- v:
    case <-time.After(p.timeout): // 超时处理
    }
}
