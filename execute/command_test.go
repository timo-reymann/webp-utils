package execute

import "testing"

func TestCommand_BuildArgs(t *testing.T) {
	templater, _ := newArgumentTemplater(map[string]interface{}{
		"special_snowflake": "yes",
	}, map[string]interface{}{})
	cmd := Command{
		Binary: "echo",
		Args: []Argument{
			{
				Name:  "test",
				Value: "value",
			},
		},
		Filename:     "",
		ArgTemplater: templater,
	}

	args, _ := cmd.BuildArgs()

	for _, arg := range args {
		if arg == "special_snowflake" {
			t.Error("Expected context variables not to appear in effective parameters")
		}
	}
}

func TestCommand_RunWithCombinedOutput(t *testing.T) {
	templater, _ := newArgumentTemplater(map[string]interface{}{
		"special_snowflake": "yes",
	}, map[string]interface{}{})
	cmd := Command{
		Binary: "echo",
		Args: []Argument{
			{
				Name:  "test",
				Value: "value",
			},
		},
		Filename:     "",
		ArgTemplater: templater,
	}

	_, err := cmd.RunWithCombinedOutput()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestCommand_String(t *testing.T) {
	templater, _ := newArgumentTemplater(map[string]interface{}{}, map[string]interface{}{})
	cmd := Command{
		Binary: "echo",
		Args: []Argument{
			{
				Name:  "test",
				Value: "value",
			},
		},
		Filename:     "",
		ArgTemplater: templater,
	}

	expected := `Command{binary=echo, args=[-test value ]}`

	if cmd.String() != expected {
		t.Errorf("Expected %s, but got %s", expected, cmd.String())
	}
}

func TestBatch(t *testing.T) {
	templater, _ := newArgumentTemplater(map[string]interface{}{}, map[string]interface{}{})

	commands := []Command{
		{
			Binary: "echo",
			Args: []Argument{
			},
			Filename:     "",
			ArgTemplater: templater,
				},
		{
			Binary: "echo",
			Args: []Argument{
			},
			Filename:     "",
			ArgTemplater: templater,
				},
	}
	_, err := Batch(commands)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}