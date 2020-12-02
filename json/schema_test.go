package json

import "testing"

func TestValidate(t *testing.T) {
	// Valid
	errors := Validate([]byte(`{"o":"test.webp"}`), "cwebp")
	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	// Invalid
	errors = Validate([]byte(`{}`), "cwebp")
	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, but got %v", errors)
	}

	// Non existent schema
	errors = Validate([]byte(`{}`), "unknown")
	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, but got %v", errors)
	}
}
