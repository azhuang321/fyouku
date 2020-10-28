package main

import (
	"demo/services/mq"
	"fmt"
)

// 路由模式的消费者demo
func main() {
	mq.ConsumerEx("fyouku.demo.direct", "direct", "one", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
