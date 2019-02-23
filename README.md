# go-copy

fast copy struct to struct for golang.

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
	"github.com/liguangsheng/go-copy"
	"github.com/modern-go/reflect2"
	"unsafe"
)

type IntToStringCopier struct{}

func (d *IntToStringCopier) Pairs() []copy.TypePair {
	return []copy.TypePair{{
		SrcType: reflect2.TypeOf(int(0)).RType(),
		DstType: reflect2.TypeOf("").RType(),
	}}
}

func (d *IntToStringCopier) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	val := *(srcType.PackEFace(srcPtr).(*int))
	str := fmt.Sprintf("%d", val)
	dstType.UnsafeSet(dstPtr, reflect2.PtrOf(str))
}

func main() {
	var src int = 42
	var dst string
	copier := copy.NewCopier()
	copier.Register(&IntToStringCopier{})
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
	copier := copy.NewCopier(copy.WithParseFunc(copy.ParseFieldByCopyTag))
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
λ go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: github.com/liguangsheng/go-copy/_benchmark
BenchmarkJinzhuCopyBig-8           10000            101620 ns/op           68848 B/op        487 allocs/op
BenchmarkDeepCopyBig-8              1000           1848412 ns/op         1907780 B/op      25759 allocs/op
BenchmarkJSONCopyBig-8             50000             27703 ns/op            4758 B/op        152 allocs/op
BenchmarkCopyBig-8                200000             11964 ns/op            2512 B/op        102 allocs/op (this repo)
BenchmarkJinzhuCopyMedium-8       200000              6971 ns/op            7320 B/op         58 allocs/op
BenchmarkDeepCopyMedium-8          50000             26620 ns/op           21672 B/op        331 allocs/op
BenchmarkJSONCopyMedium-8         500000              2817 ns/op             472 B/op         17 allocs/op
BenchmarkCopyMedium-8            1000000              1453 ns/op             272 B/op         12 allocs/op (this repo)
BenchmarkJinzhuCopySmall-8       1000000              2090 ns/op            1616 B/op         21 allocs/op
BenchmarkDeepCopySmall-8          300000              3816 ns/op            2520 B/op         49 allocs/op
BenchmarkJSONCopySmall-8         2000000               814 ns/op             104 B/op          5 allocs/op
BenchmarkCopySmall-8             3000000               586 ns/op              96 B/op          5 allocs/op (this repo)
PASS
ok      github.com/liguangsheng/go-copy/_benchmark      21.368s                                        
```

# TODO
- more descriptor
- more unit test
