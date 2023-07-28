# iridescence_go

A collection of boilerplate I find myself writing a lot.

## Contents

### Utils

#### `maps.go`

```go
// MergeMaps creates a new map and copies target and then source into it
func MergeMaps[K comparable, V any](target, source map[K]V) map[K]V
// Keys gets a slice of all the keys in a map
func Keys[K comparable, V any](self map[K]V) []K
// SortedKeys gets a slice of all the keys in a map, sorted
func SortedKeys[K constraints.Ordered, V any](self map[K]V) []K {
// Values gets a slice of all the values in a map
func Values[K comparable, V any](self map[K]V) []V
// MapToStruct converts a map to a struct by converting through json
func MapToStruct[T any](m map[string]any) (T, error)
// StructToMap converts a struct to a map by converting through json
func StructToMap[T any](t T) (map[string]interface{}, error)
```
