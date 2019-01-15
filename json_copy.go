package copy

import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JSONCopy(dst, src interface{}) {
	bytes, _ := json.Marshal(src)
	json.Unmarshal(bytes, dst)
}
