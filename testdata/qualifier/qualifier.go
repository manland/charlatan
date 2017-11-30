// generated by "charlatan -dir=testdata/qualifier -output=testdata/qualifier/qualifier.go Qualifier".  DO NOT EDIT.

package main

import (
	"fmt"
	"reflect"
	"testing"
)

// QualifyInvocation represents a single call of FakeQualifier.Qualify
type QualifyInvocation struct {
	Parameters struct {
		Ident6 fmt.Scanner
	}
	Results struct {
		Ident7 fmt.Scanner
	}
}

// NamedQualifyInvocation represents a single call of FakeQualifier.NamedQualify
type NamedQualifyInvocation struct {
	Parameters struct {
		A fmt.Scanner
		B fmt.Scanner
		C fmt.Scanner
	}
	Results struct {
		D fmt.Scanner
	}
}

/*
FakeQualifier is a mock implementation of Qualifier for testing.
Use it in your tests as in this example:

	package example

	func TestWithQualifier(t *testing.T) {
		f := &main.FakeQualifier{
			QualifyHook: func(ident6 fmt.Scanner) (ident7 fmt.Scanner) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeQualify ...
		f.AssertQualifyCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakeQualify.
*/
type FakeQualifier struct {
	QualifyHook      func(fmt.Scanner) fmt.Scanner
	NamedQualifyHook func(fmt.Scanner, fmt.Scanner, fmt.Scanner) fmt.Scanner

	QualifyCalls      []*QualifyInvocation
	NamedQualifyCalls []*NamedQualifyInvocation
}

// NewFakeQualifierDefaultPanic returns an instance of FakeQualifier with all hooks configured to panic
func NewFakeQualifierDefaultPanic() *FakeQualifier {
	return &FakeQualifier{
		QualifyHook: func(fmt.Scanner) (ident7 fmt.Scanner) {
			panic("Unexpected call to Qualifier.Qualify")
			return
		},
		NamedQualifyHook: func(fmt.Scanner, fmt.Scanner, fmt.Scanner) (d fmt.Scanner) {
			panic("Unexpected call to Qualifier.NamedQualify")
			return
		},
	}
}

// NewFakeQualifierDefaultFatal returns an instance of FakeQualifier with all hooks configured to call t.Fatal
func NewFakeQualifierDefaultFatal(t *testing.T) *FakeQualifier {
	return &FakeQualifier{
		QualifyHook: func(fmt.Scanner) (ident7 fmt.Scanner) {
			t.Fatal("Unexpected call to Qualifier.Qualify")
			return
		},
		NamedQualifyHook: func(fmt.Scanner, fmt.Scanner, fmt.Scanner) (d fmt.Scanner) {
			t.Fatal("Unexpected call to Qualifier.NamedQualify")
			return
		},
	}
}

// NewFakeQualifierDefaultError returns an instance of FakeQualifier with all hooks configured to call t.Error
func NewFakeQualifierDefaultError(t *testing.T) *FakeQualifier {
	return &FakeQualifier{
		QualifyHook: func(fmt.Scanner) (ident7 fmt.Scanner) {
			t.Error("Unexpected call to Qualifier.Qualify")
			return
		},
		NamedQualifyHook: func(fmt.Scanner, fmt.Scanner, fmt.Scanner) (d fmt.Scanner) {
			t.Error("Unexpected call to Qualifier.NamedQualify")
			return
		},
	}
}

func (_f1 *FakeQualifier) Qualify(ident6 fmt.Scanner) (ident7 fmt.Scanner) {
	invocation := new(QualifyInvocation)

	invocation.Parameters.Ident6 = ident6

	ident7 = _f1.QualifyHook(ident6)

	invocation.Results.Ident7 = ident7

	_f1.QualifyCalls = append(_f1.QualifyCalls, invocation)

	return
}

// QualifyCalled returns true if FakeQualifier.Qualify was called
func (f *FakeQualifier) QualifyCalled() bool {
	return len(f.QualifyCalls) != 0
}

// AssertQualifyCalled calls t.Error if FakeQualifier.Qualify was not called
func (f *FakeQualifier) AssertQualifyCalled(t *testing.T) {
	t.Helper()
	if len(f.QualifyCalls) == 0 {
		t.Error("FakeQualifier.Qualify not called, expected at least one")
	}
}

