package benchmark

import (
	"testing"
	"time"

	jzcopier "github.com/jinzhu/copier"
	"github.com/liguangsheng/go-copy"
	"github.com/ulule/deepcopier"
)

type TestStructMediumA struct {
	Field1  int
	Field2  int8
	Field3  int16
	Field4  int32
	Field5  int64
	Field6  float32
	Field7  float64
	Field8  string
	Field9  bool
	Field10 time.Time
}
type TestStructMediumB struct {
	Field1  int
	Field2  int8
	Field3  int16
	Field4  int32
	Field5  int64
	Field6  float32
	Field7  float64
	Field8  string
	Field9  bool
	Field10 time.Time
}

func BenchmarkJinzhuCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dest TestStructMediumB
	for i := 0; i < b.N; i++ {
		jzcopier.Copy(&dest, src)
	}
}

func BenchmarkDeepCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dest TestStructMediumB
	for i := 0; i < b.N; i++ {
		deepcopier.Copy(src).To(&dest)
	}
}

func BenchmarkJSONCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dest TestStructMediumB
	for i := 0; i < b.N; i++ {
		jsonCopy(&dest, src)
	}
}

func BenchmarkMyCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dest TestStructMediumB
	cpr := copy.New()
	cpr.Copy(&dest, src)
	for i := 0; i < b.N; i++ {
		cpr.Copy(&dest, src)
	}
}
