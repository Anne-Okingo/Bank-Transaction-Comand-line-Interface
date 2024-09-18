package functions

import (
	"testing"
)

func TestCredit(t *testing.T) {
	input1, input2 := 1000, 100
	want := 900
	got := Credit(input1, input2)

	if want != got {
		t.Errorf("Test Debit Failled : want %d, got %d", want, got)
	}
}
