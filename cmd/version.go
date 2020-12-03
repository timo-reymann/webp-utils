package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timo-reymann/webp-utils/pkg/buildinfo"
	"os"
	"text/tabwriter"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version info",
		Run: func(cmd *cobra.Command, args []string) {
			println("The following build info is available")
			writer := tabwriter.NewWriter(os.Stdout, 10, 1, 10, byte(' '), tabwriter.TabIndent)
			fmt.Fprintf(writer, "GitSha\t%s\n", buildinfo.GitSha)
			fmt.Fprintf(writer, "Version\t%s\n", buildinfo.Version)
			writer.Flush()
		},
	}

	rootCmd.AddCommand(versionCmd)
}
