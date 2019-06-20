package raw_test

import (
	"math/rand"
	"testing"

	"my-kit/store/raw"
	"my-kit/store/raw/test"
)

/*
BenchmarkNewLazyStore/Get_1000/16-4                     	10000000	       166 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Set_1000/16-4                     	10000000	       126 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Delete_1000/16-4                  	20000000	        91.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Reset_1000/16-4                   	10000000	       339 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewLazyStore/Iterate_1000/16-4                 	10000000	       307 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Get_1000/1024-4                   	 2000000	       514 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Set_1000/1024-4                   	10000000	       238 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Delete_1000/1024-4                	10000000	       127 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Reset_1000/1024-4                 	10000000	       324 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewLazyStore/Iterate_1000/1024-4               	10000000	       166 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Get_1000/262144-4                 	10000000	       169 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Set_1000/262144-4                 	10000000	       132 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Delete_1000/262144-4              	20000000	        93.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Reset_1000/262144-4               	10000000	       269 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewLazyStore/Iterate_1000/262144-4             	10000000	       247 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Get_1000000/16-4                  	 5000000	       423 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Set_1000000/16-4                  	 1000000	      1528 ns/op	     200 B/op	       0 allocs/op
BenchmarkNewLazyStore/Delete_1000000/16-4               	10000000	       147 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Reset_1000000/16-4                	10000000	       202 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewLazyStore/Iterate_1000000/16-4              	 5000000	       367 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Get_1000000/1024-4                	 5000000	       359 ns/op	      32 B/op	       1 allocs/op
BenchmarkNewLazyStore/Set_1000000/1024-4                	 2000000	       684 ns/op	     100 B/op	       0 allocs/op
BenchmarkNewLazyStore/Delete_1000000/1024-4             	10000000	       107 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewLazyStore/Reset_1000000/1024-4              	10000000	       166 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewLazyStore/Iterate_1000000/1024-4            	 5000000	       676 ns/op	      32 B/op	       1 allocs/op
*/
func BenchmarkNewLazyStore(b *testing.B) {
	//b.SkipNow()

	data := make([][]byte, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = make([]byte, 16)
		rand.Read(data[i])
	}

	i := -1
	test.Benchmark(b, func() raw.Store {
		return raw.NewLazyStore(raw.NewDefaultStore(), func(key string) (bytes []byte, e error) {
			i++
			return data[i], nil
		})
	})
}

func TestNewLazyStore(t *testing.T) {
	data := make([][]byte, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = make([]byte, 16)
		rand.Read(data[i])
	}

	i := -1
	test.GCStress(t, func() raw.Store {
		return raw.NewLazyStore(raw.NewDefaultStore(), func(key string) (bytes []byte, e error) {
			i++
			return data[i], nil
		})
	})
}