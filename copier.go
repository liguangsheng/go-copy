package copy

import (
	"reflect"
	"unsafe"

	"github.com/golang/groupcache/lru"
	"github.com/modern-go/reflect2"
)

// New a *Copier
func New(opts ...Option) *Copier {
	c := &Copier{
		cache:    lru.New(1000),
		nameFunc: NameByFieldName,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Register(
		timeToInt64{},
		int64ToTime{},
	)

	return c
}

// Copier dest copy values
type Copier struct {
	cap      int
	cache    *lru.Cache
	nameFunc NameFunc
}

// Register add typed Copier dest cache
func (c *Copier) Register(handlers ...CustomHandler) {
	for _, h := range handlers {
		dest, src := h.Samples()
		hash := hashType(dest, src)
		c.cache.Add(hash, h)
	}
}

// Unregister remove typed Copier src cache
func (c *Copier) Unregister(handlers ...Samples) {
	for _, h := range handlers {
		dest, src := h.Samples()
		hash := hashType(dest, src)
		c.cache.Remove(hash)
	}
}

// Copy values
func (c *Copier) Copy(dest, src interface{}) error {
	var (
		srcType   = indirectType(reflect.TypeOf(src))
		srcType2  = reflect2.Type2(srcType)
		srcPtr    = reflect2.PtrOf(src)
		destType  = indirectType(reflect.TypeOf(dest))
		destType2 = reflect2.Type2(destType)
		destPtr   = reflect2.PtrOf(dest)
	)

	return c.copy(destPtr, srcPtr, &_pair{
		assignable: assignable(destType2, srcType2),
		srcType:    srcType2,
		destType:   destType2,
		destSize:   destType.Size(),
		srcSize:    srcType.Size(),
	})
}

func (c *Copier) copy(dest, src unsafe.Pointer, pair *_pair) error {
	// memory copy
	if pair.assignable {
		pair.copy(dest, src)
		return nil
	}

	// use cached copy handler
	hash := hashRType(pair.destType.RType(), pair.srcType.RType())
	if handler, ok := c.cache.Get(hash); ok {
		return handler.(Handler).Copy(pair.destType, pair.srcType, dest, src)
	}

	fields := parseStructs(pair.destType, pair.srcType, c.nameFunc)
	if fields != nil {
		sh := &structsHandler{copier: c, fields: fields}
		c.cache.Add(hash, sh)
		return sh.Copy(pair.destType, pair.srcType, dest, src)
	}

	return nil
}
