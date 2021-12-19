package redis_study

import (
    `fmt`
    `github.com/go-redis/redis/v8`
    `math/rand`
    `testing`
    `time`
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func lock() int {
    // 自旋锁
    for {
        randNum := rand.Intn(10000)
        ok := rdb.SetNX(ctx, "set-nx", randNum, time.Second).Val()
        if ok {
            return randNum
        }

        fmt.Println("lock wait...")

        // 适当延迟，减少cpu压力
        time.Sleep(300*time.Millisecond)
    }
}

func unlock(randNum int) {
    // Lua脚本，保证判断跟删除原子操作
    delIfEqual := redis.NewScript(`
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end
    `)

    n, err := delIfEqual.Run(ctx, rdb, []string{"set-nx"}, randNum).Result()
    if err != nil {
        panic(err)
    }

    if n.(int64) == 1 {
        fmt.Println("unlock self", randNum)
    } else {
        fmt.Println("expire, undo unlock", randNum)
    }
}

func expire() {
    randNum := lock()
    defer unlock(randNum)

    fmt.Println("expire thread", randNum)

    time.Sleep(2*time.Second)
}

func unexpire() {
    randNum := lock()
    defer unlock(randNum)

    fmt.Println("unexpire thread", randNum)

    time.Sleep(500*time.Millisecond)
}

func Test_SetNX(t *testing.T) {
    go expire()
    go unexpire()

    time.Sleep(3*time.Second)
}


