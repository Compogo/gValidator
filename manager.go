package gValidator

import (
	"context"
	"errors"
	"fmt"
)

type Manager[T any] struct {
	validators []Validator[T]
}

func NewManager[T any](validators ...Validator[T]) *Manager[T] {
	return &Manager[T]{validators: validators}
}

func (m *Manager[T]) ValidateMany(ctx context.Context, ts ...*T) error {
	var errs []error
	for index, t := range ts {
		if err := m.Validate(ctx, t, nil); err != nil {
			errs = append(errs, fmt.Errorf("record index: %d, validate error: %w", index, err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (m *Manager[T]) Validate(ctx context.Context, new, old *T) error {
	var errs []error

	for _, validator := range m.validators {
		if err := validator.Validate(ctx, new, old); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
