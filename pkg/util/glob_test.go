package util

import "testing"

func TestFileGlobMatches(t *testing.T) {
	// Test no files found
	files, err := FileGlobMatches("testdata/*.jpeg")
	if len(files) > 0 {
		t.Error("Should find no files for non existent glob")
	}
	if err == nil || err.Error() != "glob doesnt match any file" {
		t.Errorf("Expected error, but got %v", err)
	}

	// Test existing
	files, err = FileGlobMatches("testdata/*.png")
	if len(files) != 1 {
		t.Errorf("Should find one file, but found %d", len(files))
	}
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
