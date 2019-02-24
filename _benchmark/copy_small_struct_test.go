package _benchmark

import (
	"testing"

	jzcopier "github.com/jinzhu/copier"
	"github.com/liguangsheng/go-copy"
	"github.com/ulule/deepcopier"
)

type TestStructSmallA struct {
	Field1 int
	Field7 float64
	Field8 string
}
type TestStructSmallB struct {
	Field1 int
	Field7 float64
	Field8 string
}

func BenchmarkJinzhuCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dst TestStructSmallB
	for i := 0; i < b.N; i++ {
		jzcopier.Copy(&dst, src)
	}
}

func BenchmarkDeepCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dst TestStructSmallB
	for i := 0; i < b.N; i++ {
		deepcopier.Copy(src).To(&dst)
	}
}

func BenchmarkJSONCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dst TestStructSmallB
	for i := 0; i < b.N; i++ {
		copy.JSONCopy(&dst, src)
	}
}

func BenchmarkCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dst TestStructSmallB
	cpr := copy.NewCopier()
	cpr.Copy(&dst, src)
	for i := 0; i < b.N; i++ {
		cpr.Copy(&dst, src)
	}
}
