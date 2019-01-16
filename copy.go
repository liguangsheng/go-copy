package copy

var defaultCopier *copier

func init() {
	defaultCopier = NewCopier()
}

func Copy(dst, src interface{}) error {
	return defaultCopier.Copy(dst, src)
}
