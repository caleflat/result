package result

import "fmt"

// Result[T] is a type that represents either a value of type T or an error.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a new Result[T] with the given value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// Err creates a new Result[T] with the given error.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// IsOk returns true if the result is ok.
func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the result is an error.
func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

// IsOkAnd returns true if the result is ok and the value is equal to the given
// value.
func (r *Result[T]) IsOkAnd(f func(T) bool) bool {
	return r.IsOk() && f(r.value)
}

// IsErrAnd returns true if the result is an error and the error satisfies the
// given predicate.
func (r *Result[T]) IsErrAnd(f func(error) bool) bool {
	return r.IsErr() && f(r.err)
}

// Ok returns the value if the result is ok, otherwise panics.
func (r *Result[T]) Ok() *T {
	if r.IsErr() {
		panic(r.err)
	}

	return &r.value
}

// Err returns the error if the result is an error, otherwise panics.
func (r *Result[T]) Err() *error {
	if r.IsOk() {
		panic("result is ok")
	}

	return &r.err
}

// Map maps the value if the result is ok, otherwise returns the error.
func (r *Result[T]) Map(f func(T) T) *Result[T] {
	if r.IsErr() {
		return r
	}

	return &Result[T]{value: f(r.value)}
}

// MapOr maps the value if the result is ok, otherwise returns the given value.
func (r *Result[T]) MapOr(f func(T) T, or T) T {
	if r.IsErr() {
		return or
	}

	return f(r.value)
}

// MapOrElse maps the value if the result is ok, otherwise returns the value
// returned by the given function.
func (r *Result[T]) MapOrElse(f func(T) T, or func() T) T {
	if r.IsErr() {
		return or()
	}

	return f(r.value)
}

// MapErr maps the error if the result is an error, otherwise returns the value.
func (r *Result[T]) MapErr(f func(error) error) *Result[T] {
	if r.IsOk() {
		return r
	}

	return &Result[T]{err: f(r.err)}
}

// Inspect calls the given function with the value if the result is ok, otherwise
// returns the error.
func (r *Result[T]) Inspect(f func(T)) *Result[T] {
	if r.IsErr() {
		return r
	}

	f(r.value)
	return r
}

// InspectErr calls the given function with the error if the result is an error,
// otherwise returns the value.
func (r *Result[T]) InspectErr(f func(error)) *Result[T] {
	if r.IsOk() {
		return r
	}

	f(r.err)
	return r
}

// Expect returns the value if the result is ok, otherwise panics with the given
// message.
func (r *Result[T]) Expect(msg string) T {
	if r.IsErr() {
		panic(msg)
	}

	return r.value
}

// Unwrap returns the value if the result is ok, otherwise panics.
func (r *Result[T]) Unwrap() T {
	if r.IsErr() {
		msg := fmt.Sprintf("called `Result[T].Unwrap()` on an `Err` value: %v", r.err)
		panic(msg)
	}

	return r.value
}
