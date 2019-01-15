package copy

import (
	"time"

	"github.com/modern-go/reflect2"
)

var (
	TypeTime  = reflect2.TypeOf(time.Time{})
	TypeInt64 = reflect2.TypeOf(int64(0))
	TypeInt   = reflect2.TypeOf(int(0))
)
