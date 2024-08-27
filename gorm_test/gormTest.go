package main

import (
	"fmt"
	"go_code/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var DB *gorm.DB

func init() {
	dsn := "root:12345678@tcp(localhost:3306)/student?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("err = ", err)
		panic("failed to connect database")
	}
	DB = db
}

func HanlerMysql(users []dao.User, count int, goroutineIndex int) {
	// 14 -> 3
	// 5 5 4
	// 0 1 2
	defer wg.Done()
	k := len(users) / count
	remainder := len(users) % count
	result := make([][]dao.User, 0, count)
	for i := 0; i < count; i++ {
		start := i * k
		end := start + k

		if i < remainder {
			start += i
			end += i + 1
		} else {
			start += remainder
			end += remainder
		}
		result = append(result, users[start:end])
	}
	batchCreate(result[goroutineIndex-1], 50)
}

func batchCreate(users []dao.User, nums int) {
	//preNum := 50
	for i := 0; i < len(users); i += nums {
		end := i + nums

		if end > len(users) {
			end = len(users)
		}

		preUsers := users[i:end]
		DB.Create(preUsers)
	}
}

func main() {

	nowTime := time.Now()
	var users []dao.User

	//查询所有的user
	DB.Find(&users)

	copyUser := users

	fmt.Println("len(users) = ", len(users))

	// 生成随机的名字和密码
	for i := 0; i < 960000; i++ {
		rand.Seed(time.Now().UnixNano())
		randNums := rand.Intn(1000000000)

		passWord := "123456_" + strconv.Itoa(randNums)
		userName := "yuanxiangdong_" + strconv.Itoa(randNums)
		createTime := time.Now().UnixMilli()

		user := dao.User{
			Username:   userName,
			Password:   passWord,
			CreateTime: createTime,
		}
		users = append(users, user)

	}

	usersMap := make(map[string]dao.User)

	for _, v := range users {
		usersMap[v.Username] = v
	}

	var newUsers []dao.User

	for _, v := range copyUser {
		if _, ok := usersMap[v.Username]; ok {
			delete(usersMap, v.Username)
		}
	}

	for _, v := range usersMap {
		newUsers = append(newUsers, v)
	}
	fmt.Println("耗时1：", time.Since(nowTime))

	wg = sync.WaitGroup{}

	wg.Add(10)

	for i := 1; i <= 10; i++ {
		//defer wg.Done()
		go HanlerMysql(newUsers, 10, i)
	}

	// mysql分批次创建
	//preNum := 50
	//for i := 0; i < len(newUsers); i += preNum {
	//	end := i + preNum
	//
	//	if end > len(newUsers) {
	//		end = len(newUsers)
	//	}
	//
	//	preUsers := newUsers[i:end]
	//	DB.Create(preUsers)
	//}

	fmt.Println("len(user) = ", len(users))
	fmt.Println("len(newUsers) = ", len(newUsers))
	fmt.Println("总耗时：", time.Since(nowTime))

	//user := &User{
	//  Username:   "yuanxiangdong",
	//  Password:   "123456",
	//  CreateTime: time.Now().UnixMilli(),
	//}
	//db.Create(users)
	wg.Wait()
}
