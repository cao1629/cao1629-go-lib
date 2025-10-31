package concurrency

import (
    "fmt"
    "log"
    "sync"
    "testing"
)

// https://go.dev/blog/pipelines
func gen(nums ...int) <-chan int {
    out := make(chan int)

    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()

    return out
}

// square the numbers
func sq(in <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()

    return out
}

// Fan-in: merge multiple channels into one with WaitGroup
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(len(cs))

    drain := func(c <-chan int) {
        for v := range c {
            out <- v
        }
        wg.Done()
    }

    for _, c := range cs {
        go drain(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

// in: {2,3}
// c1, c2, c3 read values from in
// when "in" sends all the values, it is closed.
// then for-loop in "output" exits.
func TestFanInFanOut(t *testing.T) {
    in := gen(2, 3)
    c1 := in
    c2 := in
    c3 := in

    out := merge(c1, c2, c3)

    for v := range out {
        t.Log(v)
    }
}

// explicit cancellation: downstream stages only receive part of values from upstream stages
// when it's enough, tell upstream stages you need to stop. (close upstream channels to avoid resource leak)
func cancellableGen(done chan struct{}, nums ...int) <-chan int {
    out := make(chan int)

    go func() {
        defer close(out)
        for _, n := range nums {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }()

    return out
}

func TestCallableGen(t *testing.T) {
    done := make(chan struct{})
    out := cancellableGen(done, 2, 3, 4, 5)

    fmt.Println(<-out)
    fmt.Println(<-out)

    close(done)
}

func cancellableSq(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        defer close(out)

        for n := range in {
            select {
            case out <- n * n:
            case <-done:
                return
            }
        }
    }()

    return out
}

// fan-in with cancellation
func cancellableMerge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }

    wg.Add(len(cs))

    for _, c := range cs {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func TestExplicitCancellation(t *testing.T) {
    done := make(chan struct{})
    in := cancellableGen(done, 1, 2, 3, 4, 5, 6, 7)

    // fan-out
    sq1 := cancellableSq(done, in)
    sq2 := cancellableSq(done, in)

    // fan-in
    out := cancellableMerge(done, sq1, sq2)

    log.Println(<-out)
    log.Println(<-out)
    log.Println(<-out)

    //for n := range out {
    //   log.Println(n)
    //}

    defer close(done)
}
