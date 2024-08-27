package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    go broadcaster()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }
}

// an outgoing message channel
type client struct {
    ch   chan<- string
    name string
}

var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string) // all incoming client messages
)

func broadcaster() {
    clients := make(map[client]bool) // all connected clients
    for {
        select {
        case msg := <-messages:
            // 用 select 如果2秒内没有信息发送过来的话
            // 每个协程单独去处理每个客户端连接的chan

            for cli := range clients {
                select {
                case cli.ch <- msg:
                case <-time.After(2 * time.Second):
                    go func(cli client) {
                        cli.ch <- msg
                    }(cli)
                    // default:
                    //     go func(cli client) {
                    //         cli.ch <- msg
                    //     }(cli)
                }

            }
        case cli := <-entering:
            clients[cli] = true
            // 告诉新进来的用户，当前有多少人在线
            var cliNames []string
            for c := range clients {
                cliNames = append(cliNames, c.name)
            }
            cli.ch <- fmt.Sprintf("%d arrival: %v\n", len(cliNames), cliNames)

        case cli := <-leaving:
            delete(clients, cli)
            close(cli.ch)
        }
    }
}
func handleConn(conn net.Conn) {
    ch := make(chan string) // outgoing client messages
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages <- who + " has arrived"
    entering <- client{
        ch:   ch,
        name: who,
    }
    expiredTime := time.NewTicker(5 * time.Minute)
    go func() {
        <-expiredTime.C
        conn.Close()
    }()

    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ": " + input.Text()
        expiredTime.Reset(5 * time.Minute)
        // messages <- input.Text()
    }
    // NOTE: ignoring potential errors from input.Err()

    leaving <- client{
        ch:   ch,
        name: who,
    }
    messages <- who + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
    }
}
