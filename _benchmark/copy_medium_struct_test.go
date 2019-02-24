package _benchmark

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
	var dst TestStructMediumB
	for i := 0; i < b.N; i++ {
		jzcopier.Copy(&dst, src)
	}
}

func BenchmarkDeepCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dst TestStructMediumB
	for i := 0; i < b.N; i++ {
		deepcopier.Copy(src).To(&dst)
	}
}

func BenchmarkJSONCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dst TestStructMediumB
	for i := 0; i < b.N; i++ {
		copy.JSONCopy(&dst, src)
	}
}

func BenchmarkCopyMedium(b *testing.B) {
	var src = TestStructMediumA{}
	var dst TestStructMediumB
	cpr := copy.NewCopier()
	cpr.Copy(&dst, src)
	for i := 0; i < b.N; i++ {
		cpr.Copy(&dst, src)
	}
}
