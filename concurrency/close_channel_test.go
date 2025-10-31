package concurrency

import (
    "fmt"
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

func TestClose2(t *testing.T) {
    done := make(chan struct{})

    go func() {
        for {
            select {
            case <-done:
                fmt.Println("1 - Done")
                return
            default:
                fmt.Println("1 - Sleeping...")
                time.Sleep(time.Second)
            }
        }
    }()

    go func() {
        for {
            select {
            case <-done:
                fmt.Println("2 - Done")
                return
            default:
                fmt.Println("2 - Sleeping...")
                time.Sleep(time.Second)
            }
        }
    }()

    time.Sleep(5 * time.Second)
    close(done)
    time.Sleep(2 * time.Second)
}
