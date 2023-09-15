package sets

import "encoding/json"

func (s Set[T]) MarshalJSON() ([]byte, error) {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}

	return json.Marshal(slice)
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	*s = make(Set[T])
	for _, elem := range slice {
		s.Add(elem)
	}

	return nil
}
