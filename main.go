package main

import (
    "fmt"
    "sync"
    "time"
)

func main34341() {
    // Create a new counter
    counter := NewCounter()

    // Demonstrate basic operations
    fmt.Println("Initial value:", counter.Get())
    fmt.Println("After increment:", counter.Increment())
    fmt.Println("After adding 5:", counter.Add(5))
    fmt.Println("Current value:", counter.Get())
    fmt.Println("After decrement:", counter.Decrement())
    fmt.Println("After setting to 100:", counter.Set(100))

    // Demonstrate thread safety with concurrent operations
    fmt.Println("\nTesting thread safety with 100 goroutines...")

    var wg sync.WaitGroup
    numGoroutines := 100
    incrementsPerGoroutine := 10

    // Reset counter for the test
    prevValue := counter.Reset()
    fmt.Printf("Reset counter, previous value was: %d\n", prevValue)

    // Start multiple goroutines that increment the counter
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < incrementsPerGoroutine; j++ {
                counter.Increment()
                time.Sleep(time.Microsecond) // Small delay to increase chance of race conditions
            }
            fmt.Printf("Goroutine %d finished\n", id)
        }(i)
    }

    // Wait for all goroutines to complete
    wg.Wait()

    expectedValue := int64(numGoroutines * incrementsPerGoroutine)
    actualValue := counter.Get()

    fmt.Printf("\nExpected final value: %d\n", expectedValue)
    fmt.Printf("Actual final value: %d\n", actualValue)

    if actualValue == expectedValue {
        fmt.Println("Thread safety test passed!")
    } else {
        fmt.Println("Thread safety test failed!")
    }
}
