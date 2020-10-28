package main

import (
	"demo/services/mq"
	"fmt"
)

// 简短模式及worker模式的消费者demo
func main() {
	// 两个交换机器10秒
	mq.ConsumerDlx("fyouku.dlx.a", "fyouku_dlx_a", "fyouku.dlx.b", "fyouku_dlx_b", 10000, callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
