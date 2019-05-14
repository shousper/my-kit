package local_test

import (
	"testing"

	"github.com/shousper/my-kit/store/raw"
	"github.com/shousper/my-kit/store/raw/local"
	"github.com/shousper/my-kit/store/raw/test"
)

/*
BenchmarkNewConcurrentMapStore/Get_1000/16-4         	10000000	       221 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Get_1000/1024-4       	10000000	       170 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Get_1000/262144-4     	10000000	       169 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Get_1000000/16-4      	 5000000	       302 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Get_1000000/1024-4    	 5000000	       289 ns/op	      32 B/op	       1 allocs/op

BenchmarkNewConcurrentMapStore/Set_1000/16-4         	10000000	       177 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Set_1000/1024-4       	10000000	       178 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Set_1000/262144-4     	10000000	       176 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Set_1000000/16-4      	 3000000	       461 ns/op	      86 B/op	       1 allocs/op
BenchmarkNewConcurrentMapStore/Set_1000000/1024-4    	 3000000	       385 ns/op	      86 B/op	       1 allocs/op

BenchmarkNewConcurrentMapStore/Delete_1000/16-4      	20000000	        91.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewConcurrentMapStore/Delete_1000/1024-4    	20000000	        91.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewConcurrentMapStore/Delete_1000/262144-4  	20000000	        91.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewConcurrentMapStore/Delete_1000000/16-4   	10000000	       107 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewConcurrentMapStore/Delete_1000000/1024-4 	10000000	       107 ns/op	       0 B/op	       0 allocs/op

BenchmarkNewConcurrentMapStore/Reset_1000/16-4       	  300000	      5701 ns/op	    2816 B/op	      65 allocs/op
BenchmarkNewConcurrentMapStore/Reset_1000/1024-4     	  300000	      3862 ns/op	    2816 B/op	      65 allocs/op
BenchmarkNewConcurrentMapStore/Reset_1000/262144-4   	  300000	      4284 ns/op	    2816 B/op	      65 allocs/op
BenchmarkNewConcurrentMapStore/Reset_1000000/16-4    	  300000	      3826 ns/op	    2816 B/op	      65 allocs/op
BenchmarkNewConcurrentMapStore/Reset_1000000/1024-4  	  300000	      4299 ns/op	    2816 B/op	      65 allocs/op

BenchmarkNewConcurrentMapStore/Iterate_1000/16-4     	   20000	     70488 ns/op	   32000 B/op	    1000 allocs/op
BenchmarkNewConcurrentMapStore/Iterate_1000/1024-4   	   20000	     73985 ns/op	   32000 B/op	    1000 allocs/op
BenchmarkNewConcurrentMapStore/Iterate_1000/262144-4 	   20000	     72363 ns/op	   32000 B/op	    1000 allocs/op
BenchmarkNewConcurrentMapStore/Iterate_1000000/16-4  	      10	 229161458 ns/op	32000000 B/op	 1000000 allocs/op
BenchmarkNewConcurrentMapStore/Iterate_1000000/1024-4         	      10	 170099232 ns/op	32000000 B/op	 1000000 allocs/op
*/
func BenchmarkNewConcurrentMapStore(b *testing.B) {
	//b.SkipNow()

	test.Benchmark(b, func() raw.Store { return local.NewConcurrentMapStore() })
}

func TestNewConcurrentMapStore(t *testing.T) {
	test.GCStress(t, func() raw.Store { return local.NewConcurrentMapStore() })
}
