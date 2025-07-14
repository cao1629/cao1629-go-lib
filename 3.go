package main

//
//import (
//    "fmt"
//    "time"
//)
//
//// Teddy represents a dog with hunger and barking behavior
//type Teddy struct {
//    hungry   chan struct{}
//    fed      chan struct{}
//    barkDone chan struct{}
//    ticker   *time.Ticker
//}
//
//func MyTeddy() *Teddy {
//    myTeddy := &Teddy{
//        hungry:   make(chan struct{}),
//        fed:      make(chan struct{}),
//        barkDone: make(chan struct{}),
//        ticker:   time.NewTicker(10 * time.Second),
//    }
//    go myTeddy.hungerCycle()
//    go myTeddy.barkingBehavior()
//    return myTeddy
//}
//
//func (t *Teddy) hungerCycle() {
//    defer t.ticker.Stop()
//
//    // Teddy starts hungry
//    t.hungry <- struct{}{}
//
//    for {
//        <-t.fed // Wait until fed
//        select {
//        case <-t.ticker.C:
//            t.hungry <- struct{}{} // Teddy gets hungry again
//        }
//    }
//}
//
//// barkingBehavior manages Teddy's barking when hungry
//func (t *Teddy) barkingBehavior() {
//    for {
//        // wait for Teddy to become hungry
//        <-t.hungry
//        fmt.Printf("Teddy is hungry and starts barking!\n")
//
//        t.barkDone = make(chan struct{})
//
//        // Start barking goroutine with ticker
//        go func() {
//            barkTicker := time.NewTicker(1 * time.Second)
//            defer barkTicker.Stop()
//
//            for {
//                select {
//                case <-t.barkDone:
//                    return
//                case <-barkTicker.C:
//                    fmt.Printf("Woof!\n")
//                }
//            }
//        }()
//
//        <-t.fed // Wait until Teddy is fed
//        close(t.barkDone)
//        fmt.Printf("%s is full and stops barking.\n")
//    }
//}
//
//// Feed feeds Teddy
//func (t *Teddy) Feed() {
//    fmt.Printf("Feeding %s...\n")
//    t.fed <- struct{}{}
//}
//
//// WaitForHunger waits for Teddy to become hungry
//func (t *Teddy) WaitForHunger() {
//    <-t.hungry
//}
//
//// Teddy's behavior simulation
//func teddyBehavior() {
//    teddy := MyTeddy()
//
//    // Simulate feeding Teddy after 5 seconds of barking
//    for {
//        teddy.WaitForHunger()
//        time.Sleep(5 * time.Second)
//        teddy.Feed()
//    }
//}
//
//func main() {
//    // Run Teddy's behavior simulation
//    teddyBehavior()
//}
