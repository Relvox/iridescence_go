package serialization

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/relvox/iridescence_go/errors"
	"github.com/relvox/iridescence_go/validation"
)

func UnmarshalYaml[T any](data []byte) (T, error) {
	var t T
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		return *new(T), fmt.Errorf("unmarshal yaml: %w", err)
	}
	if v, ok := any(t).(validation.Validable); ok {
		return errors.WrapCommaError(t, v.Validate())("unmarshal yaml: validate")
	}
	return t, nil
}

func UnmarshalYamlFile[T any](path string) (T, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return *new(T), fmt.Errorf("unmarshal yaml file: read file '%s': %w", path, err)
	}
	return errors.WrapCommaError(UnmarshalYaml[T](bs))("unmarshal yaml file '%s'", path)
}

func UnmarshalYamlFS[T any](f fs.FS, path string) (T, error) {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		return *new(T), fmt.Errorf("unmarshal yaml fs file: read fs file '%s': %w", path, err)
	}
	return errors.WrapCommaError(UnmarshalYaml[T](bs))("unmarshal yaml fs file '%s'", path)
}

func MustUnmarshalYaml[T any](data []byte) T {
	var t T
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		panic(fmt.Errorf("must unmarshal yaml: %w", err))
	}
	return t
}

func MustUnmarshalYamlFile[T any](path string) T {
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("must unmarshal yaml file: read file '%s': %w", path, err))
	}
	return MustUnmarshalYaml[T](bs)
}

func MustUnmarshalYamlFS[T any](f fs.FS, path string) T {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		panic(fmt.Errorf("must unmarshal yaml fs file: read fs file '%s': %w", path, err))
	}
	return MustUnmarshalYaml[T](bs)
}
