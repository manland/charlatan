// generated by "charlatan -dir=testdata/array -output=testdata/array/array.go Array".  DO NOT EDIT.

package main

import "reflect"

// ArrayArrayParameterInvocation represents a single call of FakeArray.ArrayParameter
type ArrayArrayParameterInvocation struct {
	Parameters struct {
		Ident1 [3]string
	}
}

// ArrayArrayReturnInvocation represents a single call of FakeArray.ArrayReturn
type ArrayArrayReturnInvocation struct {
	Results struct {
		Ident1 [3]string
	}
}

// ArraySliceParameterInvocation represents a single call of FakeArray.SliceParameter
type ArraySliceParameterInvocation struct {
	Parameters struct {
		Ident1 []string
	}
}

// ArraySliceReturnInvocation represents a single call of FakeArray.SliceReturn
type ArraySliceReturnInvocation struct {
	Results struct {
		Ident1 []string
	}
}

// ArrayTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type ArrayTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeArray is a mock implementation of Array for testing.
Use it in your tests as in this example:

	package example

	func TestWithArray(t *testing.T) {
		f := &main.FakeArray{
			ArrayParameterHook: func(ident1 [3]string) () {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeArrayParameter ...
		f.AssertArrayParameterCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeArrayParameter.
*/
type FakeArray struct {
	ArrayParameterHook func([3]string)
	ArrayReturnHook    func() [3]string
	SliceParameterHook func([]string)
	SliceReturnHook    func() []string

	ArrayParameterCalls []*ArrayArrayParameterInvocation
	ArrayReturnCalls    []*ArrayArrayReturnInvocation
	SliceParameterCalls []*ArraySliceParameterInvocation
	SliceReturnCalls    []*ArraySliceReturnInvocation
}

// NewFakeArrayDefaultPanic returns an instance of FakeArray with all hooks configured to panic
func NewFakeArrayDefaultPanic() *FakeArray {
	return &FakeArray{
		ArrayParameterHook: func([3]string) {
			panic("Unexpected call to Array.ArrayParameter")
		},
		ArrayReturnHook: func() (ident1 [3]string) {
			panic("Unexpected call to Array.ArrayReturn")
		},
		SliceParameterHook: func([]string) {
			panic("Unexpected call to Array.SliceParameter")
		},
		SliceReturnHook: func() (ident1 []string) {
			panic("Unexpected call to Array.SliceReturn")
		},
	}
}

// NewFakeArrayDefaultFatal returns an instance of FakeArray with all hooks configured to call t.Fatal
func NewFakeArrayDefaultFatal(t_sym1 ArrayTestingT) *FakeArray {
	return &FakeArray{
		ArrayParameterHook: func([3]string) {
			t_sym1.Fatal("Unexpected call to Array.ArrayParameter")
			return
		},
		ArrayReturnHook: func() (ident1 [3]string) {
			t_sym1.Fatal("Unexpected call to Array.ArrayReturn")
			return
		},
		SliceParameterHook: func([]string) {
			t_sym1.Fatal("Unexpected call to Array.SliceParameter")
			return
		},
		SliceReturnHook: func() (ident1 []string) {
			t_sym1.Fatal("Unexpected call to Array.SliceReturn")
			return
		},
	}
}

// NewFakeArrayDefaultError returns an instance of FakeArray with all hooks configured to call t.Error
func NewFakeArrayDefaultError(t_sym2 ArrayTestingT) *FakeArray {
	return &FakeArray{
		ArrayParameterHook: func([3]string) {
			t_sym2.Error("Unexpected call to Array.ArrayParameter")
			return
		},
		ArrayReturnHook: func() (ident1 [3]string) {
			t_sym2.Error("Unexpected call to Array.ArrayReturn")
			return
		},
		SliceParameterHook: func([]string) {
			t_sym2.Error("Unexpected call to Array.SliceParameter")
			return
		},
		SliceReturnHook: func() (ident1 []string) {
			t_sym2.Error("Unexpected call to Array.SliceReturn")
			return
		},
	}
}

func (f *FakeArray) Reset() {
	f.ArrayParameterCalls = []*ArrayArrayParameterInvocation{}
	f.ArrayReturnCalls = []*ArrayArrayReturnInvocation{}
	f.SliceParameterCalls = []*ArraySliceParameterInvocation{}
	f.SliceReturnCalls = []*ArraySliceReturnInvocation{}
}

func (f_sym3 *FakeArray) ArrayParameter(ident1 [3]string) {
	if f_sym3.ArrayParameterHook == nil {
		panic("Array.ArrayParameter() called but FakeArray.ArrayParameterHook is nil")
	}

	invocation_sym3 := new(ArrayArrayParameterInvocation)
	f_sym3.ArrayParameterCalls = append(f_sym3.ArrayParameterCalls, invocation_sym3)

	invocation_sym3.Parameters.Ident1 = ident1

	f_sym3.ArrayParameterHook(ident1)

	return
}

// ArrayParameterCalled returns true if FakeArray.ArrayParameter was called
func (f *FakeArray) ArrayParameterCalled() bool {
	return len(f.ArrayParameterCalls) != 0
}

// AssertArrayParameterCalled calls t.Error if FakeArray.ArrayParameter was not called
func (f *FakeArray) AssertArrayParameterCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.ArrayParameterCalls) == 0 {
		t.Error("FakeArray.ArrayParameter not called, expected at least one")
	}
}

// ArrayParameterNotCalled returns true if FakeArray.ArrayParameter was not called
func (f *FakeArray) ArrayParameterNotCalled() bool {
	return len(f.ArrayParameterCalls) == 0
}

// AssertArrayParameterNotCalled calls t.Error if FakeArray.ArrayParameter was called
func (f *FakeArray) AssertArrayParameterNotCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.ArrayParameterCalls) != 0 {
		t.Error("FakeArray.ArrayParameter called, expected none")
	}
}