// QualifyNotCalled returns true if FakeQualifier.Qualify was not called
func (f *FakeQualifier) QualifyNotCalled() bool {
	return len(f.QualifyCalls) == 0
}

// AssertQualifyNotCalled calls t.Error if FakeQualifier.Qualify was called
func (f *FakeQualifier) AssertQualifyNotCalled(t *testing.T) {
	t.Helper()
	if len(f.QualifyCalls) != 0 {
		t.Error("FakeQualifier.Qualify called, expected none")
	}
}

// QualifyCalledOnce returns true if FakeQualifier.Qualify was called exactly once
func (f *FakeQualifier) QualifyCalledOnce() bool {
	return len(f.QualifyCalls) == 1
}

// AssertQualifyCalledOnce calls t.Error if FakeQualifier.Qualify was not called exactly once
func (f *FakeQualifier) AssertQualifyCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.QualifyCalls) != 1 {
		t.Errorf("FakeQualifier.Qualify called %d times, expected 1", len(f.QualifyCalls))
	}
}

// QualifyCalledN returns true if FakeQualifier.Qualify was called at least n times
func (f *FakeQualifier) QualifyCalledN(n int) bool {
	return len(f.QualifyCalls) >= n
}

// AssertQualifyCalledN calls t.Error if FakeQualifier.Qualify was called less than n times
func (f *FakeQualifier) AssertQualifyCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.QualifyCalls) < n {
		t.Errorf("FakeQualifier.Qualify called %d times, expected >= %d", len(f.QualifyCalls), n)
	}
}

// QualifyCalledWith returns true if FakeQualifier.Qualify was called with the given values
func (_f2 *FakeQualifier) QualifyCalledWith(ident6 fmt.Scanner) (found bool) {
	for _, call := range _f2.QualifyCalls {
		if reflect.DeepEqual(call.Parameters.Ident6, ident6) {
			found = true
			break
		}
	}

	return
}

// AssertQualifyCalledWith calls t.Error if FakeQualifier.Qualify was not called with the given values
func (_f3 *FakeQualifier) AssertQualifyCalledWith(t *testing.T, ident6 fmt.Scanner) {
	t.Helper()
	var found bool
	for _, call := range _f3.QualifyCalls {
		if reflect.DeepEqual(call.Parameters.Ident6, ident6) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeQualifier.Qualify not called with expected parameters")
	}
}

// QualifyCalledOnceWith returns true if FakeQualifier.Qualify was called exactly once with the given values
func (_f4 *FakeQualifier) QualifyCalledOnceWith(ident6 fmt.Scanner) bool {
	var count int
	for _, call := range _f4.QualifyCalls {
		if reflect.DeepEqual(call.Parameters.Ident6, ident6) {
			count++
		}
	}

	return count == 1
}

// AssertQualifyCalledOnceWith calls t.Error if FakeQualifier.Qualify was not called exactly once with the given values
func (_f5 *FakeQualifier) AssertQualifyCalledOnceWith(t *testing.T, ident6 fmt.Scanner) {
	t.Helper()
	var count int
	for _, call := range _f5.QualifyCalls {
		if reflect.DeepEqual(call.Parameters.Ident6, ident6) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeQualifier.Qualify called %d times with expected parameters, expected one", count)
	}
}

// QualifyResultsForCall returns the result values for the first call to FakeQualifier.Qualify with the given values
func (_f6 *FakeQualifier) QualifyResultsForCall(ident6 fmt.Scanner) (ident7 fmt.Scanner, found bool) {
	for _, call := range _f6.QualifyCalls {
		if reflect.DeepEqual(call.Parameters.Ident6, ident6) {
			ident7 = call.Results.Ident7
			found = true
			break
		}
	}

	return
}

func (_f7 *FakeQualifier) NamedQualify(a fmt.Scanner, b fmt.Scanner, c fmt.Scanner) (d fmt.Scanner) {
	invocation := new(NamedQualifyInvocation)

	invocation.Parameters.A = a
	invocation.Parameters.B = b
	invocation.Parameters.C = c

	d = _f7.NamedQualifyHook(a, b, c)

	invocation.Results.D = d

	_f7.NamedQualifyCalls = append(_f7.NamedQualifyCalls, invocation)

	return
}

