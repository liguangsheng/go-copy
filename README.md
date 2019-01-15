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

# TODO
- parse struct field with tag
- more descriptor
- more unit test
- support slice
- use lru cache instead map cache
- optimize to improve speed