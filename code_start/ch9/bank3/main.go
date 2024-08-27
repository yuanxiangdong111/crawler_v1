package main

import (
    "sync"
)

var (
    mu      sync.Mutex
    balance int
)

func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    deposit(-amount)
    if balance < 0 {
        deposit(amount)
        // 余额不足
        return false
    }
    return true
}

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    deposit(amount)
}

func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}

func deposit(amount int) {
    balance += amount
}
