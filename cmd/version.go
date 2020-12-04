package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timo-reymann/webp-utils/pkg/buildinfo"
	"os"
	"runtime"
	"text/tabwriter"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version info",
		Run: func(cmd *cobra.Command, args []string) {
			println("The following build info is available")
			writer := tabwriter.NewWriter(os.Stdout, 10, 1, 10, byte(' '), tabwriter.TabIndent)
			_, _ = fmt.Fprintf(writer, "GitSha\t%s\n", buildinfo.GitSha)
			_, _ = fmt.Fprintf(writer, "Version\t%s\n", buildinfo.Version)
			_, _ = fmt.Fprintf(writer, "Build time\t%s\n", buildinfo.BuildTime)
			_, _ = fmt.Fprintf(writer, "Go version\t%s\n", runtime.Version())
			_, _ = fmt.Fprintf(writer, "OS/Arch\t%s/%s\n", runtime.GOOS, runtime.GOARCH)
			writer.Flush()
		},
	}

	rootCmd.AddCommand(versionCmd)
}
