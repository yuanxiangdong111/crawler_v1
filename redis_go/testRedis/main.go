package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/go-redis/redis/v8"
    "golang.org/x/net/context"
)

var rdb *redis.Client

func init() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "172.20.194.240:6379",
        Password: "T2eCGHY#S6Yrd26",
        DB:       0,
        PoolSize: 100,
    })
}

type Student struct {
    ScrollToPos_0 string `json:"scrollToPos_0"`
    Textbook_list string `json:"textbook_list"`
}

func main() {
    ctx := context.Background()
    // redisPipe := rdb.Pipeline()
    key := "galaxy.phoenix.attr:1:9992353769580"
    // keys := rdb.Keys(ctx, "galaxy.phoenix.attr*").Val()
    match := "a*"

    fields, _ := rdb.HScan(ctx, key, 0, match, 500).Val()
    fmt.Println("fields = ", fields)
}

func WriteFile(ctx context.Context, cmd []redis.Cmder) error {
    f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        panic(err)
        return err
    }
    defer f.Close()
    writer := bufio.NewWriter(f)
    for _, v := range cmd {
        tmp := v.String()
        tmp1 := []byte(tmp)
        // fmt.Println("tmp1 = ", tmp1)
        n, err := writer.Write(tmp1)
        // n, err := writer.WriteString(v.String())
        if err != nil {
            fmt.Println("err = ", err)
            return err
        }
        fmt.Printf("成功写入 %d 字节数据\n", n)
    }

    writer.Flush()
    return nil
}
