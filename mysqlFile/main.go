package main

import (
    "bytes"
    "fmt"

    "go_code/dao"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Student struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Sex     string  `json:"sex"`
    Address string  `json:"address"`
    Email   string  `json:"email"`
    Score   float64 `json:"score"`
    School  string  `json:"school"`
}

var DB *gorm.DB

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

func main() {

    // s := "{\"name\":\"袁湘东\",\"age\":26,\"sex\":\"男\",\"address\":\"北京市海淀区222\",\"email\":\"130xxxxxxxx@qq.com\",\"score\":88.88,\"school\":\"high school\"}"
    // s1 := dao.User{
    //     Username:   "yuanxiangdong",
    //     Password:   "{\"name\":\"袁湘东\",\"age\":26,\"sex\":\"男\",\"address\":\"北京市海淀区222\",\"email\":\"130xxxxxxxx@qq.com\",\"score\":88.88,\"school\":\"high school\"}",
    //     CreateTime: time.Now().Unix(),
    // }
    // var ss []dao.User
    // ss = append(ss, s1)
    // DB.Create(ss)
    // if err != nil {
    //     panic(err)
    // }
    // getFromMysqlByName("yuanxiangdong")

    // UserTables := make([]string, 0, 256)
    // for i := 0; i < 256; i++ {
    //     UserTables = append(UserTables, ConcatenateStrings("galaxy_user_", strconv.FormatInt(int64(i), 16)))
    // }
    //
    // fmt.Println("UserTables = ", UserTables)
}

func getFromMysqlByName(name string) {
    s1 := dao.User{}

    db, err := DB.Where("username = ?", name).Find(&s1).DB()
    if err != nil {
        panic(err)
    }
    fmt.Println("s1 = ", s1.Password)
    fmt.Println(db)
}

func ConcatenateStrings(s ...string) string {
    if len(s) == 0 {
        return ""
    }
    var buffer bytes.Buffer
    for _, i := range s {
        buffer.WriteString(i)
    }
    return buffer.String()
}
