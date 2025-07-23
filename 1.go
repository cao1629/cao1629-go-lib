package main

import (
    "fmt"
    "time"
)

type ElectionTicker struct {
    timer *time.Timer
    C     chan time.Time
}

func MakeElectionTicker() *ElectionTicker {
    et := &ElectionTicker{
        C: make(chan time.Time),
    }
    return et
}

func (et *ElectionTicker) Reset(interval time.Duration) {
    et.Stop()
    et.timer = time.NewTimer(interval)
    go func() {
        tick := <-et.timer.C
        et.C <- tick
    }()
}

func (et *ElectionTicker) Stop() {
    if et.timer != nil {
        et.timer.Stop()
    }
}

func main() {
    ticker := MakeElectionTicker()

    go func() {
        for {
            select {
            case <-ticker.C:
                fmt.Println(time.Now().Format("15:04:05"))
            }
        }
    }()

    go func() {
        ticker.Reset(time.Second)
    }()

    go func() {
        time.Sleep(3 * time.Second)
        ticker.Reset(2 * time.Second)
    }()

    time.Sleep(10 * time.Second)
}
