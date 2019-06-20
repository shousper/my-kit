package raw_test

import (
	"testing"

	"my-kit/store/raw"
	"my-kit/store/raw/test"
)

/*
BenchmarkNewDefaultStore/Get_1000/16-4                  	10000000	       287 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Get_1000/1024-4                	 5000000	       220 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Get_1000/262144-4              	10000000	       148 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Get_1000000/16-4               	 5000000	       315 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Get_1000000/1024-4             	 5000000	       292 ns/op	      32 B/op	       1 allocs/op

BenchmarkNewDefaultStore/Set_1000/16-4                  	10000000	       116 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Set_1000/1024-4                	10000000	       154 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Set_1000/262144-4              	10000000	       113 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Set_1000000/16-4               	 1000000	      1087 ns/op	     200 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Set_1000000/1024-4             	 3000000	       375 ns/op	      66 B/op	       0 allocs/op

BenchmarkNewDefaultStore/Delete_1000/16-4               	20000000	        90.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Delete_1000/1024-4             	10000000	       107 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Delete_1000/262144-4           	20000000	        76.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Delete_1000000/16-4            	20000000	        84.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewDefaultStore/Delete_1000000/1024-4          	10000000	       169 ns/op	       0 B/op	       0 allocs/op

BenchmarkNewDefaultStore/Reset_1000/16-4                	10000000	       185 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Reset_1000/1024-4              	10000000	       182 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Reset_1000/262144-4            	10000000	       252 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Reset_1000000/16-4             	10000000	       180 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Reset_1000000/1024-4           	10000000	       271 ns/op	      48 B/op	       1 allocs/op

BenchmarkNewDefaultStore/Iterate_1000/16-4              	10000000	       244 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Iterate_1000/1024-4            	10000000	       156 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Iterate_1000/262144-4          	10000000	       171 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Iterate_1000000/16-4           	 5000000	       265 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewDefaultStore/Iterate_1000000/1024-4         	 5000000	       260 ns/op	      32 B/op	       1 allocs/op
*/
func BenchmarkNewDefaultStore(b *testing.B) {
	//b.SkipNow()

	test.Benchmark(b, func() raw.Store { return raw.NewDefaultStore() })
}

func TestNewDefaultStore(t *testing.T) {
	test.GCStress(t, func() raw.Store { return raw.NewDefaultStore() })
}