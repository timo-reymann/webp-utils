package execute

import (
	"os/exec"
	"strings"
)

// BuildCommandsForFileList generates a list with command wrappers for cli execution
func BuildCommandsForFileList(executable string, fileNames []string, configuration map[string]interface{}) ([]Command, error) {
	executablePath, err := getExecutablePath(executable)
	if err != nil {
		return nil, err
	}

	commands := make([]Command, len(fileNames))
	for i, fileName := range fileNames {
		// Create templater for file
		templater, err := newArgumentTemplater(configuration, map[string]interface{}{
			"source_file":      fileName,
			"source_file_name": strings.Split(fileName, ".")[0],
		})

		if err != nil {
			return nil, err
		}

		// Create args for file
		args := make([]Argument, 0)
		for key, value := range configuration {
			args = append(args, *newArgument(key, value))
		}

		// Bundle and create command structure
		commands[i] = Command{
			Binary:       executablePath,
			Filename:     fileName,
			Args:         args,
			ArgTemplater: templater,
		}
	}

	return commands, nil
}

func getExecutablePath(executableName string) (string, error) {
	return exec.LookPath(executableName)
}