// NamedQualifyCalled returns true if FakeQualifier.NamedQualify was called
func (f *FakeQualifier) NamedQualifyCalled() bool {
	return len(f.NamedQualifyCalls) != 0
}

// AssertNamedQualifyCalled calls t.Error if FakeQualifier.NamedQualify was not called
func (f *FakeQualifier) AssertNamedQualifyCalled(t *testing.T) {
	t.Helper()
	if len(f.NamedQualifyCalls) == 0 {
		t.Error("FakeQualifier.NamedQualify not called, expected at least one")
	}
}

// NamedQualifyNotCalled returns true if FakeQualifier.NamedQualify was not called
func (f *FakeQualifier) NamedQualifyNotCalled() bool {
	return len(f.NamedQualifyCalls) == 0
}

// AssertNamedQualifyNotCalled calls t.Error if FakeQualifier.NamedQualify was called
func (f *FakeQualifier) AssertNamedQualifyNotCalled(t *testing.T) {
	t.Helper()
	if len(f.NamedQualifyCalls) != 0 {
		t.Error("FakeQualifier.NamedQualify called, expected none")
	}
}

// NamedQualifyCalledOnce returns true if FakeQualifier.NamedQualify was called exactly once
func (f *FakeQualifier) NamedQualifyCalledOnce() bool {
	return len(f.NamedQualifyCalls) == 1
}

// AssertNamedQualifyCalledOnce calls t.Error if FakeQualifier.NamedQualify was not called exactly once
func (f *FakeQualifier) AssertNamedQualifyCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.NamedQualifyCalls) != 1 {
		t.Errorf("FakeQualifier.NamedQualify called %d times, expected 1", len(f.NamedQualifyCalls))
	}
}

// NamedQualifyCalledN returns true if FakeQualifier.NamedQualify was called at least n times
func (f *FakeQualifier) NamedQualifyCalledN(n int) bool {
	return len(f.NamedQualifyCalls) >= n
}

// AssertNamedQualifyCalledN calls t.Error if FakeQualifier.NamedQualify was called less than n times
func (f *FakeQualifier) AssertNamedQualifyCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.NamedQualifyCalls) < n {
		t.Errorf("FakeQualifier.NamedQualify called %d times, expected >= %d", len(f.NamedQualifyCalls), n)
	}
}

// NamedQualifyCalledWith returns true if FakeQualifier.NamedQualify was called with the given values
func (_f8 *FakeQualifier) NamedQualifyCalledWith(a fmt.Scanner, b fmt.Scanner, c fmt.Scanner) (found bool) {
	for _, call := range _f8.NamedQualifyCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) {
			found = true
			break
		}
	}

	return
}

// AssertNamedQualifyCalledWith calls t.Error if FakeQualifier.NamedQualify was not called with the given values
func (_f9 *FakeQualifier) AssertNamedQualifyCalledWith(t *testing.T, a fmt.Scanner, b fmt.Scanner, c fmt.Scanner) {
	t.Helper()
	var found bool
	for _, call := range _f9.NamedQualifyCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeQualifier.NamedQualify not called with expected parameters")
	}
}

// NamedQualifyCalledOnceWith returns true if FakeQualifier.NamedQualify was called exactly once with the given values
func (_f10 *FakeQualifier) NamedQualifyCalledOnceWith(a fmt.Scanner, b fmt.Scanner, c fmt.Scanner) bool {
	var count int
	for _, call := range _f10.NamedQualifyCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) {
			count++
		}
	}

	return count == 1
}

// AssertNamedQualifyCalledOnceWith calls t.Error if FakeQualifier.NamedQualify was not called exactly once with the given values
func (_f11 *FakeQualifier) AssertNamedQualifyCalledOnceWith(t *testing.T, a fmt.Scanner, b fmt.Scanner, c fmt.Scanner) {
	t.Helper()
	var count int
	for _, call := range _f11.NamedQualifyCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeQualifier.NamedQualify called %d times with expected parameters, expected one", count)
	}
}

// NamedQualifyResultsForCall returns the result values for the first call to FakeQualifier.NamedQualify with the given values
func (_f12 *FakeQualifier) NamedQualifyResultsForCall(a fmt.Scanner, b fmt.Scanner, c fmt.Scanner) (d fmt.Scanner, found bool) {
	for _, call := range _f12.NamedQualifyCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) {
			d = call.Results.D
			found = true
			break
		}
	}

	return
}
