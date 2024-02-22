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

	// counter 에 동시 접근하는 문제로 1000이 아닐 가능성이 높습니다.
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

	// mutex 를 사용하여 동시 접근 문제를 해결했기 때문에 1000 을 보장합니다.
	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}
