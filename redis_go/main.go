package main

import (
    "fmt"

    "github.com/go-redis/redis/v8"
    "golang.org/x/net/context"
)

var rdb *redis.Client

func init() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",
        DB:       0,
        PoolSize: 100,
    })
}

type Student struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Sex     string  `json:"sex"`
    Address string  `json:"address"`
    Email   string  `json:"email"`
    Score   float64 `json:"score"`
    School  string  `json:"school"`
}

func main() {
    ctx := context.Background()
    // redisPipe := rdb.Pipeline()
    key := "user:12345678"
    // match := "a*"

    // values := map[string]string{
    //     "aaa":   "11",
    //     "aa123": "23131",
    //     "a12":   "23131",
    //     "aa23":  "2323",
    // }
    // s1 := Student{
    //     Name:    "袁湘东",
    //     Age:     26,
    //     Sex:     "男",
    //     Address: "北京市海淀区",
    //     Email:   "130xxxxxxxx@qq.com",
    //     Score:   88.88,
    //     School:  "high school",
    // }
    // s1 := Student{}
    s := "{\"name\":\"袁湘东\",\"age\":26,\"sex\":\"男\",\"address\":\"北京市海淀区222\",\"email\":\"130xxxxxxxx@qq.com\",\"score\":88.88,\"school\":\"high school\"}"
    s = ""
    fmt.Println("s = ", s)
    fmt.Println("len(s) = ", len(s))
    // values := "{'name':'袁湘东','age':26,'sex':'男','address':'北京市海淀区','email':'130xxxxxxxx@qq.com','score':88.88,'school':'high school'}"
    // fmt.Println("values = ", values)

    // marshal, err := json.Marshal(s1)
    // if err != nil {
    //     panic(err)
    // }
    // strMar := string(marshal)
    // fmt.Println("len(strMar) = ", len(strMar))

    if len(s) <= 0 {
        err := rdb.Del(ctx, key).Err()
        if err != nil {
            panic(err)
        }
        return
    }

    _, err := rdb.Set(ctx, key, s, 0).Result()
    if err != nil {
        panic(err)
    }

    result, err := rdb.Get(ctx, key).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("result = ", result)
    fmt.Println("len(result) = ", len(result))

}

func longestConsecutive(nums []int) int {
    numMap := make(map[int]bool)
    res := 0
    for _, v := range nums {
        numMap[v] = true
    }

    for i := 0; i < len(nums); i++ {
        if numMap[nums[i]-1] == false {
            currentNum := nums[i]
            currentLen := 1

            for numMap[currentNum+1] == true {
                currentNum++
                currentLen++
            }

            if currentLen > res {
                res = currentLen
            }
        }

    }

    return res

}
