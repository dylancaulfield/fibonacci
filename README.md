$ go test -bench=.

goos: windows  
goarch: amd64  
pkg: github.com/dyldawg/fibonacci  

BenchmarkFibRecursive 201 5840000 ns/op  
BenchmarkFibRecursiveMemo 51910971 23.1 ns/op  
BenchmarkFibInOrder 2723212 433 ns/op  

PASS
ok      github.com/dyldawg/fibonacci  4.669s