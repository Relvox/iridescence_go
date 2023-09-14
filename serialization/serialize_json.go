package serialization

import (
	"encoding/json"
	"os"
)

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
