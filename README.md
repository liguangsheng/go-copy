[![Go Report Card](https://goreportcard.com/badge/github.com/liguangsheng/go-copy)](https://goreportcard.com/report/github.com/liguangsheng/go-copy)

# go-copy

fast copy struct dest struct for golang.

# example

see [example_test.go](example_test.go)

# benchmark
```
goos: windows
goarch: amd64
pkg: github.com/liguangsheng/go-copy/benchmark
BenchmarkJinzhuCopyBig-8           10000            110397 ns/op           38944 B/op        675 allocs/op
BenchmarkDeepCopyBig-8               727           1635491 ns/op         1907809 B/op      25759 allocs/op
BenchmarkJSONCopyBig-8             52401             22996 ns/op            4439 B/op        142 allocs/op
BenchmarkMyCopyBig-8             1490727               798 ns/op             896 B/op          1 allocs/op
BenchmarkJinzhuCopyMedium-8       196717              5861 ns/op            4136 B/op         72 allocs/op
BenchmarkDeepCopyMedium-8          53563             22665 ns/op           21672 B/op        331 allocs/op
BenchmarkJSONCopyMedium-8         521202              2381 ns/op             440 B/op         16 allocs/op
BenchmarkMyCopyMedium-8          4255312               284 ns/op              96 B/op          1 allocs/op
BenchmarkJinzhuCopySmall-8        665781              1799 ns/op            1200 B/op         25 allocs/op
BenchmarkDeepCopySmall-8          363690              3247 ns/op            2520 B/op         49 allocs/op
BenchmarkJSONCopySmall-8         1999911               627 ns/op             104 B/op          5 allocs/op
BenchmarkMyCopySmall-8           5194210               218 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/liguangsheng/go-copy/benchmark       17.169s                               
```

