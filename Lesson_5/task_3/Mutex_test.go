package task_3

import (
	"testing"
)

var counts = 1000

func getPercent(percent int) int {
	if percent <= 0 {
		return 0
	}
	return int((counts * percent) / 100)
}

func benchmarkSetUpAddMutex(b *testing.B, percent int) {
	var set = SetUp()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(getPercent(percent))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

func benchmarkSetUpHasMutex(b *testing.B, percent int) {
	var set = SetUp()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(getPercent(percent))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetUpAddMutex(b *testing.B) {
	benchmarkSetUpAddMutex(b, 10)
	benchmarkSetUpAddMutex(b, 50)
	benchmarkSetUpAddMutex(b, 90)
}

func BenchmarkSetUpHasMutex(b *testing.B) {
	benchmarkSetUpHasMutex(b, 10)
	benchmarkSetUpHasMutex(b, 50)
	benchmarkSetUpHasMutex(b, 90)
}
