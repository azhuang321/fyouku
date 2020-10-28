package main

import (
	"demo/services/mq"
	"encoding/json"
	"fmt"

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
	/*
		这里要用死信队列的消费者
		("A交换机", "A队列", "B交换机", "B队列", A队列过期时间, 回调)
	*/
	mq.ConsumerDlx("fyouku.comment.count", "fyouku_comment_count", "fyouku.comment.dlx", "fyouku_comment_dlx", 10000, callback)
}
func callback(s string) {

	type Data struct {
		VideoId    int
		EpisodesId int
	}
	var data Data
	var sql string

	/*
		json.Unmarshal将json字符串解码到相应的数据结构
		这里的功能是把队列中的json数据转为自定义类型的数据. 当然, 其中结构是一一对应的.
	*/
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return
	}

	/*
		一下是增加评论数的代码, 和模型中的功能是一样的
	*/
	o := orm.NewOrm()
	// 修改视频的总评论数
	sql = "update video set comment=comment+1 where id=?"
	o.Raw(sql, data.VideoId).Exec()
	// 修改视频剧集的评论数
	sql = "update video_episodes set comment=comment+1 where id=?"
	o.Raw(sql, data.EpisodesId).Exec()
	/*
		更新排行榜 - 通过MQ来来实现
		排行榜是根据评论数排序的, 把video_id推给队列, 表示该视频评论数需要+1.
		这里实现的是生产者, 消费者是/mq/top/main.go
		* 特别说明:用队列的异步方式是为了避免在接口中直接执行Mysql的耗时操作, 所有上面的Mysql操作应该挪到异步任务中执行
	*/
	// 实时增加评论数
	videoObj := map[string]int{
		"videoId": data.VideoId,
	}
	videoJson, _ := json.Marshal(videoObj)
	mq.Publish("", "fyouku_top", string(videoJson))

	fmt.Printf("msg is :%s\n", s) // 测试中使用的代码, 上线可注释掉
}
