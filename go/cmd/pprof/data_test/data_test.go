package data_test

import (
	"doc/go/cmd/pprof"
	"testing"
)

func BenchmarkData(b *testing.B) {
	b.Run("block", func(b *testing.B) {
		o := pprof.Block{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})

	b.Run("cpu", func(b *testing.B) {
		o := pprof.Cpu{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("mem", func(b *testing.B) {
		o := pprof.Mem{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})

	b.Run("goroutine", func(b *testing.B) {
		o := pprof.Goroutine{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("mutex", func(b *testing.B) {
		o := pprof.Mutex{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
}
