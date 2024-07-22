package cache

import (
	"coredx/config"
	"coredx/log"
	"fmt"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var cli *minio.Client
var expire time.Duration
var err error

// Init 初始化女娲平台缓存服务器
func Init(cfg *config.Cache) {
	cli, err = minio.New(cfg.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSsl,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	expire = time.Duration(cfg.Expire) * time.Hour * 24
	log.Debug("initial cache store success")
}
