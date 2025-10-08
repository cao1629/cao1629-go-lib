package concurrency

import (
    "fmt"
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

type Data struct {
    val int
}

func TestChannelCopy(t *testing.T) {

    ch := make(chan Data)

    data := Data{1}
    go func() {
        ch <- data
    }()

    d := <-ch
    data.val = 100
    fmt.Println(d.val)
}

func TestChannelCopy2(t *testing.T) {

    ch := make(chan *Data)

    data := &Data{1}
    go func() {
        ch <- data
    }()

    d := <-ch
    data.val = 100
    fmt.Println(d.val)
}

func TestChannel3(t *testing.T) {
    ch := make(chan struct{})

    go func() {
        ch <- struct{}{}
    }()

    <-ch
    fmt.Println("hello")

    time.Sleep(3 * time.Second)
}

func TestPingPong(t *testing.T) {
    ping := make(chan struct{})
    pong := make(chan struct{})

    go func() {
        for {
            <-ping
            fmt.Println("Ping")
            time.Sleep(time.Second)
            pong <- struct{}{}
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second)
            ping <- struct{}{}
            <-pong
            fmt.Println("Pong")

        }
    }()

    time.Sleep(10 * time.Second)
}

func Test4(t *testing.T) {

}
