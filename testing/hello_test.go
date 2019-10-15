package hello

import "testing"

func TestHello(t *testing.T) {
	if v := Hello(); v != "hello" {
		t.Errorf("Expected 'hello', but got '%s'", v)
	}
}

//* the test should always be in the same package as the code it's tetsting
