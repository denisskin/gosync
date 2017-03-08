package gosync

import (
	"sync"
	"testing"
)

func BenchmarkWaitingViaEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {

		ev := NewEvent()

		go func() {
			// do somethings
			ev.Done()
		}()

		ev.Wait()
	}
}

func BenchmarkWaitingViaWaitGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			// do somethings
			wg.Done()
		}()

		wg.Wait()
	}
}

func BenchmarkWaitingViaMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {

		var mx sync.Mutex

		mx.Lock()
		go func() {
			// do somethings
			mx.Unlock()
		}()

		mx.Lock()
		mx.Unlock()
	}
}
