package concurrency

import (
    "testing"
    "time"
)

func TestTwoProducer(t *testing.T) {

    ch := make(chan int)

    go func() {
        for {
            for i := 0; i < 20; i++ {
                ch <- 1
            }
            time.Sleep(time.Second)
        }
    }()

    go func() {
        for {
            for i := 0; i < 20; i++ {
                ch <- 2
            }
            time.Sleep(time.Second)
        }
    }()

    go func() {
        for v := range ch {
            t.Log(v)
        }
    }()

    time.Sleep(5 * time.Second)
}
