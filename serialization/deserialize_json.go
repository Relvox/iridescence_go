package serialization

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/toolvox/utilgo/pkg/errs"

	"github.com/relvox/iridescence_go/errors"
)

func UnmarshalJson[T any](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		return *new(T), fmt.Errorf("unmarshal json: %w", err)
	}
	if v, ok := any(t).(errs.Validator); ok {
		return errors.WrapCommaError(t, v.Validate())("unmarshal json: validate")
	}
	return t, nil
}

func UnmarshalJsonFile[T any](path string) (T, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return *new(T), fmt.Errorf("unmarshal json file: read file '%s': %w", path, err)
	}
	return errors.WrapCommaError(UnmarshalJson[T](bs))("unmarshal json file '%s'", path)
}

func UnmarshalJsonFS[T any](f fs.FS, path string) (T, error) {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		return *new(T), fmt.Errorf("unmarshal json fs file: read fs file '%s': %w", path, err)
	}
	return errors.WrapCommaError(UnmarshalJson[T](bs))("unmarshal json` fs file '%s'", path)
}

func MustUnmarshalJson[T any](data []byte) T {
	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		panic(fmt.Errorf("must unmarshal json: %w", err))
	}
	return t
}

func MustUnmarshalJsonFile[T any](path string) T {
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("unmarshal json file: read file '%s': %w", path, err))
	}
	return MustUnmarshalJson[T](bs)
}

func MustUnmarshalJsonFS[T any](f fs.FS, path string) T {
	bs, err := fs.ReadFile(f, path)
	if err != nil {
		panic(fmt.Errorf("unmarshal json fs file: read fs file '%s': %w", path, err))
	}
	return MustUnmarshalJson[T](bs)
}
