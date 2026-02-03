package cmd

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/command/cobraKit"
	"github.com/spf13/cobra"
)

var rootCmd = cobraKit.NewSimpleCommand("sv", "自动刷短视频的工具（通过adb）.", "", rootRun)

func Execute() error {
	return rootCmd.Execute()
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println("logic...")
}
