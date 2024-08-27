package main

import (
    "net"
    "strings"
)

// 获取本机出口IP
// 和google的DNS服务器建立连接，获取本机出口IP
func GetOutBoundIP() (ip string, err error) {
    dial, err := net.Dial("udp", "8.8.8.8:53")
    if err != nil {
        return
    }
    defer dial.Close()
    addr := dial.LocalAddr().(*net.UDPAddr)
    return strings.Split(addr.String(), ":")[0], nil
}
