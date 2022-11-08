package result

import (
	"errors"
	"testing"
)

func TestResult(t *testing.T) {
	r := Ok(1)

	if !r.IsOk() {
		t.Error("expected result to be ok")
	}

	if r.IsErr() {
		t.Error("expected result to not be an error")
	}

	if !r.IsOkAnd(func(v int) bool { return v == 1 }) {
		t.Error("expected result to be ok and have value 1")
	}

	if r.IsErrAnd(func(e error) bool { return e != nil }) {
		t.Error("expected result to not be an error and have error nil")
	}

	v := r.Ok()
	if *v != 1 {
		t.Error("expected result to have value 1")
	}

	err := r.Map(func(v *int) { *v = 2 })
	if err != nil {
		t.Error("expected result to not have error")
	}

	if *v != 2 {
		t.Error("expected result to have value 2")
	}
}

func TestErrResult(t *testing.T) {
	r := Err[int](errors.New("error"))

	if r.IsOk() {
		t.Error("expected result to not be ok")
	}

	if !r.IsErr() {
		t.Error("expected result to be an error")
	}

	if r.IsOkAnd(func(v int) bool { return v == 1 }) {
		t.Error("expected result to not be ok and have value 1")
	}

	if !r.IsErrAnd(func(e error) bool { return e != nil }) {
		t.Error("expected result to be an error and have error nil")
	}

	if *r.Err() == nil {
		t.Error("expected result to have error")
	}

	if r.UnwrapErr() == nil {
		t.Error("expected result to have error")
	}
}
