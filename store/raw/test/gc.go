package test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/shousper/my-kit/store/raw"
)

func profileMemory(fn func(t *testing.T)) func(t *testing.T) {
	printMemStats := func(t *testing.T, before, after, afterGC *runtime.MemStats) {
		t.Logf("Alloc         : %d (%d)", after.Alloc-before.Alloc, after.Alloc-afterGC.Alloc)
		t.Logf("TotalAlloc    : %d (%d)", after.TotalAlloc-before.TotalAlloc, after.TotalAlloc-afterGC.TotalAlloc)
		t.Logf("Lookups       : %d (%d)", after.Lookups-before.Lookups, after.Lookups-afterGC.Lookups)
		t.Logf("Mallocs       : %d (%d)", after.Mallocs-before.Mallocs, after.Mallocs-afterGC.Mallocs)
		t.Logf("Frees         : %d (%d)", after.Frees-before.Frees, after.Frees-afterGC.Frees)
		t.Logf("HeapAlloc     : %d (%d)", after.HeapAlloc-before.HeapAlloc, after.HeapAlloc-afterGC.HeapAlloc)
		t.Logf("HeapIdle      : %d (%d)", after.HeapIdle-before.HeapIdle, after.HeapIdle-afterGC.HeapIdle)
		t.Logf("HeapInuse     : %d (%d)", after.HeapInuse-before.HeapInuse, after.HeapInuse-afterGC.HeapInuse)
		t.Logf("HeapReleased  : %d (%d)", after.HeapReleased-before.HeapReleased, after.HeapReleased-afterGC.HeapReleased)
		t.Logf("HeapObjects   : %d (%d)", after.HeapObjects-before.HeapObjects, after.HeapObjects-afterGC.HeapObjects)
		t.Logf("PauseTotalNs  : %d (%d)", after.PauseTotalNs-before.PauseTotalNs, after.PauseTotalNs-afterGC.PauseTotalNs)
		t.Logf("NumGC         : %d (%d)", after.NumGC-before.NumGC, after.NumGC-afterGC.NumGC)
		t.Logf("GCCPUFraction : %0.6f (%0.6f)", after.GCCPUFraction-before.GCCPUFraction, after.GCCPUFraction-afterGC.GCCPUFraction)
	}
	return func(t *testing.T) {
		before := runtime.MemStats{}
		after := runtime.MemStats{}
		afterGC := runtime.MemStats{}

		runtime.ReadMemStats(&before)

		fn(t)

		runtime.ReadMemStats(&after)
		runtime.GC()
		runtime.ReadMemStats(&afterGC)

		printMemStats(t, &before, &after, &afterGC)
	}
}

func GCStress(t *testing.T, fn func() raw.Store) {
	var configurations = []struct {
		capacity, size int
	}{
		{1000, 16},
		{1000, 1024},
		{1000, 262144},
		{1000000, 16},
		{1000000, 1024},
	}
	for _, cfg := range configurations {
		t.Run(fmt.Sprintf("Get %d/%d", cfg.capacity, cfg.size), func(t *testing.T) {
			set := fn()
			keys, _ := AddItems(t, set, cfg.capacity, cfg.size)

			var out interface{}
			profileMemory(func(t *testing.T) {
				for i := 0; i < cfg.capacity; i++ {
					out, _ = set.Get(keys[i])
				}
				result = out
			})(t)
		})
	}
	for _, cfg := range configurations {
		t.Run(fmt.Sprintf("Set %d/%d", cfg.capacity, cfg.size), func(t *testing.T) {
			set := fn()
			keys, values := AddItems(t, set, cfg.capacity, cfg.size)

			out := fn()
			profileMemory(func(t *testing.T) {
				for i := 0; i < cfg.capacity; i++ {
					_ = out.Set(keys[i], values[i])
				}
				result = out
			})(t)
		})
	}
	for _, cfg := range configurations {
		t.Run(fmt.Sprintf("Delete %d/%d", cfg.capacity, cfg.size), func(t *testing.T) {
			set := fn()
			keys, _ := AddItems(t, set, cfg.capacity, cfg.size)

			profileMemory(func(t *testing.T) {
				for i := 0; i < cfg.capacity; i++ {
					_ = set.Delete(keys[i])
				}
				result = set
			})(t)
		})
	}
	for _, cfg := range configurations {
		t.Run(fmt.Sprintf("Reset %d/%d", cfg.capacity, cfg.size), func(t *testing.T) {
			sets := make([]raw.Store, cfg.capacity)
			for i := 0; i < cfg.capacity; i++ {
				sets[i] = fn()
			}

			profileMemory(func(t *testing.T) {
				for i := 0; i < cfg.capacity; i++ {
					_ = sets[i].Reset()
				}
				result = sets
			})(t)
		})
	}
	for _, cfg := range configurations {
		t.Run(fmt.Sprintf("Iterate %d/%d", cfg.capacity, cfg.size), func(t *testing.T) {
			set := fn()
			_, _ = AddItems(t, set, cfg.capacity, cfg.size)

			profileMemory(func(t *testing.T) {
				for i := 0; i < cfg.capacity; i++ {
					it := set.Iterate()
					for {
						value, ok := it.Next()
						if !ok {
							break
						}
						result = value
					}
				}
			})(t)
		})
	}
}
