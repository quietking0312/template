package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Server",
	Short: "",
	Run: func(cmd *cobra.Command, arg []string) {
		if isShowVersion {
			fmt.Printf("UTC build time: %s\n", buildTime)
			fmt.Printf("Build from version: %s\n", version)
			fmt.Printf("Commit: %s\n", commit)
		}
	},
}

func Execute(bTime, v, c string) {
	buildTime = bTime
	version = v
	commit = c
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// 打包时间， 版本号
var buildTime, version, commit string

var isShowVersion bool

func init() {
	rootCmd.Flags().BoolVarP(&isShowVersion, "version", "v", false, "打印版本信息")
}
