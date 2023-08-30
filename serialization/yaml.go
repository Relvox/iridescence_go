package serialization

import (
	"os"

	"gopkg.in/yaml.v3"
)

func MustMarshalYaml[T any](obj T) []byte {
	data, err := yaml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return data
}

func MustMarshalYamlFile[T any](obj T, path string) {
	data := MustMarshalYaml(obj)
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
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

func MarshalYaml[T any](obj T, indented bool) ([]byte, error) {
	data, err := yaml.Marshal(obj)
	return data, err
}

func MarshalYamlFile[T any](obj T, path string) error {
	data, err := MarshalYaml(obj, true)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

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
