package cmd

func init() {
	rootCmd.AddCommand(createCobraCommand("cwebp"))
}
