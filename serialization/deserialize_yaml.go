package serialization

import (
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

func UnmarshalYaml[T any](data []byte) (T, error) {
	var t T
	err := yaml.Unmarshal(data, &t)
	return t, err
}

func UnmarshalYamlFile[T any](path string) (T, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		var t T
		return t, err
	}
	return UnmarshalYaml[T](bs)
}

func UnmarshalYamlFS[T any](f fs.FS, path string) (T, error) {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		var t T
		return t, err
	}
	return UnmarshalYaml[T](bs)
}

func MustUnmarshalYaml[T any](data []byte) T {
	var t T
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}
	return t
}

func MustUnmarshalYamlFile[T any](path string) T {
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return MustUnmarshalYaml[T](bs)
}

func MustUnmarshalYamlFS[T any](f fs.FS, path string) T {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		panic(err)
	}
	return MustUnmarshalYaml[T](bs)
}
