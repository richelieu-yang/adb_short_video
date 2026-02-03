package logic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionRun(cmd *cobra.Command, args []string) {
	fmt.Println("version: 0.0.1")
}
