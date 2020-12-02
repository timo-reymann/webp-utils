package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timo-reymann/webp-utils/pkg/config"
	"github.com/timo-reymann/webp-utils/pkg/execute"
	"github.com/timo-reymann/webp-utils/pkg/util"
)

func init() {
	cpwebCmd := &cobra.Command{
		Use:   "cwebp",
		Short: "Call cwebp",
		Run: func(cmd *cobra.Command, args []string) {
			fileGlob, _ := cmd.Flags().GetString("file-glob")
			files, err := util.FileGlobMatches(fileGlob)
			abort(cmd, "Error getting files for glob", err)

			configFile, _ := cmd.Flags().GetString("config")
			configuration, err := config.Load(configFile)
			abort(cmd, "Error loading configuration file", err)

			// Build arg lines
			commands, err := execute.BuildCommandsForFileList("cwebp", files, configuration)
			abort(cmd, "Error building cli for cwebp", err)

			// Execute commands
			stdout, err := execute.Batch(commands)
			if err != nil {
				_, _ = cmd.ErrOrStderr().Write([]byte(stdout))
			}
			abort(cmd, "Error executing cli", err)
		},
	}

	cpwebCmd.Flags().String("file-glob", "*.png", "File util pattern for source files to be converted")
	cpwebCmd.Flags().String("config", "config.json", "Configuration file for command execution")

	rootCmd.AddCommand(cpwebCmd)
}
