package tests

import (
	"testing"
)

func TestDiv(t *testing.T) {
	if ans := Div(5, 3); ans != 1 {
		t.Errorf("5 / 3 expected be 1, but %d got", ans)
	}
}

func Div(a, b int64) int64 {
	return a / b
}
