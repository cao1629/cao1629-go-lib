package concurrency

import (
    "context"
    "errors"
    "math/rand/v2"
    "time"
)

// Simulate fetching a random number from a remote service.
// Three possible cases:
// case 1: return a random number.
// case 2: return an error.
// case 3: time out. which is the work of the caller of this method.
func fetchRemoteRandomNumber() (int, error) {
    respTime := time.Duration(rand.IntN(2000)) * time.Millisecond
    time.Sleep(respTime)

    if rand.IntN(10) < 2 {
        return -1, errors.New("service unavailable")
    } else {
        return rand.IntN(100), nil
    }
}

// GetSquaredRandomNumber calls fetchRemoteRandomNumber with a timeout and returns
// the square of the random number on success, or an error otherwise.
func GetSquaredRandomNumber(timeout time.Duration) (int, error) {
    type result struct {
        num int
        err error
    }

    // why buffered channel?
    // because the following goroutine may be blocked forever if the timeout happens first
    resultCh := make(chan result, 1)

    go func() {
        num, err := fetchRemoteRandomNumber()
        resultCh <- result{num: num, err: err}
    }()

    select {
    case res := <-resultCh:
        if res.err != nil {
            return 0, res.err
        }
        return res.num * res.num, nil
    case <-time.After(timeout):
        return 0, errors.New("request timed out")
    }
}

// GetSquaredRandomNumberWithContext calls fetchRemoteRandomNumber using context with timeout
// and returns the square of the random number on success, or an error otherwise.
func GetSquaredRandomNumberWithContext(ctx context.Context, timeout time.Duration) (int, error) {
    type result struct {
        num int
        err error
    }

    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    resultCh := make(chan result, 1)

    go func() {
        num, err := fetchRemoteRandomNumber()
        resultCh <- result{num: num, err: err}
    }()

    select {
    case res := <-resultCh:
        if res.err != nil {
            return 0, res.err
        }
        return res.num * res.num, nil
    case <-ctx.Done():
        return 0, ctx.Err()
    }
}

// SumFiveRandomNumbers calls fetchRemoteRandomNumber 5 times and returns the sum
// of all random numbers. If any call returns an error or times out, returns an error.
func SumFiveRandomNumbers(timeout time.Duration) (int, error) {
    type result struct {
        num int
        err error
    }

    resultCh := make(chan result, 5)

    for i := 0; i < 5; i++ {
        go func() {
            num, err := fetchRemoteRandomNumber()
            resultCh <- result{num: num, err: err}
        }()
    }

    sum := 0
    for i := 0; i < 5; i++ {
        select {
        case res := <-resultCh:
            if res.err != nil {
                return 0, res.err
            }
            sum += res.num
        case <-time.After(timeout):
            return 0, errors.New("request timed out")
        }
    }

    return sum, nil
}

// MajorityVote sends RPC to 5 servers and returns true when the majority
// (3 or more) respond with success, or false when the majority respond with errors.
// Stops waiting once a majority is reached.
func MajorityVote() (bool, error) {
    resultCh := make(chan error, 5)

    for i := 0; i < 5; i++ {
        go func() {
            _, err := fetchRemoteRandomNumber()
            resultCh <- err
        }()
    }

    successCount := 0
    failureCount := 0
    majority := 3

    for i := 0; i < 5; i++ {
        err := <-resultCh
        if err == nil {
            successCount++
            if successCount >= majority {
                return true, nil
            }
        } else {
            failureCount++
            if failureCount >= majority {
                return false, nil
            }
        }
    }

    // Should never reach here since we check for majority in the loop
    return successCount >= majority, nil
}
