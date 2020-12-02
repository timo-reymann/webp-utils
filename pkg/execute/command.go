package execute

import (
	"fmt"
	"os/exec"
)

// Command represents all information necessary to execute a webp cli tool
type Command struct {
	Binary       string
	Args         []Argument
	Filename     string
	ArgTemplater *ArgumentTemplater
}

// String generates a debug(able) presentation
func (c *Command) String() string {
	args, _ := c.BuildArgs()

	return fmt.Sprintf("Command{binary=%s, args=%v}", c.Binary, args)
}

// BuildArgs constructs the final arguments used for exec.Command
func (c *Command) BuildArgs() ([]string, error) {
	args := make([]string, 0)

	for _, arg := range c.Args {
		// Ignore parameters only used for templating
		if c.ArgTemplater.IsExtraParameter(arg.Name) {
			continue
		}

		// Append all args for argument, split into name and value
		extraArg, err := arg.BuildArg(c.ArgTemplater)
		if err != nil {
			return nil, err
		}
		args = append(args, extraArg...)
	}

	// Append filename as latest parameter
	args = append(args, c.Filename)

	return args, nil
}

// RunWithCombinedOutput executes the command without separating stdout and stderr
func (c *Command) RunWithCombinedOutput() (string, error) {
	args, err := c.BuildArgs()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(c.Binary, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), err
}

// Batch executes multiple commands till a error occurs, if an error occurs the stdout is
// also included
func Batch(commands []Command) (string, error) {
	for _, line := range commands {
		println(">> Execute " + line.String())
		out, err := line.RunWithCombinedOutput()
		if err != nil {
			return out, err
		}
	}

	return "", nil
}
