package context

import (
	"context"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	ctx := context.Background()
	process(ctx)
	ctx = context.WithValue(ctx, "traceId", "xxml-2022")
	ctx = context.WithValue(ctx, "traceId2", "xxml-2022")
	process(ctx)
}
func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("traceId = %s \n", traceId)
	} else {
		fmt.Printf("no traceId   \n")
	}
}
