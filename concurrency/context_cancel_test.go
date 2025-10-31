package concurrency

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func TestWithCancel(t *testing.T) {

    work := func(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                fmt.Println("working....")
                time.Sleep(time.Second)
            }
        }
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go work(ctx)
    go work(ctx)
    time.Sleep(3 * time.Second)
}
