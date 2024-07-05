package api

import (
	"coredx/core/api/middware"
	"coredx/core/api/service"
	"coredx/core/api/ws"
)

func (s *Server) initRoute() {
	s.eng.GET("/api/v1/node/1", ws.NodeService)
	s.eng.Use(middware.Cors())
	x := s.eng.Group("/v1/user")
	x.POST("register", service.Register)

	//userNoAuthGroup.POST("/v1/user/login", service.Login)
	userGroup := s.eng.Group("/v1/user")
	userGroup.Use(middware.Access())
	userGroup.POST("/modify", service.Modify)
	userGroup.POST("/delete", service.Delete)
	userGroup.POST("/logout", service.Logout)
	userGroup.POST("/test", service.Test)
	//userGroup := s.eng.Group("/v1/user")
	//userGroup.POST("/login")
	//s.eng.POST("api/login", service.LoginService)
	//authGroup := s.eng.Group("/api/v1")
	//authGroup.Use(middware.Limit(), middware.Token())
	////限流测试接口middware
	//authGroup.POST("data/rate", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"code": 200})
	//})
	////s.eng.Use(handler.MiddleWare())
	//authGroup.GET("data/user/info", handler.GetDrUserInfo())
	//authGroup.POST("data/upload", service.UploadService)
	//authGroup.POST("data/delete", handler.DeleteData())
	//authGroup.POST("data/download", handler.DownloadData())
	//authGroup.POST("data/user/create", service.DataUserRegisterService)
	//authGroup.POST("data/user/modify", handler.DrUserModify())
	//authGroup.POST("data/encrypt", handler.Encrypt())
	//authGroup.POST("data/decrypt", handler.Decrypt())
	//authGroup.POST("data/water", handler.Water())
	//authGroup.POST("back", handler.BackUp())
	//authGroup.POST("data/user/account/info", handler.DrAccountInfo())
	//authGroup.POST("data/user/account/sync", handler.DrAccountSync())

}
