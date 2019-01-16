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
	var src = struct {
		Field1 time.Time
		Field2 int64
	}{
		Field1: time.Now(),
		Field2: time.Now().Unix(),
	}
	var dst struct {
		Field1 int64
		Field2 time.Time
	}
	if err := copy.Copy(&dst, src); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(src) // {2019-01-16 17:18:06.646149 +0800 CST m=+0.001120925 1547630286}
		fmt.Println(dst) // {1547630286 2019-01-16 17:18:06 +0800 CST}
	}
}
```

You can add your custom descriptor to support more type.
```go
package main

import (
	"fmt"
	"unsafe"

	"github.com/modern-go/reflect2"

	"github.com/liguangsheng/go-copy"
)

type IntToStringDescriptor struct{}

func (d IntToStringDescriptor) SrcType() reflect2.Type {
	return reflect2.TypeOf(int(0))
}

func (d IntToStringDescriptor) DstType() reflect2.Type {
	return reflect2.TypeOf("hahaha")
}

func (d IntToStringDescriptor) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	val := *(srcType.PackEFace(srcPtr).(*int))
	str := fmt.Sprintf("%d", val)
	dstType.UnsafeSet(dstPtr, reflect2.PtrOf(str))
}

func main() {
	var src int = 42
	var dst string
	copier := copy.NewCopier()
	copier.Register(IntToStringDescriptor{})
	if err := copier.Copy(&dst, src); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(src) // 42
		fmt.Println(dst) // 42
	}
}
```

Use different field parser
```go
package main

import (
	"fmt"

	"github.com/liguangsheng/go-copy"
)

func main() {
	var src = struct {
		Field1 int `copy:"copy_to_field2"`
		Field2 int `copy:"copy_to_field1"`
	}{
		Field1: 111,
		Field2: 222,
	}
	var dst struct {
		Field1 int `copy:"copy_to_field1"`
		Field2 int `copy:"copy_to_field2"`
	}
	copier := copy.NewCopier()
	copier.UseFieldParser(copy.ParseFieldByCopyTag)
	if err := copier.Copy(&dst, src); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(src) // {111 222}
		fmt.Println(dst) // {222 111}
	}
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