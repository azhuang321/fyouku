package models

import (
	"demo/services/mq"
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}
type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

// 保存通知消息
func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	messageId, err := o.Insert(&message)
	return messageId, err
}

// 保存消息接收人
func SendMessageUser(userId int, messageId int64) error {
	o := orm.NewOrm()
	var messageUser MessageUser
	messageUser.UserId = userId
	messageUser.MessageId = messageId
	messageUser.Status = 1
	messageUser.AddTime = time.Now().Unix()
	_, err := o.Insert(&messageUser)
	return err
}

// 发送消息到队列
func SendMessageUserMq(userId int, messageId int64) {
	// 把数据转换成json字符串
	type Data struct {
		UserId    int
		MessageId int64
	}
	var data Data
	data.UserId = userId
	data.MessageId = messageId
	dataJson, _ := json.Marshal(data)
	/*
		生产者 - 发送消息到消息推送队列
		数据库的耗时操作交给异步任务,在o/mq/send/main.go中实现
		这里的消息是发给一个用户, 如果是同时发送给很多用户的场景, 这将非常耗时
	*/
	mq.Publish("", "fyouku_send_message_user", string(dataJson))
}
