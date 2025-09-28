package main

import (
    "sync"
    "time"
)

type NumberProducer struct {
    mu       sync.RWMutex
    running  bool
    output   chan int
    stopChan chan struct{}
}

func NewNumberProducer() *NumberProducer {
    return &NumberProducer{
        output:   make(chan int),
        stopChan: make(chan struct{}),
    }
}

func (p *NumberProducer) SetRunning(running bool) {
    p.mu.Lock()
    p.running = running
    p.mu.Unlock()
}

func (p *NumberProducer) IsRunning() bool {
    p.mu.RLock()
    defer p.mu.RUnlock()
    return p.running
}

func (p *NumberProducer) Start() {
    go p.produce()
}

func (p *NumberProducer) Stop() {
    close(p.stopChan)
    close(p.output)
}

func (p *NumberProducer) Output() <-chan int {
    return p.output
}

func (p *NumberProducer) produce() {
    counter := 0
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-p.stopChan:
            return
        case <-ticker.C:
            if p.IsRunning() {
                select {
                case p.output <- counter:
                    counter++
                case <-p.stopChan:
                    return
                }
            }
        }
    }
}
