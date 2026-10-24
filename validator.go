package gValidator

type Validator[T any] interface {
	Validate(*T) error
}

type ValidateFunc[T any] func(*T) error

func (v ValidateFunc[T]) Validate(t *T) error {
	return v(t)
}
