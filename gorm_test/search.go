package main

import (
	"fmt"
	"go_code/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func main() {
	startName := "yuanxiangdong_400"
	endName := "yuanxiangdong_700967809"
	var user []dao.User
	err := DB.Where("username > ? and username < ?", startName, endName).Find(&user).Error
	if err != nil {
		fmt.Println("err = ", err)
	}

	fmt.Println("查找出来的长度 = ", len(user))
}
