package main

import (
    "fmt"
    "os"

    "github.com/gogo/protobuf/proto"
    "go_code/protoFile"
)

func write() {
    p1 := &protoFile.Phone{
        Id:     1,
        Number: "12345678",
    }

    data, _ := proto.Marshal(p1)

    // 把数据写入文件
    os.WriteFile("./test.txt", data, 0666)
}

func read() {
    // 读取文件数据
    data, _ := os.ReadFile("./test.txt")
    p1 := &protoFile.Phone{}
    // 解码数据
    proto.Unmarshal(data, p1)

    fmt.Printf("%+v\n", p1)
}

func show() {
    p1 := &protoFile.Phone{
        Id:     3232,
        Number: "9999999",
    }

    fmt.Println("id = ", p1.GetId())
}

func main() {
    // write()
    // read()
    // show()

    myPhone := &protoFile.Phone{
        Id:     111,
        Number: "444",
        Name:   "555",
        Email:  "666",
    }

    // 序列化
    marshal, err := proto.Marshal(myPhone)
    if err != nil {
        panic(err)
    }

    // 反序列化
    newPhone := &protoFile.Phone{}
    err = proto.Unmarshal(marshal, newPhone)
    if err != nil {
        panic(err)
    }
    fmt.Println("newPhone = ", newPhone)
}
