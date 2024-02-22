package main

import (
	"sync"
	"testing"
)

func TestCounterRaceCondition(t *testing.T) {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}

func TestCounterWithMutex(t *testing.T) {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}
