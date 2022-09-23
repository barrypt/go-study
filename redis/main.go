package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx context.Context = context.Background()

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       15,
		PoolSize: 20,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

//str key  存取
func strRedis() {
	rdb.Set(ctx, "name", "wushaoyu", 0)
	rdb.Set(ctx, "name12", "wushaoyu111", time.Second*1000)
	ff, kerr := rdb.SetNX(ctx, "12233", "22334", time.Minute).Result()
	fmt.Printf("key:%v 设置成功", kerr)

	if kerr != nil {
		fmt.Printf("key:%v get  err:%v", "12233", kerr)
	}
	fmt.Println("lock:", ff)
	ff1, kerr1 := rdb.SetNX(ctx, "12233", "2233456", time.Minute).Result()
	fmt.Printf("key:%v 设置成功", kerr1)

	if kerr1 != nil {
		fmt.Printf("key:%v get  err:%v", "12233", kerr1)
	}
	fmt.Println("lock1:", ff1)
	rdb.Set(ctx, "age", "123", 0)
	val, err := rdb.Get(ctx, "age").Result()
	if err != nil {
		fmt.Printf("key:%v get  err:%v", "age", err)
	}
	fmt.Printf("val: %v\n", val)
}

//Hash key  存取
func hashRedis() {

	rdb.HSet(ctx, "122", "1", 2, 3, 4, 5, 6, 7, 8, 9, 10).Result()
	vv, err1 := rdb.HGet(ctx, "122", "1").Result()
	if err1 != nil {
		fmt.Printf("HGet1: %v\n", err1)
	}
	fmt.Printf("vv: %v\n", vv)
	vv2, err2 := rdb.HGet(ctx, "122", "11").Result()
	if err2 != redis.Nil {
		fmt.Printf("HGet2: %v\n", err2)
	}
	fmt.Printf("vv2: %v\n", vv2)

	rdb.SAdd(ctx, "set", 12, 3, 4, 5, 6, 6, 6, 6, 5, 4, 4)
	rdb.LPush(ctx, "llist", 1, 2, 3, 4, 5, 5, 6, 4, 343, 3, 3, 2, 2)
	rdb.RPush(ctx, "rlist", 1, 2, 3, 4, 5, 5, 6, 4, 343, 3, 3, 2, 2)
	rdb.ZAdd(ctx, "zadd", &redis.Z{Score: 12, Member: 456}, &redis.Z{Score: 15, Member: 456789}, &redis.Z{Score: 100, Member: 45678349})
	zp, zerr := rdb.ZPopMax(ctx, "zadd", 2).Result()
	if zerr == redis.Nil {
		fmt.Printf("zp: %v\n", zp)
	}
	zp1, zerr1 := rdb.ZCard(ctx, "zadd").Result()
	if zerr1 == redis.Nil {
		fmt.Printf("zp: %v\n", zp)
	}
	fmt.Printf("zp1: %v\n", zp1)
	s, err := rdb.HGet(ctx, "user", "firstname").Result() //返回值,报错信息
	// s := rdb.HGet(ctx, "user", "firstname").Val() //返回值, 返回单个key的值
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("s: %v\n", s)
	fmt.Println("========================")
	m, _ := rdb.HGetAll(ctx, "user").Result() //map[firstname:wu lastname:shao]
	fmt.Printf("%+v\n", m)
	// fmt.Printf("m[\"firstname\"]: %v\n", m["firstname"])
	// fmt.Printf("m[\"lastname\"]: %v\n", m["lastname"])
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed,err:%v\n", err)
		return
	}
	fmt.Println("connect redis success....")

	//程序退出时释放相关资源
	defer rdb.Close()

	strRedis()
	hashRedis()

}
