package local_test

import (
	"testing"
	"time"

	"github.com/shousper/my-kit/store/raw"
	"github.com/shousper/my-kit/store/raw/local"
	"github.com/shousper/my-kit/store/raw/test"
)

/*
BenchmarkNewBigCacheStore/Get_1000/16-4                 	 5000000	       341 ns/op	      80 B/op	       3 allocs/op
BenchmarkNewBigCacheStore/Get_1000/1024-4               	 1000000	      1036 ns/op	    1088 B/op	       3 allocs/op
BenchmarkNewBigCacheStore/Get_1000/262144-4             	   10000	    108530 ns/op	  262208 B/op	       3 allocs/op
BenchmarkNewBigCacheStore/Get_1000000/16-4              	 3000000	       498 ns/op	      68 B/op	       2 allocs/op
BenchmarkNewBigCacheStore/Get_1000000/1024-4            	 1000000	      1046 ns/op	     478 B/op	       1 allocs/op

BenchmarkNewBigCacheStore/Set_1000/16-4                 	 3000000	       451 ns/op	      83 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Set_1000/1024-4               	 1000000	      2372 ns/op	    2059 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Set_1000/262144-4             	    5000	    502256 ns/op	  647604 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Set_1000000/16-4              	 2000000	       641 ns/op	      17 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Set_1000000/1024-4            	 1000000	      2595 ns/op	    1577 B/op	       0 allocs/op

BenchmarkNewBigCacheStore/Delete_1000/16-4              	20000000	        95.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Delete_1000/1024-4            	20000000	        96.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Delete_1000/262144-4          	20000000	        97.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Delete_1000000/16-4           	10000000	       127 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewBigCacheStore/Delete_1000000/1024-4         	10000000	       204 ns/op	       0 B/op	       0 allocs/op

BenchmarkNewBigCacheStore/Reset_1000/16-4               	     300	   5903314 ns/op	17448960 B/op	    4096 allocs/op
BenchmarkNewBigCacheStore/Reset_1000/1024-4             	     100	  21488047 ns/op	17448960 B/op	    4096 allocs/op
BenchmarkNewBigCacheStore/Reset_1000/262144-4           	     100	  18338171 ns/op	17448960 B/op	    4096 allocs/op
BenchmarkNewBigCacheStore/Reset_1000000/16-4            	     100	  19721385 ns/op	17448960 B/op	    4096 allocs/op
BenchmarkNewBigCacheStore/Reset_1000000/1024-4          	     100	  18561838 ns/op	17448960 B/op	    4096 allocs/op

BenchmarkNewBigCacheStore/Iterate_1000/16-4             	 1000000	      2019 ns/op	     164 B/op	       5 allocs/op
BenchmarkNewBigCacheStore/Iterate_1000/1024-4           	  500000	      2176 ns/op	    1184 B/op	       5 allocs/op
BenchmarkNewBigCacheStore/Iterate_1000/262144-4         	   20000	     65395 ns/op	  262292 B/op	       5 allocs/op
BenchmarkNewBigCacheStore/Iterate_1000000/16-4          	  100000	     13181 ns/op	    3232 B/op	       5 allocs/op
BenchmarkNewBigCacheStore/Iterate_1000000/1024-4        	  100000	     11981 ns/op	    3472 B/op	       5 allocs/op
*/
func BenchmarkNewBigCacheStore(b *testing.B) {
	//b.SkipNow()

	test.Benchmark(b, func() raw.Store { return local.NewBigCacheStore(1 * time.Second) })
}

func TestNewBigCacheStore(t *testing.T) {
	test.GCStress(t, func() raw.Store { return local.NewBigCacheStore(1 * time.Second) })
}
