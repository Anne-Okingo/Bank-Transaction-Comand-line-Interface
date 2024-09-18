package functions

import (
	"testing"
)

func TestDebit(t *testing.T) {
	input1, input2 := 1000, 100
	want := 1100
	got := Debit(input1, input2)

	if want != got {
		t.Errorf("Test Debit Failled : want %d, got %d", want, got)
	}
}
