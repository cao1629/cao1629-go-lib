package concurrency

import (
    "testing"
    "time"
)

func TestCloseChannel(t *testing.T) {
    done := make(chan struct{})

    go func() {
        for {
            select {
            case <-done:
                t.Log("Done")
            }
        }
    }()

    time.Sleep(time.Second)
    time.Sleep(3 * time.Second)
}
