package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type GoDemoController struct {
	beego.Controller
}

func (c *GoDemoController) Demo() {
	for i := 0; i < 10; i++ {
		// go cal(i)
		go func(i int) {
			fmt.Printf("i = %d\n", i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	c.Ctx.WriteString("demo")
}

/*
	Channel的demo, 并发写入个数字, 再按照写入顺序打印出来
*/
func (c *GoDemoController) ChannelDemo() {
	Channel := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int, Channel chan bool) {
			fmt.Printf("i = %d\n", i)
			time.Sleep(2 * time.Second)
			Channel <- true
		}(i, Channel)
	}
	for j := 0; j < 10; j++ {
		bo := <-Channel
		fmt.Println(bo)
	}
	close(Channel)
	c.Ctx.WriteString("ChannelDemo")
}

/*
	select调度器
*/
func (c *GoDemoController) SelectDemo() {
	oneChannel := make(chan int, 20)
	twoChannel := make(chan string, 20)
	go func(oneChannel chan int) {
		for i := 0; i < 10; i++ {
			oneChannel <- i
			fmt.Println("<- one = ", i)
		}
	}(oneChannel)

	twoChannel <- "a"
	twoChannel <- "b"

	go func() {
		for i := 0; i < 30; i++ {
			select {
			case a := <-oneChannel:
				fmt.Println("get one = ", a)
			case b := <-twoChannel:
				fmt.Println("get two = ", b)
			default:
				fmt.Println("no message")
			}
		}
	}()
	c.Ctx.WriteString("SelectDemo")
}

// 模拟任务池
func (c *GoDemoController) TaskDemo() {
	// 接收任务
	taskChannel := make(chan int, 20)
	// 处理任务
	resChannel := make(chan int, 20)
	// 关闭任务
	closeChannel := make(chan bool, 5)

	// 造一个任务队列
	go func() {
		for i := 0; i < 50; i++ {
			taskChannel <- i
		}
		close(taskChannel)
	}()
	// 5个协程同时处理taskChannel
	for i := 0; i < 5; i++ {
		go func(taskChannel chan int, resChannel chan int, closeChannel chan bool, i int) {
			for t := range taskChannel {
				resChannel <- t
				fmt.Println("do ", i, t)
			}
			closeChannel <- true
		}(taskChannel, resChannel, closeChannel, i)
	}
	// 异步监听器: 监听这5个任务全部完成
	go func() {
		for i := 0; i < 5; i++ {
			<-closeChannel
			fmt.Println("<-closeChannel ", i)
		}
		close(resChannel)
		close(closeChannel)
	}()
	// for循环channel时, 当channel关闭以后会退出循环
	for r := range resChannel {
		fmt.Println("res:", r)
	}
	c.Ctx.WriteString("TaskDemo")
}

/*
改造
/crr/golang/beego/demo/controllers/comment.go 中List接口
*/
