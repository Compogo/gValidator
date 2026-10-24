package gValidator

import "context"

type Validator[T any] interface {
	Validate(context.Context, *T) error
}

type ValidateFunc[T any] func(context.Context, *T) error

func (v ValidateFunc[T]) Validate(ctx context.Context, t *T) error {
	return v(ctx, t)
}
