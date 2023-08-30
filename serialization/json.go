package serialization

import (
	"encoding/json"
	"os"
)

func MustMarshalJson[T any](obj T, indented bool) []byte {
	var data []byte
	var err error
	if indented {
		data, err = json.MarshalIndent(obj, "", "    ")
	} else {
		data, err = json.Marshal(obj)
	}
	if err != nil {
		panic(err)
	}
	return data
}

func MustMarshalJsonFile[T any](obj T, path string) {
	data := MustMarshalJson(obj, true)
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
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

func MarshalJson[T any](obj T, indented bool) ([]byte, error) {
	var data []byte
	var err error
	if indented {
		data, err = json.MarshalIndent(obj, "", "    ")
	} else {
		data, err = json.Marshal(obj)
	}
	return data, err
}

func MarshalJsonFile[T any](obj T, path string) error {
	data, err := MarshalJson(obj, true)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

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
