package serialization

import (
	"encoding/json"
	"io/fs"
	"os"
)

func UnmarshalJson[T any](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	return t, err
}

func UnmarshalJsonFile[T any](path string) (T, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		var t T
		return t, err
	}
	return UnmarshalJson[T](bs)
}

func UnmarshalJsonFS[T any](f fs.FS, path string) (T, error) {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		var t T
		return t, err
	}
	return UnmarshalJson[T](bs)
}

func MustUnmarshalJson[T any](data []byte) T {
	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}
	return t
}

func MustUnmarshalJsonFile[T any](path string) T {
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return MustUnmarshalJson[T](bs)
}

func MustUnmarshalJsonFS[T any](f fs.FS, path string) T {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		panic(err)
	}
	return MustUnmarshalJson[T](bs)
}
