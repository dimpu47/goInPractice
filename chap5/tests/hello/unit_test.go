package hello

import (
	"testing"
)

func Testhello(t *testing.T) {
	if v := hello(); v != "hello" {
		t.Errorf("Expected 'hello', but got %s", v)
	}
}
