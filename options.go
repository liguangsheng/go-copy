package copy

type Option func(cpr *copier)

func WithCacheSize(size int) Option {
	return func(cpr *copier) {
		cpr.cacheSize = size
	}
}

func WithParseFunc(fn FieldParseFunc) Option {
	return func(cpr *copier) {
		cpr.fieldParser = fn
	}
}
