package redisStore

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var RedisDb *redis.Client

// 创建 redis 链接
func init() {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		//连接失败
		println(err)
	}
}

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// 实现设置captcha的方法
func (r *RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	//time.Minute*2：有效时间2分钟
	err := RedisDb.Set(ctx, key, value, time.Minute*2).Err()

	return err
}

// 实现获取captcha的方法
func (r *RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		//clear为true，不管验证成功或失败都会删除此验证码
		err := RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// 实现验证captcha的方法
func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
