package test01

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"testing"
)

var RDB *redis.Client

func TestRedis(t *testing.T) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	err := RDB.Set(ctx, "user", "yuanxiangdong", 0).Err()
	if err != nil {
		fmt.Println("err = ", err)
		panic(err)
	}
	result, err := RDB.Get(ctx, "user").Result()
	if err != nil {
		fmt.Println("err = ", err)
		panic(err)
	}
	fmt.Println("result = ", result)
}