// ArrayParameterCalledOnce returns true if FakeArray.ArrayParameter was called exactly once
func (f *FakeArray) ArrayParameterCalledOnce() bool {
	return len(f.ArrayParameterCalls) == 1
}

// AssertArrayParameterCalledOnce calls t.Error if FakeArray.ArrayParameter was not called exactly once
func (f *FakeArray) AssertArrayParameterCalledOnce(t ArrayTestingT) {
	t.Helper()
	if len(f.ArrayParameterCalls) != 1 {
		t.Errorf("FakeArray.ArrayParameter called %d times, expected 1", len(f.ArrayParameterCalls))
	}
}

// ArrayParameterCalledN returns true if FakeArray.ArrayParameter was called at least n times
func (f *FakeArray) ArrayParameterCalledN(n int) bool {
	return len(f.ArrayParameterCalls) >= n
}

// AssertArrayParameterCalledN calls t.Error if FakeArray.ArrayParameter was called less than n times
func (f *FakeArray) AssertArrayParameterCalledN(t ArrayTestingT, n int) {
	t.Helper()
	if len(f.ArrayParameterCalls) < n {
		t.Errorf("FakeArray.ArrayParameter called %d times, expected >= %d", len(f.ArrayParameterCalls), n)
	}
}

// ArrayParameterCalledWith returns true if FakeArray.ArrayParameter was called with the given values
func (f_sym4 *FakeArray) ArrayParameterCalledWith(ident1 [3]string) bool {
	for _, call_sym4 := range f_sym4.ArrayParameterCalls {
		if reflect.DeepEqual(call_sym4.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertArrayParameterCalledWith calls t.Error if FakeArray.ArrayParameter was not called with the given values
func (f_sym5 *FakeArray) AssertArrayParameterCalledWith(t ArrayTestingT, ident1 [3]string) {
	t.Helper()
	var found_sym5 bool
	for _, call_sym5 := range f_sym5.ArrayParameterCalls {
		if reflect.DeepEqual(call_sym5.Parameters.Ident1, ident1) {
			found_sym5 = true
			break
		}
	}

	if !found_sym5 {
		t.Error("FakeArray.ArrayParameter not called with expected parameters")
	}
}

// ArrayParameterCalledOnceWith returns true if FakeArray.ArrayParameter was called exactly once with the given values
func (f_sym6 *FakeArray) ArrayParameterCalledOnceWith(ident1 [3]string) bool {
	var count_sym6 int
	for _, call_sym6 := range f_sym6.ArrayParameterCalls {
		if reflect.DeepEqual(call_sym6.Parameters.Ident1, ident1) {
			count_sym6++
		}
	}

	return count_sym6 == 1
}

// AssertArrayParameterCalledOnceWith calls t.Error if FakeArray.ArrayParameter was not called exactly once with the given values
func (f_sym7 *FakeArray) AssertArrayParameterCalledOnceWith(t ArrayTestingT, ident1 [3]string) {
	t.Helper()
	var count_sym7 int
	for _, call_sym7 := range f_sym7.ArrayParameterCalls {
		if reflect.DeepEqual(call_sym7.Parameters.Ident1, ident1) {
			count_sym7++
		}
	}

	if count_sym7 != 1 {
		t.Errorf("FakeArray.ArrayParameter called %d times with expected parameters, expected one", count_sym7)
	}
}

func (f_sym8 *FakeArray) ArrayReturn() (ident1 [3]string) {
	if f_sym8.ArrayReturnHook == nil {
		panic("Array.ArrayReturn() called but FakeArray.ArrayReturnHook is nil")
	}

	invocation_sym8 := new(ArrayArrayReturnInvocation)
	f_sym8.ArrayReturnCalls = append(f_sym8.ArrayReturnCalls, invocation_sym8)

	ident1 = f_sym8.ArrayReturnHook()

	invocation_sym8.Results.Ident1 = ident1

	return
}

// SetArrayReturnStub configures Array.ArrayReturn to always return the given values
func (f_sym9 *FakeArray) SetArrayReturnStub(ident1 [3]string) {
	f_sym9.ArrayReturnHook = func() [3]string {
		return ident1
	}
}

// ArrayReturnCalled returns true if FakeArray.ArrayReturn was called
func (f *FakeArray) ArrayReturnCalled() bool {
	return len(f.ArrayReturnCalls) != 0
}

// AssertArrayReturnCalled calls t.Error if FakeArray.ArrayReturn was not called
func (f *FakeArray) AssertArrayReturnCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.ArrayReturnCalls) == 0 {
		t.Error("FakeArray.ArrayReturn not called, expected at least one")
	}
}

// ArrayReturnNotCalled returns true if FakeArray.ArrayReturn was not called
func (f *FakeArray) ArrayReturnNotCalled() bool {
	return len(f.ArrayReturnCalls) == 0
}

// AssertArrayReturnNotCalled calls t.Error if FakeArray.ArrayReturn was called
func (f *FakeArray) AssertArrayReturnNotCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.ArrayReturnCalls) != 0 {
		t.Error("FakeArray.ArrayReturn called, expected none")
	}
}

