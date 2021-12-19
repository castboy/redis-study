package redis_study

import (
	"context"
	`fmt`
	"github.com/go-redis/redis/v8"
	`strconv`
	`testing`
	`time`
)

var ctx = context.Background()
var rdb *redis.Client
var incrby *redis.Script

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.HMSet(ctx, "voucher", "id", 0, "amount", 10, "name", "voucher").Err(); err != nil {
		panic(err)
	}

	// Lua脚本
	incrby = redis.NewScript(`
		if (redis.call("HEXISTS", KEYS[1], KEYS[2]) == 1) then
			local stock = tonumber(redis.call("HGET", KEYS[1], KEYS[2]))
			if (stock > 0) then
				redis.call("HINCRBY", KEYS[1], KEYS[2], -1)
				return stock
			else
				return 0
			end
		else
			return 0
		end
    `)
}

func Test_SecKill_Wrong(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			current, _ := strconv.Atoi(rdb.HGet(ctx, "voucher", "amount").Val())
			if current > 0 {
				rdb.HIncrBy(ctx, "voucher", "amount", -1)
			}
		}()
	}

	time.Sleep(10*time.Second)
}

func Test_SecKill_Right(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			n, err := incrby.Run(ctx, rdb, []string{"voucher", "amount"}).Result()
			if n.(int64) > 0 {
				fmt.Println(n, err)
			}
		}()
	}

	time.Sleep(10*time.Second)
}
