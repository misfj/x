package main

import (
	"context"
	"coredx/client"
	"coredx/config"
	"coredx/core"
	"coredx/db"
	"coredx/global"
	"coredx/log"
	"coredx/store/cache"
	"coredx/store/ipfs"
	"coredx/store/minio"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	config.Init("config.yaml", "private.key", "public.pub")
	//dir := ""
	//var err error
	////加载代理程序
	//if dir, err = osext.ExecutableFolder(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//args := strings.Split(config.Config.Watertwo.Args, " ")
	//exe.Go(filepath.Join(dir, "exe", config.Config.Watertwo.Name), args)
}

// 生成swagger文档
// cd  core/api/  && swag init  --parseDependency  -g  .\service\app.go  docs
func main() {
	log.Init(config.GetLogging())
	defer log.Sync()
	cache.Init(config.GetCache())
	minio.Init(config.GetStore())
	ipfs.Init(config.GetBack())
	db.Init(config.GetDb())
	//fmt.Println(model.AppInfoQuery.FindByAppName("数据资产"))
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfbmFtZSI6IuS9oOWlveWwj-iJviIsImlzcyI6ImNvcmVkIiwiZXhwIjoxNzMwMjc5MDA0LCJpYXQiOjE3MjI1MDMwMDR9.uGa4MLd_sk8qNejQenkJbvyO6ObruvXg2AZ9hFnzoXk"
	//parseToken, err2 := middleware.ParseToken(token, config.GetJwt().Secret)
	//
	//fmt.Println(parseToken.ExpiresAt, err2)
	global.InitGlobal(20)
	go global.Global.Start()
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

	ctx = context.WithValue(
		context.Background(),
		"node-platform",
		config.GetNode().NodePlatform)

	ctx = context.WithValue(
		ctx,
		"ws",
		config.GetWs().ListenAddress)

	newCtx, cancel := context.WithCancel(ctx)
	server.Run(newCtx)

	// 处理SIGINT和SIGTERM信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-interrupt
	log.Info("recv signal:", sig)

	cancel()
	server.Close()
	server.Wait()

}
