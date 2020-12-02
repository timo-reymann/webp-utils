package execute

import "testing"

func TestBuildCommandsForFileList(t *testing.T) {

	commands, err := BuildCommandsForFileList("echo", []string{"foo.txt", "bar.txt"}, map[string]interface{}{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	cmd := commands[0]

	if cmd.Binary == "echo" {
		t.Error("Resolution of binary does not work")
	}

	_, err = BuildCommandsForFileList("not-good-99", []string{"foo.txt", "bar.txt"}, map[string]interface{}{})
	if err == nil {
		t.Errorf("Invalid binary not detected")
	}
}
