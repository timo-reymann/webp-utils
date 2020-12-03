package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timo-reymann/webp-utils/pkg/config"
	"github.com/timo-reymann/webp-utils/pkg/execute"
	"github.com/timo-reymann/webp-utils/pkg/util"
	"os"
)

func abort(cmd *cobra.Command, prefix string, err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%s:\n %s\n", prefix, err.Error())
	os.Exit(2)
}

func processCommand(executable string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fileGlob, _ := cmd.Flags().GetString("file-glob")
		files, err := util.FileGlobMatches(fileGlob)
		abort(cmd, "Error getting files for glob", err)

		// Load configuration
		configFile, _ := cmd.Flags().GetString("config")
		configuration, err := config.Load(configFile, executable)
		abort(cmd, "Error loading configuration file", err)

		// Build arg lines
		commands, err := execute.BuildCommandsForFileList(executable, files, configuration)
		abort(cmd, "Error building cli for cwebp", err)

		// Execute commands
		stdout, err := execute.Batch(commands)
		if err != nil {
			_, _ = cmd.ErrOrStderr().Write([]byte(stdout + "\n"))
		}
		abort(cmd, "Error executing "+executable, err)
	}
}

func addWebpDefaultFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().String("file-glob", "*.png", "File util pattern for source files to be converted")
	cmd.Flags().String("config", "config.json", "Configuration file for command execution")
	return cmd
}

func createCobraCommand(executable string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   executable,
		Short: "Call " + executable,
		Run:   processCommand(executable),
	}
	addWebpDefaultFlags(cmd)
	return cmd
}
