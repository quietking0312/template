package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/core/config"
	"server/core/router"
	"server/core/utils/mylog"
)

var server = &cobra.Command{
	Use:   "server",
	Short: "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.InitViperConfigFile(serverConfigFile, version); err != nil {
			fmt.Println("配置文件加载失败：", err)
			return err
		}
		mylog.Init() // 初始化日志模块
		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {

		addr := fmt.Sprintf(":%d", config.GetConfig().Server.Port)
		lis, err := router.GetListen(addr)
		if err != nil {
			return err
		}
		err = router.Run(lis)
		return nil
	},
}

var serverConfigFile string

func init() {
	rootCmd.AddCommand(server)
	server.Flags().StringVarP(&serverConfigFile, "config", "c", "server", "直接加载的配置文件")
}
