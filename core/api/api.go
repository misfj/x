package api

import (
	"context"
	"coredx/config"
	"coredx/core/api/middware"
	"coredx/core/node"
	"coredx/log"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt/rsa"
)

type Server struct {
	config    *config.API
	eng       *gin.Engine
	httpServe *http.Server
	ctx       context.Context
	ID        string
	Public    string
	Private   string
}

func e() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	return gin.New()
}
func New(config *config.API) *Server {
	//todo 将节点公私密钥写入数据库.如果作为超级节点，将自己的公钥上报给女娲云链平台.在以后的业务用的着.目前放在内存
	pair, err := rsa.GenerateRsaKeyBase64(2048)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	idMd5 := md5.Sum([]byte(pair.PrivateKey))
	e := e()
	e.MaxMultipartMemory = 1 << 30
	return &Server{
		config:  config,
		eng:     e,
		ID:      hex.EncodeToString(idMd5[:]),
		Public:  pair.PublicKey,
		Private: pair.PrivateKey,
	}
}

func (s *Server) Init() error {
	node.NewGlobalNodes()
	s.initRoute()

	return nil
}

func (s *Server) Name() string {
	return "APIServer"
}

func (s *Server) Startup(ctx context.Context) (err error) {
	s.ctx = ctx
	middware.Init(s.ctx)

	log.Info("APIServer startup...")
	log.Debugf("api address:%s", s.config.ListenAddress)
	if !s.config.EnableTLS {
		s.httpServe = &http.Server{
			Addr:    s.config.ListenAddress,
			Handler: s.eng,
		}
		//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//	err := v.RegisterValidation("TID", req.TID)
		//	if err != nil {
		//		return err
		//	}
		//}
		if err := s.httpServe.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err)
		}
	} else {
		//if err = s.eng.RunTLS(s.config.ListenAddress,
		//	s.config.CertFile,
		//	s.config.KeyFile); err != nil {
		//
		//	log.Error(err)
		//	return err
		//}
	}
	return nil
}

func (s *Server) Close() error {
	log.Info("APIServer close..")

	return s.httpServe.Shutdown(s.ctx)
}
