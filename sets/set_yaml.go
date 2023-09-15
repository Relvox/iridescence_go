package sets

import "gopkg.in/yaml.v3"

func (s Set[T]) MarshalYAML() (interface{}, error) {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}
	return slice, nil
}

func (s *Set[T]) UnmarshalYAML(value *yaml.Node) error {
	var slice []T
	for _, v := range value.Content {
		var t T
		err := v.Decode(&t)
		if err != nil {
			return err
		}
		slice = append(slice, t)
	}

	*s = make(Set[T])
	for _, elem := range slice {
		s.Add(elem)
	}

	return nil
}
