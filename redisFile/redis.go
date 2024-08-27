package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

var (
	redisPre   = "xd.yuan.redis"
	redisServe = "test"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	timeNow := time.Now()
	ctx := context.Background()
	pd := rdb.Pipeline()
	//xd.yuan.redis:test:yuan0:dong0

	for i := 0; i < 300000; i++ {
		data := make(map[string]string, 1)
		uidStr := fmt.Sprintf("yuan%d", i)
		oidStr := fmt.Sprintf("dong%d", i)
		data[uidStr] = oidStr
		//rdb.HMSet(ctx, GenRedisKey(uidStr, oidStr), data)
		//rdb.Del(ctx, GenRedisKey(uidStr, oidStr))
		pd.HMSet(ctx, GenRedisKey(uidStr, oidStr), data)
		//pd.HDel(ctx, GenRedisKey(uidStr, oidStr), uidStr)
	}

	_, err := pd.Exec(ctx)
	if err != nil {
		fmt.Println("err = ", err)
		panic(err)
	}
	fmt.Println("cost : ", time.Since(timeNow))
}

func GenRedisKey(uid, oid string) string {
	return fmt.Sprintf("%s:%s:%s:%s", redisPre, redisServe, uid, oid)
}
