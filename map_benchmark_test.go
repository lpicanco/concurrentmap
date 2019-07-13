package concurrentmap

import (
	"sync"
	"testing"
)

func BenchmarkSyncMapPut(b *testing.B) {
	var m sync.Map

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
}

func BenchmarkPut(b *testing.B) {
	m := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}
}

func BenchmarkSyncMapGet(b *testing.B) {
	var m sync.Map

	for i := 0; i < 100; i++ {
		m.Store(i, 42)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Load(i)
	}
}

func BenchmarkGet(b *testing.B) {
	m := New()

	for i := 0; i < 100; i++ {
		m.Put(i, 42)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}

func BenchmarkSyncMapPutGetConcurrent(b *testing.B) {
	var m sync.Map

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Store(i, 42)
			m.Load(i)
			i++
		}
	})
}

func BenchmarkPutGetConcurrent(b *testing.B) {
	m := New()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Put(i, 42)
			m.Get(i)
			i++
		}
	})
}

func BenchmarkSyncMapGetPutRemoveConcurrent(b *testing.B) {
	var m sync.Map

	var wg sync.WaitGroup
	wg.Add(b.N * 2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go func(i int) {
			m.Store(i, i)
			wg.Done()
		}(i)

		if i%10 == 3 {
			wg.Add(1)
			go func(i int) {
				m.Delete(i)
				wg.Done()
			}(i - 1)
		}

		go func(i int) {
			m.Load(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func BenchmarkGetPutRemoveConcurrent(b *testing.B) {
	m := New()

	var wg sync.WaitGroup
	wg.Add(b.N * 2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go func(i int) {
			m.Put(i, i)
			wg.Done()
		}(i)

		if i%10 == 3 {
			wg.Add(1)
			go func(i int) {
				m.Remove(i)
				wg.Done()
			}(i - 1)
		}

		go func(i int) {
			m.Get(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
