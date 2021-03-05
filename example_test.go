package copy_test

import (
	"fmt"
	"github.com/liguangsheng/go-copy"
	"github.com/modern-go/reflect2"
	"time"
	"unsafe"
)

// Example

func Example() {
	var src = struct {
		Field1 time.Time
		Field2 int64
	}{
		Field1: time.Now(),
		Field2: time.Now().Unix(),
	}
	var dest = struct {
		Field1 int64
		Field2 time.Time
	}{}
	if err := copy.Copy(&dest, src); err != nil {
		fmt.Println(err)
	}
}

// Example of custom handlers.

type IntToString struct{}

func (d *IntToString) Samples() (dest, src interface{}) {
	return "", int(0)
}

func (d *IntToString) Copy(destType, srcType reflect2.Type, destPtr, srcPtr unsafe.Pointer) error {
	val := *(srcType.PackEFace(srcPtr).(*int))
	str := fmt.Sprintf("%d", val)
	destType.UnsafeSet(destPtr, reflect2.PtrOf(str))
	return nil
}

func ExampleHandler() {
	var src int = 42
	var dest string
	copier := copy.New()
	copier.Register(&IntToString{})
	if err := copier.Copy(&dest, src); err != nil {
		fmt.Println(err)
	}
}

// Example of name function
func ExampleNameFunc() {
	var src = struct {
		Field1 int `copy:"copy_dest_field2"`
		Field2 int `copy:"copy_dest_field1"`
	}{
		Field1: 111,
		Field2: 222,
	}
	var dest struct {
		Field1 int `copy:"copy_dest_field1"`
		Field2 int `copy:"copy_dest_field2"`
	}
	copier := copy.New(copy.WithCacheSize(10000), copy.WithNameFunc(copy.NameByCopyTag))
	if err := copier.Copy(&dest, src); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(src)  // {111 222}
		fmt.Println(dest) // {222 111}
	}
}
