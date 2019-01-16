# go-copy

go-copy provides object to object copy for golang.

# example
```go
package main

import (
    "fmt"
    "time"

    "github.com/liguangsheng/go-copy"
)

func main() {
    var src = time.Now()
    var dst int64 = 0
    cpr := copy.NewCopier()
    cpr.Register(copy.TypeInt64.RType(), copy.TypeTime.RType(), &copy.TimeToInt64Descriptor{})
    if err := cpr.Copy(&dst, src); err != nil {
        fmt.Println(err)
    }
    fmt.Println(dst) // 1547545588
}
```

# benchmark
```
➜  go-copy git:(master) ✗ go test -bench=. -benchtime=3s -benchmem
goos: darwin
goarch: amd64
pkg: github.com/liguangsheng/go-copy
BenchmarkJinzhuCopySmall-4    	 2000000	      3506 ns/op	    1616 B/op	      21 allocs/op
BenchmarkDeepCopySmall-4      	 1000000	      5497 ns/op	    2520 B/op	      49 allocs/op
BenchmarkJSONCopySmall-4      	 5000000	      1228 ns/op	     104 B/op	       5 allocs/op
BenchmarkCopySmall-4          	 2000000	      1816 ns/op	     224 B/op	      13 allocs/op
BenchmarkJinzhuCopyMedium-4   	  500000	     11583 ns/op	    7320 B/op	      58 allocs/op
BenchmarkDeepCopyMedium-4     	  100000	     37363 ns/op	   21672 B/op	     331 allocs/op
BenchmarkJSONCopyMedium-4     	 1000000	      4098 ns/op	     472 B/op	      17 allocs/op
BenchmarkCopyMedium-4         	 1000000	      3865 ns/op	     624 B/op	      34 allocs/op
BenchmarkJinzhuCopyBig-4      	   30000	    137160 ns/op	   68848 B/op	     487 allocs/op
BenchmarkDeepCopyBig-4        	    2000	   2409910 ns/op	 1907778 B/op	   25759 allocs/op
BenchmarkJSONCopyBig-4        	  100000	     35770 ns/op	    4759 B/op	     152 allocs/op
BenchmarkCopyBig-4            	  200000	     32842 ns/op	    5745 B/op	     304 allocs/op
PASS
ok  	github.com/liguangsheng/go-copy	71.753s
```

# TODO
- parse struct field with tag
- more descriptor
- more unit test
- support slice
- use lru cache instead map cache
- optimize to improve speed