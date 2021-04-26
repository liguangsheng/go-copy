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
BenchmarkJinzhuCopyBig-8                   10000            114140 ns/op           38944 B/op        675 allocs/op
BenchmarkDeepCopyBig-8                       732           1482533 ns/op         1907792 B/op      25759 allocs/op
BenchmarkJSONCopyBig-8                     50223             23857 ns/op            4438 B/op        142 allocs/op
BenchmarkThisRepoCopyBig-8               1486052               807 ns/op             896 B/op          1 allocs/op
BenchmarkJinzhuCopyMedium-8               226514              5669 ns/op            4136 B/op         72 allocs/op
BenchmarkDeepCopyMedium-8                  53980             20994 ns/op           21672 B/op        331 allocs/op
BenchmarkJSONCopyMedium-8                 521784              2466 ns/op             440 B/op         16 allocs/op
BenchmarkThisRepoCopyMedium-8            3804498               297 ns/op              96 B/op          1 allocs/op
BenchmarkJinzhuCopySmall-8                750013              1745 ns/op            1200 B/op         25 allocs/op
BenchmarkDeepCopySmall-8                  406654              3108 ns/op            2520 B/op         49 allocs/op
BenchmarkJSONCopySmall-8                 1889121               613 ns/op             104 B/op          5 allocs/op
BenchmarkThisRepoCopySmall-8             5519410               226 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/liguangsheng/go-copy/benchmark       17.516s
```

