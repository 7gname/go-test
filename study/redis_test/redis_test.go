package redis_test

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
	"strconv"
	"testing"
	"math/rand"
)

const KeyPre  = "test:ms:count:"

func initData() <-chan string {
	ch := make(chan string)
	cli := redis.NewClient(&redis.Options{
		Addr:"10.2.18.162:6379",
	})
	go func() {
		var i int64 = 0
		for ; i < 3600; i++ {
			key := fmt.Sprintf("%s%d",KeyPre, time.Now().Unix() + i)
			cli.Set(key, rand.Intn(1000), time.Hour)
		}
		ch <- "over"
	}()
	return ch
}

func TestPipeline(t *testing.T)  {
	cli := redis.NewClient(&redis.Options{
		Addr:"10.2.18.162:6379",
	})
	defer cli.Close()
	t.Logf("redis connected\n")

	//t.Logf("start init data\n")
	//初始化数据
	//<-initData()
	//t.Logf("init data end\n")

	pipe := cli.Pipeline()
	var i int64
	for i = 0; i < 10; i++ {
		key := fmt.Sprintf("%s%d", KeyPre, time.Now().Unix() + i)
		pipe.Get(key)
	}

	cmders, err := pipe.Exec()
	if err != nil {
		t.Errorf("Err[%+v]", err)
	}
	var total int64
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringCmd)
		val, err := cmd.Result()
		fmt.Printf("cmd[%s], val[%s], err[%v]\n",cmder.Args(), val, err)
		v, _ := strconv.ParseInt(val, 10, 64)
		total += v
	}
	fmt.Printf("[%s]last minute total:[%d]\n", time.Now().Format("2006-01-02 15:04:05"), total)
}