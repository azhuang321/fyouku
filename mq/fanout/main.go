package main

import (
	"demo/services/mq"
	"fmt"
)

// 发布订阅(广播)模式的消费者
func main() {
	mq.ConsumerEx("fyouku.demo.fanout", "fanout", "", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
