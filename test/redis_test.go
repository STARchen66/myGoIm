
package test

import (
	"context"
	"github.com/gin-gonic/gin"
	"testing"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

//// TestPublish 测试发布消息到redis
//func TestPublish(t *testing.T) {
//	msg := "当前时间: " + time.Now().Format("15:04:05")
//	err := utils.Publish(ctx, utils.PublishKey, msg)
//	if err != nil {
//
//		fmt.Println(err)
//	}
//}

func channelTest() {

	ch := make(chan int)
	ch2 := make(chan struct{}, 10)

	var data int
	ch <- data

	val := <-ch
	val2, ok := <-ch2
	println(val, val2, ok)
}

func testMain(t *testing.T) {
	main()
}

func main() {

	engine := gin.Default()
	engine.Use(handle)
	engine.POST("/ping", func(ctx *gin.Context) {
		print("pang...")
	})
}

var handle = func(ctx *gin.Context) {
	print("ping...")
}
