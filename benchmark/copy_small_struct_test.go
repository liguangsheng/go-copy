package benchmark

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
	var dest TestStructSmallB
	for i := 0; i < b.N; i++ {
		jzcopier.Copy(&dest, src)
	}
}

func BenchmarkDeepCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dest TestStructSmallB
	for i := 0; i < b.N; i++ {
		deepcopier.Copy(src).To(&dest)
	}
}

func BenchmarkJSONCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dest TestStructSmallB
	for i := 0; i < b.N; i++ {
		JSONCopy(&dest, src)
	}
}

func BenchmarkMyCopySmall(b *testing.B) {
	var src = TestStructSmallA{}
	var dest TestStructSmallB
	cpr := copy.New()
	cpr.Copy(&dest, src)
	for i := 0; i < b.N; i++ {
		cpr.Copy(&dest, src)
	}
}