// ArrayReturnCalledOnce returns true if FakeArray.ArrayReturn was called exactly once
func (f *FakeArray) ArrayReturnCalledOnce() bool {
	return len(f.ArrayReturnCalls) == 1
}

// AssertArrayReturnCalledOnce calls t.Error if FakeArray.ArrayReturn was not called exactly once
func (f *FakeArray) AssertArrayReturnCalledOnce(t ArrayTestingT) {
	t.Helper()
	if len(f.ArrayReturnCalls) != 1 {
		t.Errorf("FakeArray.ArrayReturn called %d times, expected 1", len(f.ArrayReturnCalls))
	}
}

// ArrayReturnCalledN returns true if FakeArray.ArrayReturn was called at least n times
func (f *FakeArray) ArrayReturnCalledN(n int) bool {
	return len(f.ArrayReturnCalls) >= n
}

// AssertArrayReturnCalledN calls t.Error if FakeArray.ArrayReturn was called less than n times
func (f *FakeArray) AssertArrayReturnCalledN(t ArrayTestingT, n int) {
	t.Helper()
	if len(f.ArrayReturnCalls) < n {
		t.Errorf("FakeArray.ArrayReturn called %d times, expected >= %d", len(f.ArrayReturnCalls), n)
	}
}

func (f_sym10 *FakeArray) SliceParameter(ident1 []string) {
	if f_sym10.SliceParameterHook == nil {
		panic("Array.SliceParameter() called but FakeArray.SliceParameterHook is nil")
	}

	invocation_sym10 := new(ArraySliceParameterInvocation)
	f_sym10.SliceParameterCalls = append(f_sym10.SliceParameterCalls, invocation_sym10)

	invocation_sym10.Parameters.Ident1 = ident1

	f_sym10.SliceParameterHook(ident1)

	return
}

