package _benchmark

import (
	jzcopier "github.com/jinzhu/copier"
	"github.com/liguangsheng/go-copy"
	"github.com/ulule/deepcopier"
	"testing"
	"time"
)

type TestStructBigA struct {
	Field0  int
	Field1  int8
	Field2  int16
	Field3  int32
	Field4  int64
	Field5  float32
	Field6  float64
	Field7  string
	Field8  bool
	Field9  time.Time
	Field10 int
	Field11 int8
	Field12 int16
	Field13 int32
	Field14 int64
	Field15 float32
	Field16 float64
	Field17 string
	Field18 bool
	Field19 time.Time
	Field20 int
	Field21 int8
	Field22 int16
	Field23 int32
	Field24 int64
	Field25 float32
	Field26 float64
	Field27 string
	Field28 bool
	Field29 time.Time
	Field30 int
	Field31 int8
	Field32 int16
	Field33 int32
	Field34 int64
	Field35 float32
	Field36 float64
	Field37 string
	Field38 bool
	Field39 time.Time
	Field40 int
	Field41 int8
	Field42 int16
	Field43 int32
	Field44 int64
	Field45 float32
	Field46 float64
	Field47 string
	Field48 bool
	Field49 time.Time
	Field50 int
	Field51 int8
	Field52 int16
	Field53 int32
	Field54 int64
	Field55 float32
	Field56 float64
	Field57 string
	Field58 bool
	Field59 time.Time
	Field60 int
	Field61 int8
	Field62 int16
	Field63 int32
	Field64 int64
	Field65 float32
	Field66 float64
	Field67 string
	Field68 bool
	Field69 time.Time
	Field70 int
	Field71 int8
	Field72 int16
	Field73 int32
	Field74 int64
	Field75 float32
	Field76 float64
	Field77 string
	Field78 bool
	Field79 time.Time
	Field80 int
	Field81 int8
	Field82 int16
	Field83 int32
	Field84 int64
	Field85 float32
	Field86 float64
	Field87 string
	Field88 bool
	Field89 time.Time
	Field90 int
	Field91 int8
	Field92 int16
	Field93 int32
	Field94 int64
	Field95 float32
	Field96 float64
	Field97 string
	Field98 bool
	Field99 time.Time
}
type TestStructBigB struct {
	Field0  int
	Field1  int8
	Field2  int16
	Field3  int32
	Field4  int64
	Field5  float32
	Field6  float64
	Field7  string
	Field8  bool
	Field9  time.Time
	Field10 int
	Field11 int8
	Field12 int16
	Field13 int32
	Field14 int64
	Field15 float32
	Field16 float64
	Field17 string
	Field18 bool
	Field19 time.Time
	Field20 int
	Field21 int8
	Field22 int16
	Field23 int32
	Field24 int64
	Field25 float32
	Field26 float64
	Field27 string
	Field28 bool
	Field29 time.Time
	Field30 int
	Field31 int8
	Field32 int16
	Field33 int32
	Field34 int64
	Field35 float32
	Field36 float64
	Field37 string
	Field38 bool
	Field39 time.Time
	Field40 int
	Field41 int8
	Field42 int16
	Field43 int32
	Field44 int64
	Field45 float32
	Field46 float64
	Field47 string
	Field48 bool
	Field49 time.Time
	Field50 int
	Field51 int8
	Field52 int16
	Field53 int32
	Field54 int64
	Field55 float32
	Field56 float64
	Field57 string
	Field58 bool
	Field59 time.Time
	Field60 int
	Field61 int8
	Field62 int16
	Field63 int32
	Field64 int64
	Field65 float32
	Field66 float64
	Field67 string
	Field68 bool
	Field69 time.Time
	Field70 int
	Field71 int8
	Field72 int16
	Field73 int32
	Field74 int64
	Field75 float32
	Field76 float64
	Field77 string
	Field78 bool
	Field79 time.Time
	Field80 int
	Field81 int8
	Field82 int16
	Field83 int32
	Field84 int64
	Field85 float32
	Field86 float64
	Field87 string
	Field88 bool
	Field89 time.Time
	Field90 int
	Field91 int8
	Field92 int16
	Field93 int32
	Field94 int64
	Field95 float32
	Field96 float64
	Field97 string
	Field98 bool
	Field99 time.Time
}

func BenchmarkJinzhuCopyBig(b *testing.B) {
	var src = TestStructBigA{}
	var dst TestStructBigB
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jzcopier.Copy(&dst, src)
	}
}

func BenchmarkDeepCopyBig(b *testing.B) {
	var src = TestStructBigA{}
	var dst TestStructBigB
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deepcopier.Copy(src).To(&dst)
	}
}

func BenchmarkJSONCopyBig(b *testing.B) {
	var src = TestStructBigA{}
	var dst TestStructBigB
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy.JSONCopy(&dst, src)
	}
}

func BenchmarkCopyBig(b *testing.B) {
	var src = TestStructBigA{}
	var dst TestStructBigB
	cpr := copy.NewCopier()
	cpr.Copy(&dst, src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpr.Copy(&dst, src)
	}
}
