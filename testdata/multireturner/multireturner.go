// generated by "charlatan -dir=testdata/multireturner -output=testdata/multireturner/multireturner.go Multireturner".  DO NOT EDIT.

package main

// MultireturnerMultiReturnInvocation represents a single call of FakeMultireturner.MultiReturn
type MultireturnerMultiReturnInvocation struct {
	Results struct {
		Ident1 string
		Ident2 int
	}
}

// MultireturnerNamedReturnInvocation represents a single call of FakeMultireturner.NamedReturn
type MultireturnerNamedReturnInvocation struct {
	Results struct {
		A int
		B int
		C int
		D int
	}
}

// MultireturnerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type MultireturnerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeMultireturner is a mock implementation of Multireturner for testing.
Use it in your tests as in this example:

	package example

	func TestWithMultireturner(t *testing.T) {
		f := &main.FakeMultireturner{
			MultiReturnHook: func() (ident1 string, ident2 int) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeMultiReturn ...
		f.AssertMultiReturnCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeMultiReturn.
*/
type FakeMultireturner struct {
	MultiReturnHook func() (string, int)
	NamedReturnHook func() (int, int, int, int)

	MultiReturnCalls []*MultireturnerMultiReturnInvocation
	NamedReturnCalls []*MultireturnerNamedReturnInvocation
}

// NewFakeMultireturnerDefaultPanic returns an instance of FakeMultireturner with all hooks configured to panic
func NewFakeMultireturnerDefaultPanic() *FakeMultireturner {
	return &FakeMultireturner{
		MultiReturnHook: func() (ident1 string, ident2 int) {
			panic("Unexpected call to Multireturner.MultiReturn")
		},
		NamedReturnHook: func() (a int, b int, c int, d int) {
			panic("Unexpected call to Multireturner.NamedReturn")
		},
	}
}

// NewFakeMultireturnerDefaultFatal returns an instance of FakeMultireturner with all hooks configured to call t.Fatal
func NewFakeMultireturnerDefaultFatal(t_sym1 MultireturnerTestingT) *FakeMultireturner {
	return &FakeMultireturner{
		MultiReturnHook: func() (ident1 string, ident2 int) {
			t_sym1.Fatal("Unexpected call to Multireturner.MultiReturn")
			return
		},
		NamedReturnHook: func() (a int, b int, c int, d int) {
			t_sym1.Fatal("Unexpected call to Multireturner.NamedReturn")
			return
		},
	}
}

// NewFakeMultireturnerDefaultError returns an instance of FakeMultireturner with all hooks configured to call t.Error
func NewFakeMultireturnerDefaultError(t_sym2 MultireturnerTestingT) *FakeMultireturner {
	return &FakeMultireturner{
		MultiReturnHook: func() (ident1 string, ident2 int) {
			t_sym2.Error("Unexpected call to Multireturner.MultiReturn")
			return
		},
		NamedReturnHook: func() (a int, b int, c int, d int) {
			t_sym2.Error("Unexpected call to Multireturner.NamedReturn")
			return
		},
	}
}

func (f *FakeMultireturner) Reset() {
	f.MultiReturnCalls = []*MultireturnerMultiReturnInvocation{}
	f.NamedReturnCalls = []*MultireturnerNamedReturnInvocation{}
}

func (f_sym3 *FakeMultireturner) MultiReturn() (ident1 string, ident2 int) {
	if f_sym3.MultiReturnHook == nil {
		panic("Multireturner.MultiReturn() called but FakeMultireturner.MultiReturnHook is nil")
	}

	invocation_sym3 := new(MultireturnerMultiReturnInvocation)
	f_sym3.MultiReturnCalls = append(f_sym3.MultiReturnCalls, invocation_sym3)

	ident1, ident2 = f_sym3.MultiReturnHook()

	invocation_sym3.Results.Ident1 = ident1
	invocation_sym3.Results.Ident2 = ident2

	return
}

// SetMultiReturnStub configures Multireturner.MultiReturn to always return the given values
func (f_sym4 *FakeMultireturner) SetMultiReturnStub(ident1 string, ident2 int) {
	f_sym4.MultiReturnHook = func() (string, int) {
		return ident1, ident2
	}
}

// MultiReturnCalled returns true if FakeMultireturner.MultiReturn was called
func (f *FakeMultireturner) MultiReturnCalled() bool {
	return len(f.MultiReturnCalls) != 0
}

// AssertMultiReturnCalled calls t.Error if FakeMultireturner.MultiReturn was not called
func (f *FakeMultireturner) AssertMultiReturnCalled(t MultireturnerTestingT) {
	t.Helper()
	if len(f.MultiReturnCalls) == 0 {
		t.Error("FakeMultireturner.MultiReturn not called, expected at least one")
	}
}

// MultiReturnNotCalled returns true if FakeMultireturner.MultiReturn was not called
func (f *FakeMultireturner) MultiReturnNotCalled() bool {
	return len(f.MultiReturnCalls) == 0
}

// AssertMultiReturnNotCalled calls t.Error if FakeMultireturner.MultiReturn was called
func (f *FakeMultireturner) AssertMultiReturnNotCalled(t MultireturnerTestingT) {
	t.Helper()
	if len(f.MultiReturnCalls) != 0 {
		t.Error("FakeMultireturner.MultiReturn called, expected none")
	}
}

// MultiReturnCalledOnce returns true if FakeMultireturner.MultiReturn was called exactly once
func (f *FakeMultireturner) MultiReturnCalledOnce() bool {
	return len(f.MultiReturnCalls) == 1
}

// AssertMultiReturnCalledOnce calls t.Error if FakeMultireturner.MultiReturn was not called exactly once
func (f *FakeMultireturner) AssertMultiReturnCalledOnce(t MultireturnerTestingT) {
	t.Helper()
	if len(f.MultiReturnCalls) != 1 {
		t.Errorf("FakeMultireturner.MultiReturn called %d times, expected 1", len(f.MultiReturnCalls))
	}
}

// MultiReturnCalledN returns true if FakeMultireturner.MultiReturn was called at least n times
func (f *FakeMultireturner) MultiReturnCalledN(n int) bool {
	return len(f.MultiReturnCalls) >= n
}

// AssertMultiReturnCalledN calls t.Error if FakeMultireturner.MultiReturn was called less than n times
func (f *FakeMultireturner) AssertMultiReturnCalledN(t MultireturnerTestingT, n int) {
	t.Helper()
	if len(f.MultiReturnCalls) < n {
		t.Errorf("FakeMultireturner.MultiReturn called %d times, expected >= %d", len(f.MultiReturnCalls), n)
	}
}

func (f_sym5 *FakeMultireturner) NamedReturn() (a int, b int, c int, d int) {
	if f_sym5.NamedReturnHook == nil {
		panic("Multireturner.NamedReturn() called but FakeMultireturner.NamedReturnHook is nil")
	}

	invocation_sym5 := new(MultireturnerNamedReturnInvocation)
	f_sym5.NamedReturnCalls = append(f_sym5.NamedReturnCalls, invocation_sym5)

	a, b, c, d = f_sym5.NamedReturnHook()

	invocation_sym5.Results.A = a
	invocation_sym5.Results.B = b
	invocation_sym5.Results.C = c
	invocation_sym5.Results.D = d

	return
}

// SetNamedReturnStub configures Multireturner.NamedReturn to always return the given values
func (f_sym6 *FakeMultireturner) SetNamedReturnStub(a int, b int, c int, d int) {
	f_sym6.NamedReturnHook = func() (int, int, int, int) {
		return a, b, c, d
	}
}

// NamedReturnCalled returns true if FakeMultireturner.NamedReturn was called
func (f *FakeMultireturner) NamedReturnCalled() bool {
	return len(f.NamedReturnCalls) != 0
}

// AssertNamedReturnCalled calls t.Error if FakeMultireturner.NamedReturn was not called
func (f *FakeMultireturner) AssertNamedReturnCalled(t MultireturnerTestingT) {
	t.Helper()
	if len(f.NamedReturnCalls) == 0 {
		t.Error("FakeMultireturner.NamedReturn not called, expected at least one")
	}
}

// NamedReturnNotCalled returns true if FakeMultireturner.NamedReturn was not called
func (f *FakeMultireturner) NamedReturnNotCalled() bool {
	return len(f.NamedReturnCalls) == 0
}

// AssertNamedReturnNotCalled calls t.Error if FakeMultireturner.NamedReturn was called
func (f *FakeMultireturner) AssertNamedReturnNotCalled(t MultireturnerTestingT) {
	t.Helper()
	if len(f.NamedReturnCalls) != 0 {
		t.Error("FakeMultireturner.NamedReturn called, expected none")
	}
}

// NamedReturnCalledOnce returns true if FakeMultireturner.NamedReturn was called exactly once
func (f *FakeMultireturner) NamedReturnCalledOnce() bool {
	return len(f.NamedReturnCalls) == 1
}

// AssertNamedReturnCalledOnce calls t.Error if FakeMultireturner.NamedReturn was not called exactly once
func (f *FakeMultireturner) AssertNamedReturnCalledOnce(t MultireturnerTestingT) {
	t.Helper()
	if len(f.NamedReturnCalls) != 1 {
		t.Errorf("FakeMultireturner.NamedReturn called %d times, expected 1", len(f.NamedReturnCalls))
	}
}

// NamedReturnCalledN returns true if FakeMultireturner.NamedReturn was called at least n times
func (f *FakeMultireturner) NamedReturnCalledN(n int) bool {
	return len(f.NamedReturnCalls) >= n
}

// AssertNamedReturnCalledN calls t.Error if FakeMultireturner.NamedReturn was called less than n times
func (f *FakeMultireturner) AssertNamedReturnCalledN(t MultireturnerTestingT, n int) {
	t.Helper()
	if len(f.NamedReturnCalls) < n {
		t.Errorf("FakeMultireturner.NamedReturn called %d times, expected >= %d", len(f.NamedReturnCalls), n)
	}
}
