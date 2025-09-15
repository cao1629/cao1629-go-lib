package concurrency

import (
    "fmt"
    "testing"
    "time"
)

func TestGrowingTicker(t *testing.T) {
    interval := 1 * time.Second
    ticker := time.NewTicker(interval)
    growCh := make(chan struct{})
    done := make(chan struct{})

    go func() {
        for {
            select {
            case <-growCh:
                interval += time.Second
                ticker.Reset(interval)
            case tick := <-ticker.C:
                fmt.Printf("Tick at %v with interval %v\n", tick, interval)
            case <-done:
                return
            }
        }
    }()

    time.Sleep(5 * time.Second)
    growCh <- struct{}{}
    time.Sleep(6 * time.Second)
    done <- struct{}{}
}
