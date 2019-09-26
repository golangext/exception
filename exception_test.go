package exception

import (
	"io"
	"testing"
)

func TestCatchAll(t *testing.T) {
	defer All.Catch(func(e Exception) {
		t.Logf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	NotImplemented.ThrowNew(nil, "didn't feel like it")
}

func TestCatchGenericIO(t *testing.T) {
	defer IO.Catch(func(e Exception) {
		t.Logf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	panic(io.EOF)
}

func TestRethrow(t *testing.T) {
	defer All.Catch(func(e Exception) {
		t.Logf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	defer NullPointer.Catch(func(e Exception) {
		t.Fatalf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	EOF.ThrowNew(nil, "An EOF happened")
}

func TestGenericPanic(t *testing.T) {
	defer All.Catch(func(e Exception) {
		t.Fatalf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	defer NullPointer.Catch(func(e Exception) {
		t.Fatalf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	panic(3)
}

var nullOrIO = FilterOf(NullPointer, IO)

func TestFilterNegative(t *testing.T) {
	defer Panic.Catch(func(e Exception) {
		t.Logf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	defer nullOrIO.Catch(func(e Exception) {
		t.Fatalf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	panic(3)
}

func TestFilterPositive1(t *testing.T) {
	defer nullOrIO.Catch(func(e Exception) {
		t.Logf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	defer Panic.Catch(func(e Exception) {
		t.Fatalf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	EOF.ThrowNew(nil, "An EOF happened")
}

func TestFilterPositive2(t *testing.T) {
	defer nullOrIO.Catch(func(e Exception) {
		t.Logf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	defer Panic.Catch(func(e Exception) {
		t.Fatalf("Caught IO: %v type: %v", e.Error(), e.Type())
	})

	panic(io.EOF)
}
