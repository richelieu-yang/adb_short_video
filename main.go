package main

import (
	"fmt"
	"os"

	"github.com/richelieu-yang/chimera/v3/src/command/cobraKit"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobraKit.NewSimpleCommand("sv", "自动刷短视频的工具（通过adb）.", "", func(cmd *cobra.Command, args []string) {
		fmt.Println("logic...")
	})

	versionCmd := cobraKit.NewSimpleCommand("version", "Print the version number of newApp.", "", func(cmd *cobra.Command, args []string) {
		fmt.Println("version: 0.0.1")
	})

	// 添加子命令
	rootCmd.AddCommand(versionCmd)

	os.Args = []string{"sv", "version", "yjs", "ylx", "wmq"}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
