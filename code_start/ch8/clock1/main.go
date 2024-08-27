package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    listen, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }

    for {
        accept, err := listen.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConn(accept)
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil {
            return
        }
        time.Sleep(1 * time.Second)
    }
}
