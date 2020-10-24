package controllers

import (
	redisClient "demo/services/redis"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type RedisDemoController struct {
	beego.Controller
}

func (c *RedisDemoController) Demo() {
	cli := redisClient.PoolConnect()
	defer cli.Close() // 函数退出时调用

	// 创建键值对
	_, err := cli.Do("SET", "username", "frog")
	if err == nil {
		// 设置过期时间
		cli.Do("expire", "username", 1000)
	}

	// 读取缓存的值
	r, err := redis.String(cli.Do("get", "username"))
	if err == nil {
		fmt.Println(1)
		fmt.Println(r)
		// 获取过期时间
		ttl, _ := redis.Int64(cli.Do("ttl", "username"))
		fmt.Println(ttl)
	} else {
		fmt.Println(2)
		fmt.Println(err)
	}
}
