package main

import (
	"os"

	"github.com/richelieu-yang/adb_short_video/src/cmd"
)

func main() {
	os.Args = []string{"sv" /*"--addr", "127.0.0.1:5555",*/, "--clean", "--verbose"}
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
