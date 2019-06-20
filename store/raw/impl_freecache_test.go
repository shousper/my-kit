package raw_test

import (
	"testing"
	"time"

	"my-kit/store/raw"
	"my-kit/store/raw/test"
)

/*
BenchmarkNewFreeCacheStore/Get_1000/16-4         	10000000	       116 ns/op	       4 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Get_1000/1024-4       	 3000000	       406 ns/op	     157 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Get_1000/262144-4     	10000000	       164 ns/op	     357 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Get_1000000/16-4      	 5000000	       220 ns/op	       0 B/op	       0 allocs/op

BenchmarkNewFreeCacheStore/Set_1000/16-4         	 5000000	       271 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Set_1000/1024-4       	 3000000	       470 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Set_1000/262144-4     	   30000	     41447 ns/op	       1 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Set_1000000/16-4      	 3000000	       600 ns/op	      22 B/op	       0 allocs/op

BenchmarkNewFreeCacheStore/Delete_1000/16-4      	20000000	        68.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Delete_1000/1024-4    	20000000	        67.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Delete_1000/262144-4  	20000000	        73.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFreeCacheStore/Delete_1000000/16-4   	10000000	       100 ns/op	       0 B/op	       0 allocs/op

BenchmarkNewFreeCacheStore/Reset_1000/16-4       	      50	 592913780 ns/op	2148532224 B/op	     512 allocs/op
BenchmarkNewFreeCacheStore/Reset_1000/1024-4     	       5	1866044417 ns/op	2148532230 B/op	     512 allocs/op
BenchmarkNewFreeCacheStore/Reset_1000/262144-4   	       5	1900542113 ns/op	2148532230 B/op	     512 allocs/op
BenchmarkNewFreeCacheStore/Reset_1000000/16-4    	       5	1904515339 ns/op	2148532230 B/op	     512 allocs/op

BenchmarkNewFreeCacheStore/Iterate_1000/16-4     	 1000000	    281166 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewFreeCacheStore/Iterate_1000/1024-4   	 1000000	    104178 ns/op	     953 B/op	       4 allocs/op
BenchmarkNewFreeCacheStore/Iterate_1000/262144-4 	   10000	    319699 ns/op	  134134 B/op	       2 allocs/op
BenchmarkNewFreeCacheStore/Iterate_1000000/16-4            -
*/
func BenchmarkNewFreeCacheStore(b *testing.B) {
	b.SkipNow()

	test.Benchmark(b, func() raw.Store { return raw.NewFreeCacheStore(2 * 1024 * 1024 * 1024, 1 * time.Second) })
}

func TestNewFreeCacheStore(t *testing.T) {
	test.GCStress(t, func() raw.Store { return raw.NewFreeCacheStore(2 * 1024 * 1024 * 1024, 1 * time.Second) })
}