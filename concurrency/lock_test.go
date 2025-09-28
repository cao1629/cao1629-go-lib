package concurrency

import (
    "log"
    "sync"
    "testing"
    "time"
)

func TestReentrantLock(t *testing.T) {
    lock := sync.Mutex{}

    go func() {
        lock.Lock()
        lock.Lock()
        lock.Unlock()
    }()

    time.Sleep(time.Second)

    go func() {
        lock.Lock()
        log.Println("lock acquired")
    }()

    time.Sleep(5 * time.Second)
}
