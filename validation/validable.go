package validation

type Validable interface {
	Validate() error
}
