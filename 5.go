package main

//
//import (
//    "fmt"
//    "time"
//)
//
//func main() {
//    feed := make(chan struct{})
//    done := make(chan struct{})
//
//    go dogTeddy(feed, done)
//
//    // Simulate feeding Teddy every 15 seconds
//    for {
//        time.Sleep(15 * time.Second)
//        fmt.Println("You fed Teddy.")
//        feed <- struct{}{}
//    }
//}
//
//func dogTeddy(feed <-chan struct{}, done chan struct{}) {
//    const hungerTimeout = 10 * time.Second
//    const barkInterval = 1 * time.Second
//
//    lastFed := time.Now()
//    barking := false
//    barkTicker := time.NewTicker(barkInterval)
//    barkTicker.Stop() // Start stopped
//
//    for {
//        select {
//        case <-feed:
//            lastFed = time.Now()
//            if barking {
//                fmt.Println("Teddy stops barking. Thank you for the food!")
//                barkTicker.Stop()
//                barking = false
//            }
//        case <-time.After(1 * time.Second):
//            if !barking && time.Since(lastFed) >= hungerTimeout {
//                fmt.Println("Teddy is hungry and starts barking!")
//                barkTicker = time.NewTicker(barkInterval)
//                barking = true
//            }
//        case <-barkTicker.C:
//            if barking {
//                fmt.Println("Teddy: Woof! Feed me!")
//            }
//        }
//    }
//}
