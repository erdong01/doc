package selecttimeout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeOut(t *testing.T) {
	result := make(chan string)
	timeout := time.After(3 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		result <- "服务端结果"
	}()
	for {
		select {
		case v := <-result:
			fmt.Println(v)
		case <-timeout:
			fmt.Println("网络访问超时了")
			return
		default:
			fmt.Println("等待...")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestContext(t *testing.T) {
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		worker(ctx, "1号员工")
	}()
	go func() {
		worker(ctx, "2号员工")
	}()

	time.Sleep(time.Second * 5)
	stop()
	time.Sleep(time.Second)
}
func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "下班咯~~~")
			return
		default:
			fmt.Println(name, "认真工作中，请勿打扰...")
		}
		time.Sleep(1 * time.Second)
	}
}
