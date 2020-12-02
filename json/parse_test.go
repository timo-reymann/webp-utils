package json

import "testing"

func TestParse(t *testing.T) {
	_, err := Parse([]byte(`{"foo":"bar"}`))
	if err != nil {
		t.Errorf("Valid json expected to throw no error, but got %v", err)
	}
}
