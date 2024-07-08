package api

import (
	"coredx/core/api/middware"
	"coredx/core/api/service"
	"coredx/core/api/ws"
)

func (s *Server) initRoute() {
	//绑定全局跨域
	s.eng.Use(middware.Cors())
	s.eng.GET("/api/v1/node/1", ws.NodeService)
	x := s.eng.Group("/v1/user")
	x.POST("register", service.Register)
	x.POST("login", service.Login)

	//用户组
	userGroup := s.eng.Group("/v1/user")
	userGroup.Use(middware.Access())
	userGroup.POST("/modify", service.Modify)
	userGroup.POST("/delete", service.Delete)
	//这个API只能给用管理员使用,以后考虑权限控制
	userGroup.POST("/list", service.List)
	userGroup.POST("/get", service.Get)
	userGroup.POST("/space/info", service.SpaceInfo)
	userGroup.POST("/space/expand", service.SpaceExpand)
	userGroup.POST("/upgrade", service.UpGrade)
	userGroup.POST("/test", service.Test)

	//管理组
	managerGroup := s.eng.Group("/v1/manager")

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
