package benchmark

import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JSONCopy(dest, src interface{}) error {
	bytes, _ := json.Marshal(src)
	return json.Unmarshal(bytes, dest)
}
