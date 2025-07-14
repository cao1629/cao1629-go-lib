package main

import (
    "fmt"
    "time"
)

// Approach 2: Stop and recreate ticker
// "Pause" gets rid of the old ticker. "Resume" creates a new one.
// We can have different tickers at different times, but we need one channel to receive ticks
// from these different tickers.
func pausableTickerExample2() {
    fmt.Println("\n=== Approach 2: Stop and recreate ticker ===")

    var ticker *time.Ticker
    tickerChan := make(chan time.Time)
    stop := make(chan bool)

    startTicker := func() {
        ticker = time.NewTicker(1 * time.Second)

        go func() {
            fmt.Println("Ticker started")
            defer ticker.Stop()
            for {
                // 有一种可能  tickerChan <-t  ====> stop  ===> <-tickerChan
                select {
                case t := <-ticker.C:
                    tickerChan <- t
                case <-stop:
                    fmt.Printf("Ticker stopped\n")
                    return
                }
            }
        }()
    }

    // 有可能我stopTicker的时候时候 tickerChan里面还有数没有被消费掉
    stopTicker := func() {
        stop <- true
    }

    // consumer
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-tickerChan:
                fmt.Printf("Tick\n")
            case <-done:
                return
            }
        }
    }()

    // Control the ticker
    startTicker()
    time.Sleep(3 * time.Second)

    stopTicker()
    time.Sleep(3 * time.Second)

    startTicker()
    time.Sleep(3 * time.Second)

    stopTicker()

    done <- true
}

func main234234324() {
    //pausableTickerExample1()
    pausableTickerExample2()
    //pausableTickerExample3()
}
