package inits

import (
	"fmt"
	"github.com/GUOSAITONG/MyConfig/config"

	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	var err error
	rediss := config.Config.Redis
	addr := fmt.Sprintf("%s:%d", rediss.Host, rediss.Port)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: rediss.Password, // no password set
		DB:       0,               // use default DB
	})
	result, err := config.Rdb.Ping(config.Ctx).Result()
	if err != nil {
		return
	}
	fmt.Println(result, err)
	log.Println("redis 连接成功")
}
