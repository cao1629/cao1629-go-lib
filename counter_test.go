package main

import (
	"sync"
	"testing"
)

func TestNewCounter(t *testing.T) {
	counter := NewCounter()
	if counter.Get() != 0 {
		t.Errorf("NewCounter() should initialize with value 0, got %d", counter.Get())
	}
}

func TestNewCounterWithValue(t *testing.T) {
	initialValue := int64(42)
	counter := NewCounterWithValue(initialValue)
	if counter.Get() != initialValue {
		t.Errorf("NewCounterWithValue(%d) should initialize with value %d, got %d", initialValue, initialValue, counter.Get())
	}
}

func TestIncrement(t *testing.T) {
	counter := NewCounter()

	// Test single increment
	result := counter.Increment()
	if result != 1 {
		t.Errorf("First increment should return 1, got %d", result)
	}
	if counter.Get() != 1 {
		t.Errorf("Counter value should be 1 after increment, got %d", counter.Get())
	}

	// Test multiple increments
	counter.Increment()
	counter.Increment()
	if counter.Get() != 3 {
		t.Errorf("Counter value should be 3 after three increments, got %d", counter.Get())
	}
}

func TestDecrement(t *testing.T) {
	counter := NewCounterWithValue(5)

	// Test single decrement
	result := counter.Decrement()
	if result != 4 {
		t.Errorf("Decrement should return 4, got %d", result)
	}
	if counter.Get() != 4 {
		t.Errorf("Counter value should be 4 after decrement, got %d", counter.Get())
	}

	// Test decrement to negative
	counter.Set(1)
	counter.Decrement()
	counter.Decrement()
	if counter.Get() != -1 {
		t.Errorf("Counter should handle negative values, got %d", counter.Get())
	}
}

func TestAdd(t *testing.T) {
	counter := NewCounter()

	// Test positive addition
	result := counter.Add(10)
	if result != 10 {
		t.Errorf("Add(10) should return 10, got %d", result)
	}

	// Test negative addition
	result = counter.Add(-3)
	if result != 7 {
		t.Errorf("Add(-3) should return 7, got %d", result)
	}

	// Test zero addition
	result = counter.Add(0)
	if result != 7 {
		t.Errorf("Add(0) should return 7, got %d", result)
	}
}

func TestGet(t *testing.T) {
	counter := NewCounterWithValue(100)

	// Multiple gets should return the same value
	for i := 0; i < 5; i++ {
		if counter.Get() != 100 {
			t.Errorf("Get() should consistently return 100, got %d on iteration %d", counter.Get(), i)
		}
	}
}

func TestSet(t *testing.T) {
	counter := NewCounter()

	// Test setting positive value
	result := counter.Set(50)
	if result != 50 {
		t.Errorf("Set(50) should return 50, got %d", result)
	}
	if counter.Get() != 50 {
		t.Errorf("Counter value should be 50 after Set(50), got %d", counter.Get())
	}

	// Test setting negative value
	counter.Set(-25)
	if counter.Get() != -25 {
		t.Errorf("Counter should handle negative values, got %d", counter.Get())
	}

	// Test setting zero
	counter.Set(0)
	if counter.Get() != 0 {
		t.Errorf("Counter should be 0 after Set(0), got %d", counter.Get())
	}
}

func TestReset(t *testing.T) {
	counter := NewCounterWithValue(42)

	// Test reset returns previous value
	prevValue := counter.Reset()
	if prevValue != 42 {
		t.Errorf("Reset() should return previous value 42, got %d", prevValue)
	}

	// Test counter is now zero
	if counter.Get() != 0 {
		t.Errorf("Counter should be 0 after reset, got %d", counter.Get())
	}

	// Test reset when already zero
	prevValue = counter.Reset()
	if prevValue != 0 {
		t.Errorf("Reset() should return 0 when counter was already 0, got %d", prevValue)
	}
}

func TestConcurrentIncrements(t *testing.T) {
	counter := NewCounter()
	numGoroutines := 100
	incrementsPerGoroutine := 100

	var wg sync.WaitGroup

	// Start multiple goroutines incrementing concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	expectedValue := int64(numGoroutines * incrementsPerGoroutine)
	actualValue := counter.Get()

	if actualValue != expectedValue {
		t.Errorf("Concurrent increments failed: expected %d, got %d", expectedValue, actualValue)
	}
}

func TestConcurrentMixedOperations(t *testing.T) {
	counter := NewCounterWithValue(1000)
	numGoroutines := 50
	operationsPerGoroutine := 100

	var wg sync.WaitGroup

	// Start goroutines with mixed operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				switch j % 4 {
				case 0:
					counter.Increment()
				case 1:
					counter.Decrement()
				case 2:
					counter.Add(2)
				case 3:
					counter.Get() // Read operation
				}
			}
		}(i)
	}

	wg.Wait()

	// We can't predict the exact final value due to mixed operations,
	// but we can verify the counter still works correctly
	initialGet := counter.Get()
	counter.Increment()
	if counter.Get() != initialGet+1 {
		t.Errorf("Counter state corrupted after concurrent operations")
	}
}

func TestConcurrentReadsAndWrites(t *testing.T) {
	counter := NewCounterWithValue(0)
	numReaders := 50
	numWriters := 10
	readsPerGoroutine := 100
	writesPerGoroutine := 10

	var wg sync.WaitGroup

	// Start reader goroutines
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < readsPerGoroutine; j++ {
				counter.Get()
			}
		}()
	}

	// Start writer goroutines
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < writesPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	expectedWrites := int64(numWriters * writesPerGoroutine)
	if counter.Get() != expectedWrites {
		t.Errorf("Expected %d writes, got %d", expectedWrites, counter.Get())
	}
}

func TestDataRace(t *testing.T) {
	// This test should be run with -race flag: go test -race
	counter := NewCounter()
	numGoroutines := 10

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Mix of operations that could cause data races if not properly synchronized
			counter.Increment()
			counter.Get()
			counter.Add(1)
			counter.Decrement()
			counter.Set(counter.Get() + 1)
		}()
	}

	wg.Wait()

	// The exact final value doesn't matter,
	// what matters is that no data race is detected
}

// Benchmark tests
func BenchmarkIncrement(b *testing.B) {
	counter := NewCounter()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
}

func BenchmarkGet(b *testing.B) {
	counter := NewCounterWithValue(1000)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Get()
		}
	})
}

func BenchmarkMixedOperations(b *testing.B) {
	counter := NewCounter()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			switch i % 4 {
			case 0:
				counter.Increment()
			case 1:
				counter.Get()
			case 2:
				counter.Add(1)
			case 3:
				counter.Decrement()
			}
			i++
		}
	})
}
