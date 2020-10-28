package main

import (
	"demo/models"
	"demo/services/mq"
	"encoding/json"
	"fmt"
	"strconv"

	redis "demo/services/redis"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
import (
	"demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
*/
func main() {
	/*
		引入beego框架
	*/
	// 引入配置文件
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)
	// 消费者: 对应/models/comment.go -> func SaveComment方法中的生产者
	mq.Consumer("", "fyouku_top", callback)
}
func callback(s string) {
	type Data struct {
		VideoId int
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	// 获取视频数据
	videoInfo, err := models.RedisGetVideoInfo(data.VideoId)
	if err != nil {
		return
	}
	conn := redis.PoolConnect()
	defer conn.Close()

	// 更新排行榜
	/*
		排行榜的对应有序集合的key
	*/
	redisChannelKey := "video:top:channel:channelId:" + strconv.Itoa(videoInfo.ChannelId)
	redisTypeKey := "video:top:type:typeId:" + strconv.Itoa(videoInfo.TypeId)
	/*
		对有序集合中指定成员的分数加上增量
		这两个都是在修改排行榜中某视频的排序值
	*/
	conn.Do("zincrby", redisChannelKey, 1, data.VideoId)
	conn.Do("zincrby", redisTypeKey, 1, data.VideoId)
	// 测试中使用的代码, 上线可注释掉
	fmt.Printf("msg is :%s\n", s)
}
