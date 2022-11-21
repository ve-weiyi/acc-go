package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	logger.Debug("redis init invoke")

	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	res, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("ping 出错：", err)
	}
	logger.Debug(res)

	//1 Set : expiration=0表示无过期时间
	res, err = rdb.Set("name", "lqz", 30*time.Second).Result()
	if err != nil {
		fmt.Println("设置数据失败:", err)
	}
	fmt.Println(res) // OK

	//2 Get
	res, err = rdb.Get("name").Result()
	if err != nil {
		fmt.Println("设置数据失败:", err)
	}
	fmt.Println(res) // OK

	//3 SetNX:key不存在时才设置（新增操作）
	rdb.SetNX("name", 19, 0) // name存在，不会修改

	//4 SetXX:key存在时才设置（修改操作）
	rdb.SetXX("name", "pyy", 0)
	rdb.SetXX("hobby", "football", 0) // 不会新增成功

	// 5 Incr
	rdb.Incr("age")

	// 6 strlen
	l, _ := rdb.StrLen("name").Result()
	fmt.Println(l)
}
