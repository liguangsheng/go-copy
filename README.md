# go-copy

> **This project is no longer actively maintained.** With Go generics now mature, the recommended approaches for struct conversion are:
> 1. **Hand-written conversion functions** — compile-time type safety, zero runtime overhead, refactor-friendly
> 2. **Code generation** (e.g. [convergen](https://github.com/reedom/convergen)) — auto-generated conversion code, balancing convenience and performance
>
> Reflection/unsafe-based copying has no compile-time protection when fields change, making it easy to introduce silent bugs during refactoring. Unless you have a highly dynamic use case (e.g. ORM internals), this approach is not recommended.

A high-performance Go struct copying library. Uses `unsafe` pointer operations and contiguous memory merging to achieve fast data copying between structs.

## How It Works

Core path: `Copier.Copy(dest, src)` →

1. **Type analysis** — recursively expands embedded structs, extracts field names via `NameFunc`
2. **Field matching** — builds src→dest `fieldMapping` pairs by name
3. **Contiguous merging** — adjacent fields with contiguous memory layout are merged into a single `copy` (memcpy)
4. **Caching** — resolved type mappings are cached in `sync.Map`, subsequent calls hit the cache directly
5. **Execution** — operates on memory directly via `unsafe.Pointer`, bypassing standard reflection overhead

Key optimizations:
- Merges multiple contiguous assignable fields into a single memory copy
- Automatic numeric type conversion (with sign extension)
- Built-in `time.Time ↔ int64` handler, extensible via custom `TypedHandler`

## Usage

```go
import copy "github.com/liguangsheng/go-copy"

type User struct {
    Name string
    Age  int
}

type UserDTO struct {
    Name string
    Age  int
}

// Global convenience function
src := User{Name: "Alice", Age: 30}
var dest UserDTO
copy.Copy(&dest, &src)

// Reuse a Copier instance (recommended, avoids repeated type resolution)
cpr := copy.New()
cpr.Copy(&dest, &src)
```

### Field Matching Strategies

```go
// Default: match by field name
cpr := copy.New()

// Match by json tag
cpr := copy.New(copy.WithNameFunc(copy.NameByJSONTag))

// Match by copy tag
cpr := copy.New(copy.WithNameFunc(copy.NameByCopyTag))
```

## Benchmark

Environment: Linux amd64, Intel i5-13600KF

```
goos: linux
goarch: amd64
pkg: github.com/liguangsheng/go-copy/benchmark
cpu: 13th Gen Intel(R) Core(TM) i5-13600KF

BenchmarkCopy/jinzhu/small-4              1000000        1173 ns/op      1024 B/op       6 allocs/op
BenchmarkCopy/deepcopier/small-4           487996        2253 ns/op      2368 B/op      28 allocs/op
BenchmarkCopy/json/small-4                1641702         740 ns/op       264 B/op       5 allocs/op
BenchmarkCopy/go-copy/small-4            40390062          30 ns/op         0 B/op       0 allocs/op

BenchmarkCopy/jinzhu/medium-4              293874        3959 ns/op      3544 B/op      11 allocs/op
BenchmarkCopy/deepcopier/medium-4           66758       18172 ns/op     20784 B/op     226 allocs/op
BenchmarkCopy/json/medium-4               626685        2057 ns/op       408 B/op       6 allocs/op
BenchmarkCopy/go-copy/medium-4           38516486          30 ns/op         0 B/op       0 allocs/op

BenchmarkCopy/jinzhu/big-4                 13569       89225 ns/op     33088 B/op      74 allocs/op
BenchmarkCopy/deepcopier/big-4               968     1303780 ns/op   1862921 B/op   20209 allocs/op
BenchmarkCopy/json/big-4                   71494       16368 ns/op      2235 B/op      15 allocs/op
BenchmarkCopy/go-copy/big-4             30907712          33 ns/op         0 B/op       0 allocs/op
```

| Library | Small (3 fields) | Medium (10 fields) | Big (100 fields) |
|---------|-----------------|-------------------|-----------------|
| json (encoding/json) | ~25x slower | ~68x slower | ~491x slower |
| jinzhu/copier | ~39x slower | ~131x slower | ~2,675x slower |
| ulule/deepcopier | ~75x slower | ~601x slower | ~39,091x slower |

go-copy maintains ~30 ns/op with zero allocations across all sizes. The advantage grows with struct size.

## License

MIT
