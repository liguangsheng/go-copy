package copy

var _globalCopier = New()

// Copy values
func Copy(dest, src interface{}) error {
	return _globalCopier.Copy(dest, src)
}
