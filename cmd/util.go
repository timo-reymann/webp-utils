package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func abort(cmd *cobra.Command, prefix string, err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprint(cmd.ErrOrStderr(), prefix+": "+err.Error())
	os.Exit(2)
}
