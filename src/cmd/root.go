package cmd

import (
	"time"

	"github.com/richelieu-yang/chimera/v3/src/android/adbKit"
	"github.com/richelieu-yang/chimera/v3/src/command/cobraKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = cobraKit.NewSimpleCommand("sv", "自动刷短视频的工具（通过adb）.", "", rootRun)

var (
	// addr 必选的本地标识
	addr string

	// clean 本地标识
	clean bool

	// verbose 本地标识
	verbose bool

	minSec int

	maxSec int
)

func init() {
	// 主命令的本地标识 addr
	rootCmd.Flags().StringVarP(&addr, "addr", "", "", "ADB address.")
	// 将本地标识 addr 设置为必选项
	if err := rootCmd.MarkFlagRequired("addr"); err != nil {
		panic(err)
	}

	rootCmd.Flags().BoolVarP(&clean, "clean", "", false, "Clean before connecting?")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "", false, "More console output?")

	rootCmd.Flags().IntVarP(&minSec, "min", "", 30, "min duration")
	rootCmd.Flags().IntVarP(&maxSec, "max", "", 60, "max duration")
}

func Execute() error {
	return rootCmd.Execute()
}

func rootRun(cmd *cobra.Command, args []string) {
	console.Infof("addr: %s", addr)
	console.Infof("clean: %t", clean)
	console.Infof("verbose: %t", verbose)
	console.Infof("minSec: %d", minSec)
	console.Infof("maxSec: %d", maxSec)

	ins := adbKit.NewInstance("127.0.0.1:5555", true, true)
	if err := ins.Initialize(); err != nil {
		console.Fatalf("fail to initialize: %+v", err)
	}

	count := 0
	for {
		count++

		sec := randomKit.Int(minSec, maxSec+1)
		d := time.Second * time.Duration(sec)
		console.Info("Sleep starts.", zap.Int("count", count), zap.String("duration", d.String()))
		time.Sleep(d)
		console.Info("Sleep ends.", zap.Int("count", count), zap.String("duration", d.String()))

		if err := ins.SwipeLikeAHumanBeing(820, 1200, 880, 500, 450, 20, 50); err != nil {
			console.Error("Fail to swipe.", zap.Error(err))
			continue
		}
		console.Info("Manager to swipe.")
	}
}
