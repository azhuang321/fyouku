package main

import (
	"demo/services/mq"
	"fmt"
)

// 主题模式的的消费者
func main() {
	mq.ConsumerEx("fyouku.demo.topic", "topic", "#", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
