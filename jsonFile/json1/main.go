package main

import (
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
    "go_code/dao"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type TestData struct {
    ResData []JsonData `json:"res_data"`
}

type JsonData struct {
    Data string `json:"data"`
}

var DB *gorm.DB
var rdb *redis.Client

func init() {
    dsn := "root:12345678@tcp(localhost:3306)/student?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        // Logger: newLogger,
        // Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        fmt.Println("err = ", err)
        panic("failed to connect database")
    }
    DB = db
}

func init() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",
        DB:       0,
        // PoolSize: 5,
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

type ResParams struct {
    Data Student `json:"field"`
}

type Address struct {
    City  string `json:"city"`
    State string `json:"state"`
}

type Person struct {
    Name    string    `json:"name"`
    Age     int       `json:"age"`
    Address []Address `json:"address"`
}

func main() {

    // jsonStr := "{\"name\":\"袁湘东\",\"age\":26,\"sex\":\"男\",\"address\":\"北京市海淀区111\",\"email\":\"130xxxxxxxx@qq.com\",\"score\":88.88,\"school\":\"High school\"}"
    // ctx := context.Background()
    // r := ResParams{
    //     Student{
    //         Name:    "袁湘东",
    //         Age:     26,
    //         Sex:     "男",
    //         Address: "北京市海淀区111",
    //         Email:   "130xxxxxxxx@qq.com",
    //         Score:   88.88,
    //         School:  "High school",
    //     },
    // }
    // bytes, _ := json.Marshal(r.Data)
    // jsonStr := string(bytes)

    // redis 存储设置过期时间
    // galaxy.phoenix.board_person:oid:uid
    // key := "galaxy.phoenix.board_person:1:123456"
    // err2 := rdb.SetEX(ctx, key, jsonStr, 5000*time.Second).Err()
    // if err2 != nil {
    //     panic(err2)
    // }

    // 获取redis中的数据
    // result, err2 := rdb.Get(ctx, key).Result()
    // if err2 != nil {
    //     panic(err2)
    // }
    // fmt.Println("result = ", result)

    // mysql创建数据
    // err := DB.Model(&dao.GalaxyBoardPerson{}).Create(&dao.GalaxyBoardPerson{
    //     // ID:        0,
    //     Oid:       1,
    //     Uid:       "1234",
    //     Name:      "yuanxiangdong",
    //     Data:      jsonStr,
    //     CreatedAt: time.Time{},
    //     UpdatedAt: time.Time{},
    // }).Error
    // if err != nil {
    //     panic(err)
    // }

    // mysql 查找数据
    // Rest := dao.GalaxyBoardPerson{}
    uid := "1234"
    oid := 1
    // DB.Where("uid = ? and oid = ?", uid, oid).Find(&Rest)
    // fmt.Printf("Rest = %+v", Rest)

    // 将数据序列化
    // marshal, err := json.Marshal(Rest)
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println("marshal = ", string(marshal))
    // fmt.Printf("%+v", Resp.Data)

    // 写入redis中
    // key := "galaxy.phoenix.board_person:1:123456"
    // err2 := rdb.SetEX(ctx, key, string(marshal), 5000*time.Second).Err()
    // if err2 != nil {
    //     panic(err2)
    // }

    // mysql更新数据
    err := DB.Where("uid = ? and oid = ?", uid, oid).Updates(&dao.GalaxyBoardPerson{
        Oid:       1,
        Uid:       "1234",
        Name:      "yuanxiangdong",
        Data:      "dadasdadsadsads",
        CreatedAt: time.Time{},
        UpdatedAt: time.Time{},
    }).Error
    if err != nil {
        panic(err)
    }

}
