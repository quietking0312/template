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

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var isShowVersion bool

func init() {
	rootCmd.Flags().BoolVarP(&isShowVersion, "version", "v", false, "打印版本信息")
}
