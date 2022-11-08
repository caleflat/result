package result

import (
	"testing"
)

func TestResult(t *testing.T) {
	r := Ok(1)
	r.value = 1
	r.err = nil

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

	r.Map(func(v int) int { return v + 1 })
	if *v != 2 {
		t.Error("expected result to have value 2")
	}
}
