package main

import (
	"context"
	"coredx/client"
	"coredx/config"
	"coredx/core"
	"coredx/db"
	"coredx/log"
	"coredx/store/cache"
	"coredx/store/ipfs"
	"coredx/store/minio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wumansgy/goEncrypt/rsa"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var configFile string
var nodeMode string

func x() {
	_, e := os.Stat("./private.key")
	_, e1 := os.Stat("./private.key")
	if e != nil && e1 != nil {
		e = gen()
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
	} else {
		fmt.Println("无需创建")
		if e != nil {
			e = os.Remove("./private.key")
			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			} else {
				e = os.Remove("./public.pub")
				if e != nil {
					fmt.Println(e)
					os.Exit(1)
				}
			}
			e = gen()
			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			}
		}
	}
}
func gen() error {
	fmt.Println("Initializing core service...")

	pair, err := rsa.GenerateRsaKeyBase64(2048)
	if err != nil {
		return err
	}
	sk, err := os.Create("./private.key")
	if err != nil {
		return err
	}
	pk, err := os.Create("./public.pub")
	if err != nil {
		return err
	}
	_, err = io.Copy(sk, strings.NewReader(pair.PrivateKey))
	if err != nil {
		return err
	}
	_, err = io.Copy(pk, strings.NewReader(pair.PublicKey))
	if err != nil {
		return err
	}
	return nil
}

// 添加 init 子命令
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the core service",
	Long:  `generate publicKey and privateKey,please protect your  privateKey`,
	Run: func(cmd *cobra.Command, args []string) {

		x()
		// 在这里添加初始化服务的代码
	},
}
var coredCmd = &cobra.Command{
	Use:   "cored",
	Short: "cored is the core service of super nodes and cloud platforms",
	Run: func(cmd *cobra.Command, args []string) {
		//cfgFile := "x.json"
		fmt.Println(configFile)
		fmt.Println(nodeMode)
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
		err = ipfs.Init(config.BackupConfig())
		if err != nil {
			log.Error(err)
		}
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
	if err := coredCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
