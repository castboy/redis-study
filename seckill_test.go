// 秒杀，保证不超卖和一人一单

package redis_study

import (
	`bytes`
	`fmt`
	"github.com/go-redis/redis/v8"
	`runtime`
	`strconv`
	`testing`
	`time`
)



var incrby *redis.Script
var reentrant *redis.Script

func init() {
	if err := rdb.HMSet(ctx, "voucher", "id", 0, "amount", 10, "name", "voucher").Err(); err != nil {
		panic(err)
	}

	// Lua脚本，保证不超卖
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

func Test_SecKill(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			lockName := "seckillMutex:" + "activityID:" + "userID"

			// 活动持续时间1000s
			expire := "1000"

			// 保证一人一单
			if reentrantLock(lockName, expire) {
				decr_right()
			}
		}()
	}

	time.Sleep(10*time.Second)
}

func Test_Reentrant(t *testing.T) {
	go func() {
		lockName := "Test_Reentrant"
		expire := "1000"

		reentrantLock(lockName, expire)
		reentrantLock(lockName, expire)
		reentrantLock(lockName, expire)

		time.Sleep(10*time.Second)

		reentrantUnlock(lockName)
		reentrantUnlock(lockName)
		reentrantUnlock(lockName)
	}()

	time.Sleep(20*time.Second)
}

func decr_right() {
	n, err := incrby.Run(ctx, rdb, []string{"voucher", "amount"}).Result()
	if n.(int64) > 0 {
		fmt.Println(n, err)
	}
}

// 可重入锁上锁，保证一人一单
func reentrantLock(lockName, expire string) bool {
	goID := getGoroutineID()

	reentrant := redis.NewScript(`
		local key = KEYS[1]
		local threadID = ARGV[1]
		local releaseTime = ARGV[2] -- 第三个参数，锁的自动释放时间
	
		if (redis.call('exists', key) == 0) then -- 判断锁是否已经存在
			redis.call('hset', key, threadID, '1') -- 不存在则获取锁
			redis.call('expire', key, releaseTime) -- 设置有效期
			return 1 -- 返回结果
		end
	
		if (redis.call('hexists', key, threadID) == 1) then -- 锁已经存在，判断是否是当前thread
			redis.call('hincrby', key, threadID, '1') -- 如果是自己，则重入次数+1
			redis.call('expire', key, releaseTime) -- 设置有效期
			return 1 -- 返回结果
		end

		return 0 -- 代码走到这里，说明获取锁的不是自己，获取锁失败
	`)

	n, err := reentrant.Run(ctx, rdb, []string{lockName}, strconv.FormatInt(int64(goID), 10), expire).Result()
	if err != nil {
		panic(err)
	}

	return n.(int64) == 1
}

// 可重入锁解锁
func reentrantUnlock(lockName string) {
	goID := getGoroutineID()

	reentrant := redis.NewScript(`
		local key = KEYS[1]
		local threadID = ARGV[1]

		if (redis.call('hexists', key, threadID) == 0) then -- 判断当前锁是否是被自己持存
			return -1 -- 如果不是自己，则直接返回
		end

		local count = redis.call('hincrby', key, threadID, -1) -- 是自己的锁，则重入次数-1

		if (count == 0) then -- 判断重入次数是否为0
			redis.call('del', key) -- 等于0说明可以释放锁，直接删除
			return 0
		end

		return count
	`)

	n, err := reentrant.Run(ctx, rdb, []string{lockName}, strconv.FormatInt(int64(goID), 10)).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}

func getGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}