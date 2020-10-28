package main

import (
	"demo/models"
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
	mq.Consumer("", "fyouku_send_message_user", callback)
}
func callback(s string) {
	type Data struct {
		UserId    int
		MessageId int64
	}
	var data Data
	/*
		json.Unmarshal将json字符串解码到相应的数据结构
		这里的功能是把队列中的json数据转为自定义类型的数据. 当然, 其中结构是一一对应的.
	*/
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return
	}
	// 调取模型, 完成消息推送
	models.SendMessageUser(data.UserId, data.MessageId)
	fmt.Printf("msg is :%s\n", s) // 测试中使用的代码, 上线可注释掉
}
