package main

import (
    "fmt"
    "time"
)

func main333() {
    fmt.Println("=== Channel Direction Types Demo ===")

    // Create a bidirectional channel
    ch := make(chan int, 2)

    // Example 1: Send-only channel (chan<- int)
    go producer(ch) // Pass bidirectional channel to send-only parameter

    // Example 2: Receive-only channel (<-chan int)
    go consumer(ch) // Pass bidirectional channel to receive-only parameter

    // Example 3: Bidirectional channel (chan int)
    go bidirectionalWorker(ch)

    // Give goroutines time to work
    time.Sleep(100 * time.Millisecond)

    // Close the channel
    close(ch)

    fmt.Println("Demo completed!")
}

// Producer function - can only SEND to channel
func producer(sendChan chan<- int) {
    fmt.Println("Producer: Sending values...")
    sendChan <- 10
    sendChan <- 20

    // This would cause a compile error:
    // value := <-sendChan  // ERROR: invalid operation

    fmt.Println("Producer: Done sending")
}

// Consumer function - can only RECEIVE from channel
func consumer(receiveChan <-chan int) {
    fmt.Println("Consumer: Receiving values...")

    for i := 0; i < 2; i++ {
        value := <-receiveChan
        fmt.Printf("Consumer: Received %d\n", value)
    }

    // This would cause a compile error:
    // receiveChan <- 30  // ERROR: invalid operation

    fmt.Println("Consumer: Done receiving")
}

// Bidirectional worker - can both send and receive
func bidirectionalWorker(ch chan int) {
    fmt.Println("Bidirectional: Can do both operations")

    // Can send
    ch <- 99

    // Can receive
    value := <-ch
    fmt.Printf("Bidirectional: Got %d\n", value)
}

// Real-world example: Pipeline pattern
func demonstratePipeline() {
    fmt.Println("\n=== Pipeline Pattern Example ===")

    // Stage 1: Generate numbers
    numbers := make(chan int)
    go generateNumbers(numbers)

    // Stage 2: Square the numbers
    squares := make(chan int)
    go squareNumbers(numbers, squares)

    // Stage 3: Print results
    printResults(squares)
}

// Stage 1: Only sends numbers (producer)
func generateNumbers(out chan<- int) {
    defer close(out)
    for i := 1; i <= 5; i++ {
        out <- i
    }
}

// Stage 2: Receives from one channel, sends to another
func squareNumbers(in <-chan int, out chan<- int) {
    defer close(out)
    for num := range in {
        out <- num * num
    }
}

// Stage 3: Only receives (consumer)
func printResults(in <-chan int) {
    for result := range in {
        fmt.Printf("Result: %d\n", result)
    }
}

// Comparison with your Work function
func yourWorkFunction(num int, resultChan chan<- int) {
    // This is a send-only channel parameter
    // Function can only send results, not receive

    var sum int
    for i := 0; i < num; i++ {
        sum += i
    }

    resultChan <- sum // ✅ Can send
    // value := <-resultChan  // ❌ Would not compile
}

func alternativeWithBidirectional(num int, resultChan chan int) {
    // This accepts bidirectional channel
    // Function CAN both send and receive (but shouldn't receive in this case)

    var sum int
    for i := 0; i < num; i++ {
        sum += i
    }

    resultChan <- sum // ✅ Can send
    // value := <-resultChan  // ✅ Can receive (but logically wrong)
}
