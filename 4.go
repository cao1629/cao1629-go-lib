package main

import (
    "fmt"
    "time"
)

// Teddy represents a dog that barks until fed
type Teddy struct {
    fed chan struct{}
}

// StartBarking makes Teddy bark every second until fed
func (t *Teddy) StartBarking() {
    fmt.Printf("Teddy is hungry and starts barking!\n")

    // Create a C that ticks every second
    ticker := time.Tick(1 * time.Second)

    for {
        select {
        case <-t.fed:
            fmt.Printf("Teddy got fed and stopped barking!\n")
            return
        case <-ticker:
            fmt.Printf("Woof!\n")
        }
    }
}

// Feed feeds Teddy and stops the barking
func (t *Teddy) Feed() {
    fmt.Printf("Feeding...\n")
    select {
    case t.fed <- struct{}{}:
    default:
        fmt.Printf("Teddy is already fed.\n")
    }
}

func main4() {
    teddy := Teddy{fed: make(chan struct{})}

    go func() {
        time.Sleep(5 * time.Second)
        teddy.Feed()
    }()

    time.Sleep(10 * time.Second)
    fmt.Println("Program finished")
}
