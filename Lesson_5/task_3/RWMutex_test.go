package task_3

import (
	"testing"
)

var countRW = 1000

func getRWPercent(percent int) int {
	if percent <= 0 {
		return 0
	}
	return (countRW * percent) / 100
}

func benchmarkSetUpAddRWMutex(b *testing.B, percent int) {
	var set = SetUpRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(getRWPercent(percent))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRW(1)
			}
		})
	})
}

func benchmarkSetUpHasRWMutex(b *testing.B, percent int) {
	var set = SetUpRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(getRWPercent(percent))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRW(1)
			}
		})
	})
}

func BenchmarkSetUpAddRWMutex(b *testing.B) {
	benchmarkSetUpAddRWMutex(b, 10)
	benchmarkSetUpAddRWMutex(b, 50)
	benchmarkSetUpAddRWMutex(b, 90)
}

func BenchmarkSetUpHasRWMutex(b *testing.B) {
	benchmarkSetUpHasRWMutex(b, 10)
	benchmarkSetUpHasRWMutex(b, 50)
	benchmarkSetUpHasRWMutex(b, 90)
}
