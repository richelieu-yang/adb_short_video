package cmd

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/command/cobraKit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = cobraKit.NewSimpleCommand("version", "Print the version number of newApp.", "", versionRun)

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println("version: 0.0.1")
}
