package validation

type Checker interface {
	// Name returns checker name
	Name() string

	// Check tests input value with expectation
	Check(v interface{}, expects string) error
}
