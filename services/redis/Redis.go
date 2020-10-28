package redisClient

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

// 直接连接
func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
	return pool
}

// 连接池
func PoolConnect() redis.Conn {
	pool := &redis.Pool{
		MaxIdle:     1,                 // 最大的空闲连接数
		MaxActive:   10,                // 最大连接数
		IdleTimeout: 180 * time.Second, // 空闲链接超时时间
		Wait:        true,              // 超过最大连接数时, true-等待, false-报错
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
			// c, err := redis.Dial("tcp", "192.168.11.125:6379", redis.DialDatabase(1), redis.DialPassword("123456"))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return pool.Get()
}

/*
改造 - 接口总的编码示例:

/crr/golang/beego/demo/models/video.go
	// 增加redis缓存 - 获取视频详情
	func RedisGetVideoInfo(videoId int) (Video, error) {
/crr/golang/beego/demo/controllers/video.go
	video, err := models.RedisGetVideoInfo(videoId)



*/
