package values

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
)

func ptr[A any](a A) *A { return &a }

func valueStrData(arr arrow.Array) []*string {
	res := make([]*string, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = ptr(arr.ValueStr(i))
		}
	}
	return res
}

func float16Value(arr *array.Float16) []*float32 {
	res := make([]*float32, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = ptr(arr.Value(i).Float32())
		}
	}
	return res
}

func byteArrValue(arr primitive[[]byte]) [][]byte {
	res := make([][]byte, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = arr.Value(i)
		}
	}
	return res
}
