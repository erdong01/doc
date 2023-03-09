package selecttimeout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	ctx, stop := context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		worker(ctx, "打工人1")
	}()
	go func() {
		worker(ctx, "打工人2")
	}()

	time.Sleep(5 * time.Second)
	stop()
	fmt.Println("???")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("下班了")
			return
		default:
			fmt.Println("认证摸鱼中，请切勿打扰。。。")
		}
		time.Sleep(1 * time.Second)
	}
}
