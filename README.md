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
BenchmarkJinzhuCopySmall-4    	  500000	      2801 ns/op	    1616 B/op	      21 allocs/op
BenchmarkDeepCopySmall-4      	  300000	      5107 ns/op	    2520 B/op	      49 allocs/op
BenchmarkJSONCopySmall-4      	 1000000	      1091 ns/op	     104 B/op	       5 allocs/op
BenchmarkCopySmall-4          	 1000000	      1430 ns/op	     224 B/op	      13 allocs/op
BenchmarkJinzhuCopyMedium-4   	  200000	      9247 ns/op	    7320 B/op	      58 allocs/op
BenchmarkDeepCopyMedium-4     	   50000	     35786 ns/op	   21672 B/op	     331 allocs/op
BenchmarkJSONCopyMedium-4     	  300000	      4008 ns/op	     472 B/op	      17 allocs/op
BenchmarkCopyMedium-4         	  300000	      4050 ns/op	     624 B/op	      34 allocs/op
BenchmarkJinzhuCopyBig-4      	   10000	    150307 ns/op	   68848 B/op	     487 allocs/op
BenchmarkDeepCopyBig-4        	     500	   2503473 ns/op	 1907779 B/op	   25759 allocs/op
BenchmarkJSONCopyBig-4        	   50000	     35980 ns/op	    4759 B/op	     152 allocs/op
BenchmarkCopyBig-4            	   50000	     33829 ns/op	    5745 B/op	     304 allocs/op
```

# TODO
- parse struct field with tag
- more descriptor
- more unit test
- support slice
- use lru cache instead map cache
- optimize to improve speed