// SliceParameterCalled returns true if FakeArray.SliceParameter was called
func (f *FakeArray) SliceParameterCalled() bool {
	return len(f.SliceParameterCalls) != 0
}

// AssertSliceParameterCalled calls t.Error if FakeArray.SliceParameter was not called
func (f *FakeArray) AssertSliceParameterCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.SliceParameterCalls) == 0 {
		t.Error("FakeArray.SliceParameter not called, expected at least one")
	}
}

// SliceParameterNotCalled returns true if FakeArray.SliceParameter was not called
func (f *FakeArray) SliceParameterNotCalled() bool {
	return len(f.SliceParameterCalls) == 0
}

// AssertSliceParameterNotCalled calls t.Error if FakeArray.SliceParameter was called
func (f *FakeArray) AssertSliceParameterNotCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.SliceParameterCalls) != 0 {
		t.Error("FakeArray.SliceParameter called, expected none")
	}
}

// SliceParameterCalledOnce returns true if FakeArray.SliceParameter was called exactly once
func (f *FakeArray) SliceParameterCalledOnce() bool {
	return len(f.SliceParameterCalls) == 1
}

// AssertSliceParameterCalledOnce calls t.Error if FakeArray.SliceParameter was not called exactly once
func (f *FakeArray) AssertSliceParameterCalledOnce(t ArrayTestingT) {
	t.Helper()
	if len(f.SliceParameterCalls) != 1 {
		t.Errorf("FakeArray.SliceParameter called %d times, expected 1", len(f.SliceParameterCalls))
	}
}

// SliceParameterCalledN returns true if FakeArray.SliceParameter was called at least n times
func (f *FakeArray) SliceParameterCalledN(n int) bool {
	return len(f.SliceParameterCalls) >= n
}

// AssertSliceParameterCalledN calls t.Error if FakeArray.SliceParameter was called less than n times
func (f *FakeArray) AssertSliceParameterCalledN(t ArrayTestingT, n int) {
	t.Helper()
	if len(f.SliceParameterCalls) < n {
		t.Errorf("FakeArray.SliceParameter called %d times, expected >= %d", len(f.SliceParameterCalls), n)
	}
}

