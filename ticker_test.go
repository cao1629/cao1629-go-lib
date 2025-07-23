package main

import (
    "fmt"
    "testing"
    "time"
)

func TestTicker(t *testing.T) {
    start := time.Now()
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    elapsed := func() time.Duration {
        return time.Since(start).Round(time.Millisecond)
    }
    for {
        select {
        case <-tick:
            fmt.Printf("[%6s] tick.\n", elapsed())
        case <-boom:
            fmt.Printf("[%6s] BOOM!\n", elapsed())
            return
        default:
            fmt.Printf("[%6s]     .\n", elapsed())
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func TestTicker2(t *testing.T) {

    // Create a C that fires every 2 seconds
    ticker := time.Tick(2 * time.Second)

    // Listen for ticks
    for i := 0; i < 5; i++ {
        t := <-ticker
        fmt.Printf("Tick at %s\n", t.Format("15:04:05"))
    }

    fmt.Println("Done with time.Tick example")
}
