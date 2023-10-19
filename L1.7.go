package main

//Реализовать конкурентную запись данных в map.
import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (sm *SafeMap) Put(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.m[key]
}

func main() {
	safeMap := SafeMap{
		m: make(map[string]int),
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		safeMap.Put("key1", 1)
	}()

	go func() {
		defer wg.Done()
		safeMap.Put("key2", 2)
	}()

	wg.Wait()

	fmt.Println(safeMap.Get("key1"))
	fmt.Println(safeMap.Get("key2"))
}
