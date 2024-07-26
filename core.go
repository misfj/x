package main

import (
	"context"
	"coredx/client"
	"coredx/config"
	"coredx/core"
	"coredx/db"
	"coredx/exe"
	"coredx/log"
	"coredx/store/cache"
	"coredx/store/ipfs"
	"coredx/store/minio"
	"fmt"
	"github.com/kardianos/osext"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

// import (
//
//	"context"
//	"coredx/client"
//	"coredx/config"
//	"coredx/core"
//	"coredx/db"
//	"coredx/log"
//	"coredx/store/cache"
//	"coredx/store/minio"
//	"coredx/utils"
//	"fmt"
//	"os"
//	"os/signal"
//	"syscall"
//
//	"github.com/spf13/cobra"
//
// )
//
// var configFile string
// var nodeMode string
// // 添加 init 子命令
//
//	var initCmd = &cobra.Command{
//		Use:   "init",
//		Short: "Initialize the core service",
//		Long:  `generate publicKey and privateKey,please protect your  privateKey`,
//		Run: func(cmd *cobra.Command, args []string) {
//		},
//	}
//
//	var coredCmd = &cobra.Command{
//		Use:   "cored",
//		Short: "cored is the core service of super nodes and cloud platforms",
//		Run: func(cmd *cobra.Command, args []string) {
//
//			utils.CheckAndCreateFiles("public.pub", "private.key")
//			//初始化日志
//			fmt.Println("begin")
//			config.Init(configFile)
//			fmt.Println("end")
//			err := log.Init(config.LoggingConfig())
//			if err != nil {
//				fmt.Println(err)
//				os.Exit(1)
//			}
//			fmt.Println("end2")
//			//初始化数据库
//			db.Init(config.DbConfig())
//			fmt.Println("end3")
//			//初始化备份服务器
//			//err = ipfs.Init(config.BackupConfig())
//			//if err != nil {
//			//	log.Error(err)
//			//}
//			//初始化平台缓存服务器
//			cache.Init(config.CacheConfig())
//			fmt.Println("end3")
//			//初始化存储服务器
//			fmt.Println(config.StoreConfig())
//			minio.Init(config.StoreConfig())
//			fmt.Println("end4")
//			client.Init()
//
//			server := core.New()
//			err = server.Init()
//			if err != nil {
//				log.Error(err)
//				os.Exit(1)
//			}
//
//			ctx, cancel := context.WithCancel(context.Background())
//
//			server.Run(ctx)
//
//			// 处理SIGINT和SIGTERM信号
//			interrupt := make(chan os.Signal, 1)
//			signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
//
//			sig := <-interrupt
//			log.Info("recv signal:", sig)
//			cancel()
//			server.Close()
//			server.Wait()
//		},
//	}
//
//	func init() {
//		coredCmd.Flags().StringVarP(&configFile, "config", "c", "data.ini", "config file path")
//		//coredCmd.Flags().StringVarP(&nodeMode, "cloud", "m", "data.ini", "choose node mode")
//		coredCmd.AddCommand(initCmd)
//	}
//
//	func main() {
//		if err := coredCmd.Execute(); err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//	}
func init() {
	config.Init("config.yaml", "private.key", "public.pub")
	fmt.Println("------------------------------------------------")
	dir := ""
	var err error
	//加载代理程序
	if dir, err = osext.ExecutableFolder(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(config.Config.Watertwo)
	//configJson := filepath.Join(dir, "exe", "config.json")
	//fmt.Printf("waterProgram config  file:%s\n", configJson)
	// args := []string{":9090", configJson}
	//args := make([]string, 0, 2)
	args := strings.Split(config.Config.Watertwo.Args, " ")
	exe.Init(filepath.Join(dir, "exe", config.Config.Watertwo.Name), args)
}
func main() {

	log.Init(config.GetLogging())
	cache.Init(config.GetCache())
	minio.Init(config.GetStore())
	ipfs.Init(config.GetBack())
	db.Init(config.GetDb())
	client.Init()
	server := core.New()
	err := server.Init()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	//ctx 带两个值 plat代表平台地址(带了表示作为超级节点,不带表示作为平台) ws:(带参数表示作为平台,另外启动端口)
	/*
		平台:node-platform == "", ws !=""
		超级节点:node-platform != "", ws ==""
	*/
	var ctx context.Context

	ctx = context.WithValue(context.Background(), "node-platform", config.GetNode().NodePlatform)
	ctx = context.WithValue(ctx, "ws", config.GetWs().ListenAddress)

	newCtx, cancel := context.WithCancel(ctx)
	server.Run(newCtx)
	// 处理SIGINT和SIGTERM信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-interrupt
	log.Info("recv signal:", sig)
	cancel()
	server.Close()
	server.Wait()

}
