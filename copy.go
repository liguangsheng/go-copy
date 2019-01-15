package copy

var defaultCopier = NewCopier()

func Copy(dst, src interface{}) error {
	return defaultCopier.Copy(dst, src)
}
