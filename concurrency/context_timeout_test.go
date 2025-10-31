package concurrency

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func TestContextTimeout(t *testing.T) {

    callRPC := func(resultCh chan int) {
        time.Sleep(2 * time.Second)
        resultCh <- 100
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resultCh := make(chan int)

    go callRPC(resultCh)

    select {
    case <-ctx.Done():
        t.Log("timeout")
    case result := <-resultCh:
        t.Log("result = ", result)
    }

    time.Sleep(2 * time.Second)
}

// an RPC call with timeout mechanism
func getValue() (int, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resultCh := make(chan int)

    go sendRpcRequest(resultCh)

    select {
    case val := <-resultCh:
        return val, nil
    case <-ctx.Done():
        return 0, ctx.Err()
    }
}

// Simulate an RPC call
func sendRpcRequest(resultCh chan int) {
    time.Sleep(2 * time.Second)
    resultCh <- 100
}

func TestRpcTimeout(t *testing.T) {
    if result, err := getValue(); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
