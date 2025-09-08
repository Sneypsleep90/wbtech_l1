package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (s *SafeMap) NewData(key string, data int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = data
}
func (s *SafeMap) GetData(key string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.data[key]
	return value, ok
}
func main() {
	newSafeMap := NewMap()
	var wg sync.WaitGroup
	for i := 1000; i <= 10000; i = i + 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			newSafeMap.NewData(fmt.Sprintf("bitcoin:%d\n", i), i)
		}(i)
	}
	wg.Wait()
	fmt.Println(newSafeMap.data)
}
