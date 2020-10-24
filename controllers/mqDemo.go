package controllers

import (
	"demo/services/mq"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type MqDemoController struct {
	beego.Controller
}

/*
	简单模式和work工作模式push方法
	对应的从队列读取消息方法,写在(/mq/demo/main.go), 因为要单独的去执行
	以当时测试的环境为例,运行这个实例
	首先访问接口 http://127.0.0.1:8099/mq/push 生产者开始向队列推送消息
	然后开启一个或多个命令行窗口,在/crr/golang/beego/demo/mq/demo路径下执行bee run
	此时消费者开始消费队列中的消息
*/
func (c *MqDemoController) Push() {
	// 向队列推送消息, 一秒一次
	go func() {
		count := 0
		for {
			if count >= 100 {
				break
			}
			// ()(队列名)(消息)
			mq.Publish("", "fyouku_demo", "hello"+strconv.Itoa(count))
			count++
			time.Sleep(1 * time.Second)
			fmt.Println(count)

		}
	}()
	// 这里要在浏览器端打印一个消息, 不然可能会执行操作, 这和队列操作无关, 只是为了保证运行
	c.Ctx.WriteString("hello")
}

/*
	订阅(广播)模式push
	对应的从队列读取消息方法,写在(/mq/fanout/main.go), 因为要单独的去执行
	以当时测试的环境为例,运行这个实例
	首先访问接口 http://127.0.0.1:8099/mq/PushFanout 生产者开始向队列推送消息
	然后开启一个或多个命令行窗口,在/crr/golang/beego/demo/mq/fanout 路径下执行bee run
	此时消费者开始消费队列中的消息
	这种模式的消息会给每一个消费者分配一个队列, 推送消息时同时推给每一个队列,从而是每个消费者获得消息
*/
func (c *MqDemoController) PushFanout() {
	go func() {
		count := 0
		for {
			if count >= 100 {
				break
			}
			// (交换机)(模式)(路由)(消息)
			mq.PublishEx("fyouku.demo.fanout", "fanout", "", "fanout"+strconv.Itoa(count))
			count++
			time.Sleep(1 * time.Second)
			fmt.Println(count)
		}
	}()
	c.Ctx.WriteString("fanout")
}

/*
	路由模式push
	其实就是在交换机上指定了路由的key, 用算法实现消息的分发
	不同的路由, 对应不同的消费者
	如果消费者不启动, 则不会创建对应的路由
*/
func (c *MqDemoController) PushDirect() {
	go func() {
		count := 0
		for {
			if count >= 100 {
				break
			}
			// (交换机)(模式)(路由)(消息)
			if count%2 == 0 {
				mq.PublishEx("fyouku.demo.direct", "direct", "two", "direct"+strconv.Itoa(count))
			} else {
				mq.PublishEx("fyouku.demo.direct", "direct", "one", "direct"+strconv.Itoa(count))
			}
			count++
			time.Sleep(1 * time.Second)
			fmt.Println(count)
		}
	}()
	c.Ctx.WriteString("direct")
}

/*
	主题模式push
*/
func (c *MqDemoController) PushTopic() {
	go func() {
		count := 0
		for {
			if count >= 100 {
				break
			}
			// (交换机)(模式)(路由)(消息)
			if count%2 == 0 {
				mq.PublishEx("fyouku.demo.topic", "topic", "fyouku.video", "fyouku.video"+strconv.Itoa(count))
			} else {
				mq.PublishEx("fyouku.demo.topic", "topic", "user.fyouku", "user.fyouku"+strconv.Itoa(count))
			}
			count++
			time.Sleep(1 * time.Second)
			fmt.Println(count)
		}
	}()
	c.Ctx.WriteString("direct")
}
func (c *MqDemoController) PushTopicTwo() {
	go func() {
		count := 0
		for {
			if count >= 100 {
				break
			}
			// (交换机)(模式)(路由)(消息)
			if count%2 == 0 {
				mq.PublishEx("fyouku.demo.topic", "topic", "a.frog.name", "a.frog.name"+strconv.Itoa(count))
			} else {
				mq.PublishEx("fyouku.demo.topic", "topic", "b.frog.name", "b.frog.name"+strconv.Itoa(count))
			}
			count++
			time.Sleep(1 * time.Second)
			fmt.Println(count)
		}
	}()
	c.Ctx.WriteString("direct")
}
