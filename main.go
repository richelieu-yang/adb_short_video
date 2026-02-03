package main

import (
	"os"

	"github.com/richelieu-yang/adb_short_video/src/cmd"
)

func main() {
	os.Args = []string{"sv", "version"}
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
