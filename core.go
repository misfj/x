package main

import (
	"context"
	"coredx/client"
	"coredx/config"
	"coredx/core"
	"coredx/db"
	"coredx/log"
	"coredx/metrics/process"
	"coredx/store/cache"
	"coredx/store/minio"
	"coredx/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var configFile string
var nodeMode string

// 添加 init 子命令
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the core service",
	Long:  `generate publicKey and privateKey,please protect your  privateKey`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
var coredCmd = &cobra.Command{
	Use:   "cored",
	Short: "cored is the core service of super nodes and cloud platforms",
	Run: func(cmd *cobra.Command, args []string) {

		utils.CheckAndCreateFiles("public.pub", "private.key")
		//初始化日志
		config.Init(configFile)
		err := log.Init(config.LoggingConfig())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//初始化数据库
		db.Init(config.DbConfig())
		//初始化备份服务器
		//err = ipfs.Init(config.BackupConfig())
		//if err != nil {
		//	log.Error(err)
		//}
		//初始化平台缓存服务器
		cache.Init(config.CacheConfig())
		//初始化存储服务器
		minio.Init(config.StoreConfig())
		client.Init()

		server := core.New()
		err = server.Init()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		ctx, cancel := context.WithCancel(context.Background())

		server.Run(ctx)

		// 处理SIGINT和SIGTERM信号
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

		sig := <-interrupt
		log.Info("recv signal:", sig)
		cancel()
		server.Close()
		server.Wait()
	},
}

func init() {
	coredCmd.Flags().StringVarP(&configFile, "config", "c", "data.ini", "config file path")
	//coredCmd.Flags().StringVarP(&nodeMode, "cloud", "m", "data.ini", "choose node mode")
	coredCmd.AddCommand(initCmd)
}

func main() {
	//if err := coredCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//fmt.Println(system.GetCpuPercent())
	//fmt.Println(system.GetMemPercent())
	//fmt.Println(system.GetDiskPercent())
	//
	p := process.GetProcessInfo()

	for _, s := range p {
		fmt.Println(s)
	}
	//
	//err := process.KillProcess(15132)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//s := process.GetProcessNetInfo()
	//
	//for _, u := range s {
	//	fmt.Println(u)
	//}
}
