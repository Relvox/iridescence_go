package serialization

import (
	"os"

	"gopkg.in/yaml.v3"
)

func MarshalYaml[T any](obj T) ([]byte, error) {
	data, err := yaml.Marshal(obj)
	return data, err
}

func MarshalYamlFile[T any](obj T, path string) error {
	data, err := MarshalYaml(obj)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

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
