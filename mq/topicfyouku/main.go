package main

import (
	"demo/services/mq"
	"fmt"
)

// 主题模式的消费者demo
func main() {
	mq.ConsumerEx("fyouku.demo.topic", "topic", "fyouku.*", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
