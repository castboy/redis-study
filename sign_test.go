package redis_study

import (
    `fmt`
    `github.com/go-redis/redis/v8`
    `os`
    `testing`
    `time`
)

func init() {

}

// 统计连续签到次数
func Test_Sign(t *testing.T) {
    // today := time.Now().Day()
    // 假设今天是20211221
    today := 17
    offset := int64(today - 1)

    signKey := "user:sign:userID:202112"

    // 查看是否已经签到
    n, err := rdb.GetBit(ctx, signKey, offset).Result()
    if err != nil {
        panic(err)
    }

    if n == 1 {
        fmt.Println("今日已经签到，无需再签到")
        return
    }

    // 签到
    err = rdb.SetBit(ctx, signKey, offset, 1).Err()
    if err != nil {
        panic(err)
    }

    fmt.Println("签到成功")

    // 统计连续签到的次数
    res := rdb.BitField(ctx, signKey, "get", "u22", 0).Val()
    count := 0

    // 21是今日
    for i := 21; i > 0; i-- {
        if res[0] == (res[0] >> 1) << 1 {
            if i != 21 { // 允许今日未签到
                break
            }
        } else {
            count++
        }

        res[0] >>=  1
    }

    fmt.Println("连续签到次数：", count)
}

// 统计某月的签到次数
func Test_SignCount(t *testing.T) {
    key := "user:sign:userID:" + time.Now().Format("200601")

    // 0, 1, 2, 3 四个字符中1的位数
    fmt.Println(rdb.BitCount(ctx, key, &redis.BitCount{Start: 0, End: 3}).Val())
}

// 获取用户签到情况
 func Test_SignInfo(t *testing.T) {
    days := count(2021, 12)
    font := make([]int, days)
    arg := fmt.Sprintf("u%d", days)

    res := rdb.BitField(ctx, "user:sign:userID:202112", "get", arg, 0).Val()
    fmt.Println(res[0])

    for i := days - 1; i >= 0; i-- {
        if (res[0] >> 1) << 1 == res[0] {
            font[i] = 0
        } else {
            font[i] = 1
        }

        res[0] >>= 1
    }

    fmt.Println(font)
 }

func count(year int, month int) (days int) {
    if month != 2 {
        if month == 4 || month == 6 || month == 9 || month == 11 {
            days = 30

        } else {
            days = 31
            fmt.Fprintln(os.Stdout, "The month has 31 days");
        }
    } else {
        if (((year % 4) == 0 && (year % 100) != 0) || (year % 400) == 0) {
            days = 29
        } else {
            days = 28
        }
    }
    fmt.Fprintf(os.Stdout, "The %d-%d has %d days.\n", year, month, days)
    return
}