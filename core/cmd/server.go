package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/core/config"
	"server/core/dao"
	"server/core/logic"
	"server/core/router"
	"server/core/utils/mylog"
	"time"
)

var server = &cobra.Command{
	Use:   "server",
	Short: "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.InitViperConfigFile(serverConfigFile, version); err != nil {
			fmt.Println("配置文件加载失败：", err)
			return err
		}

		mylog.Init()                // 初始化日志模块
		time.Sleep(5 * time.Second) // 使用docker 启动时， mysql 启动比server 启动慢，需要延时
		if err := dao.InitDB(); err != nil {
			fmt.Println("数据库初始化失败：", err)
			return err
		}
		if err := logic.Common.Init(); err != nil {
			fmt.Println("common 初始化失败：", err)
			return err
		}
		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {

		addr := fmt.Sprintf(":%d", config.GetConfig().Server.Port)
		lis, err := router.GetListen(addr)
		if err != nil {
			return err
		}
		// err = router.Run(lis)
		err = router.RunHttpServer(lis)
		return err
	},
}

var serverConfigFile string

func init() {
	rootCmd.AddCommand(server)
	server.Flags().StringVarP(&serverConfigFile, "config", "c", "server", "直接加载的配置文件")
}
