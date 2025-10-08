package concurrency

import (
    "context"
    "testing"
    "time"
)

type Result struct {
    Value int
}

// why slow? wait for a message from a channel
func slowOperation(ctx context.Context) (Result, error) {
    select {
    case <-time.After(2 * time.Second):
        return Result{Value: 10}, nil
    case <-ctx.Done():
        return Result{}, ctx.Err()
    }
}

func processWork() (Result, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    return slowOperation(ctx)
}

func task(resultCh chan int) {
    time.Sleep(2 * time.Second)
    resultCh <- 42
}

func TestRunTask(t *testing.T) {
    resultCh := make(chan int)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    go task(resultCh)

    select {
    case <-resultCh:
        t.Log("Task completed")
    case <-ctx.Done():
        t.Log("Task timed out")
    }

}
