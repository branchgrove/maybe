package maybe

import (
	"github.com/branchgrove/kaeru"
)

// Contains some or none value based on ok
type Maybe[T any] struct {
	value T
	ok    bool
}

// Ensures compatability with kaeru
func (o *Maybe[T]) ParseAny(v any) error {
	if v == nil {
		o.ok = false
		return nil
	}

	err := kaeru.Parse(v, &o.value)

	if err != nil {
		return err
	}

	o.ok = true

	return nil
}

// Default value for kaeru parsing
func (o *Maybe[T]) SetDefault() {
	*o = None[T]()
}

// Create a new Maybe containing a value
func Some[T any](v T) Maybe[T] {
	return Maybe[T]{
		value: v,
		ok:    true,
	}
}

// Create a new Maybe without a value
func None[T any]() Maybe[T] {
	return Maybe[T]{
		ok: false,
	}
}

// Returns the value or panics if it is None
func (o Maybe[T]) Unwrap() T {
	if o.ok == false {
		panic("Tried to unwrap None value")
	}

	return o.value
}

// Returns the value and ok, if ok is true there is a value
func (o Maybe[T]) Get() (T, bool) {
	return o.value, o.ok
}

func (o Maybe[T]) IsNone() bool {
	return !o.ok
}

func (o Maybe[T]) IsSome() bool {
	return o.ok
}