// SliceParameterCalledWith returns true if FakeArray.SliceParameter was called with the given values
func (f_sym11 *FakeArray) SliceParameterCalledWith(ident1 []string) bool {
	for _, call_sym11 := range f_sym11.SliceParameterCalls {
		if reflect.DeepEqual(call_sym11.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertSliceParameterCalledWith calls t.Error if FakeArray.SliceParameter was not called with the given values
func (f_sym12 *FakeArray) AssertSliceParameterCalledWith(t ArrayTestingT, ident1 []string) {
	t.Helper()
	var found_sym12 bool
	for _, call_sym12 := range f_sym12.SliceParameterCalls {
		if reflect.DeepEqual(call_sym12.Parameters.Ident1, ident1) {
			found_sym12 = true
			break
		}
	}

	if !found_sym12 {
		t.Error("FakeArray.SliceParameter not called with expected parameters")
	}
}

// SliceParameterCalledOnceWith returns true if FakeArray.SliceParameter was called exactly once with the given values
func (f_sym13 *FakeArray) SliceParameterCalledOnceWith(ident1 []string) bool {
	var count_sym13 int
	for _, call_sym13 := range f_sym13.SliceParameterCalls {
		if reflect.DeepEqual(call_sym13.Parameters.Ident1, ident1) {
			count_sym13++
		}
	}

	return count_sym13 == 1
}

// AssertSliceParameterCalledOnceWith calls t.Error if FakeArray.SliceParameter was not called exactly once with the given values
func (f_sym14 *FakeArray) AssertSliceParameterCalledOnceWith(t ArrayTestingT, ident1 []string) {
	t.Helper()
	var count_sym14 int
	for _, call_sym14 := range f_sym14.SliceParameterCalls {
		if reflect.DeepEqual(call_sym14.Parameters.Ident1, ident1) {
			count_sym14++
		}
	}

	if count_sym14 != 1 {
		t.Errorf("FakeArray.SliceParameter called %d times with expected parameters, expected one", count_sym14)
	}
}

func (f_sym15 *FakeArray) SliceReturn() (ident1 []string) {
	if f_sym15.SliceReturnHook == nil {
		panic("Array.SliceReturn() called but FakeArray.SliceReturnHook is nil")
	}

	invocation_sym15 := new(ArraySliceReturnInvocation)
	f_sym15.SliceReturnCalls = append(f_sym15.SliceReturnCalls, invocation_sym15)

	ident1 = f_sym15.SliceReturnHook()

	invocation_sym15.Results.Ident1 = ident1

	return
}

// SetSliceReturnStub configures Array.SliceReturn to always return the given values
func (f_sym16 *FakeArray) SetSliceReturnStub(ident1 []string) {
	f_sym16.SliceReturnHook = func() []string {
		return ident1
	}
}

// SliceReturnCalled returns true if FakeArray.SliceReturn was called
func (f *FakeArray) SliceReturnCalled() bool {
	return len(f.SliceReturnCalls) != 0
}

// AssertSliceReturnCalled calls t.Error if FakeArray.SliceReturn was not called
func (f *FakeArray) AssertSliceReturnCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.SliceReturnCalls) == 0 {
		t.Error("FakeArray.SliceReturn not called, expected at least one")
	}
}

// SliceReturnNotCalled returns true if FakeArray.SliceReturn was not called
func (f *FakeArray) SliceReturnNotCalled() bool {
	return len(f.SliceReturnCalls) == 0
}

// AssertSliceReturnNotCalled calls t.Error if FakeArray.SliceReturn was called
func (f *FakeArray) AssertSliceReturnNotCalled(t ArrayTestingT) {
	t.Helper()
	if len(f.SliceReturnCalls) != 0 {
		t.Error("FakeArray.SliceReturn called, expected none")
	}
}

// SliceReturnCalledOnce returns true if FakeArray.SliceReturn was called exactly once
func (f *FakeArray) SliceReturnCalledOnce() bool {
	return len(f.SliceReturnCalls) == 1
}

// AssertSliceReturnCalledOnce calls t.Error if FakeArray.SliceReturn was not called exactly once
func (f *FakeArray) AssertSliceReturnCalledOnce(t ArrayTestingT) {
	t.Helper()
	if len(f.SliceReturnCalls) != 1 {
		t.Errorf("FakeArray.SliceReturn called %d times, expected 1", len(f.SliceReturnCalls))
	}
}

// SliceReturnCalledN returns true if FakeArray.SliceReturn was called at least n times
func (f *FakeArray) SliceReturnCalledN(n int) bool {
	return len(f.SliceReturnCalls) >= n
}

// AssertSliceReturnCalledN calls t.Error if FakeArray.SliceReturn was called less than n times
func (f *FakeArray) AssertSliceReturnCalledN(t ArrayTestingT, n int) {
	t.Helper()
	if len(f.SliceReturnCalls) < n {
		t.Errorf("FakeArray.SliceReturn called %d times, expected >= %d", len(f.SliceReturnCalls), n)
	}
}
