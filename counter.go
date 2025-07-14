package main

import (
    "sync"
)

// Counter represents a thread-safe counter
type Counter struct {
    mu    sync.RWMutex
    value int64
}

// NewCounter creates a new counter with initial value of 0
func NewCounter() *Counter {
    return &Counter{
        value: 0,
    }
}

// NewCounterWithValue creates a new counter with the specified initial value
func NewCounterWithValue(initial int64) *Counter {
    return &Counter{
        value: initial,
    }
}

// Increment increases the counter by 1 and returns the new value
func (c *Counter) Increment() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
    return c.value
}

// Decrement decreases the counter by 1 and returns the new value
func (c *Counter) Decrement() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value--
    return c.value
}

// Add adds the specified value to the counter and returns the new value
func (c *Counter) Add(delta int64) int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value += delta
    return c.value
}

// Get returns the current value of the counter
func (c *Counter) Get() int64 {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.value
}

// Set sets the counter to the specified value and returns the new value
func (c *Counter) Set(value int64) int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value = value
    return c.value
}

// Reset sets the counter to 0 and returns the previous value
func (c *Counter) Reset() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    prev := c.value
    c.value = 0
    return prev
}
