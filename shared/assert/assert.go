package assert

import (
	"fmt"
	"testing"
)

func True(t testing.TB, condition bool, msg ...string) {
	t.Helper()
	boolEqual(t, true, condition, msg...)
}

func False(t testing.TB, condition bool, msg ...string) {
	t.Helper()
	boolEqual(t, false, condition, msg...)
}

func boolEqual(t testing.TB, wanted, got bool, msg ...string) {
	t.Helper()
	if wanted != got {
		if len(msg) == 0 {
			msg = []string{fmt.Sprintf("should be %t", wanted)}
		}
		t.Errorf(msg[0])
	}
}

func IntEqual(t testing.TB, wanted, got int, msg ...string) {
	t.Helper()
	if wanted != got {
		if len(msg) == 0 {
			msg = []string{fmt.Sprintf("wanted %d but got %d", wanted, got)}
		}
		t.Errorf(msg[0])
	}
}

// SliceIntEqual
// note: for testing, nil-slice and empty slice are equal
func SliceIntEqual(t testing.TB, wanted, got []int) {
	t.Helper()
	if len(wanted) != len(got) {
		t.Errorf("wanted %v but got %v", wanted, got)
		return
	}

	for i := range wanted {
		IntEqual(t, wanted[i], got[i], fmt.Sprintf("wanted %v but got %v", wanted, got))
	}
}