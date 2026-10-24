package gValidator

import "context"

type Validator[T any] interface {
	Validate(context.Context, *T, *T) error
}

type ValidateFunc[T any] func(context.Context, *T, *T) error

func (v ValidateFunc[T]) Validate(ctx context.Context, new, old *T) error {
	return v(ctx, new, old)
}
