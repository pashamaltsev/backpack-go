package utils

import (
	"unsafe"

	"github.com/go-json-experiment/json"
)

func BytesToString(bytes []byte) string {
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

func StringToBytes(string_ string) (bytes []byte) {
	return unsafe.Slice(unsafe.StringData(string_), len(string_))
}

func Ptr[T any](v T) *T {
	return &v
}

func StructToMap[M1 ~map[K]V, K comparable, V any](obj any) (result map[string]V) {
	raw, _ := json.Marshal(obj)
	json.Unmarshal(raw, &result)
	return
}
