package api

import (
	_ "coredx/core/api/docs"
	"coredx/core/api/middleware"
	"coredx/core/api/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initRoute() {
	//绑定全局跨域
	s.eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.eng.Use(middleware.Cors(), middleware.Logger())

	s.eng.POST("/v1/app/login", service.AppLogin)
	// s.eng.GET("/api/v1/node/1", ws.NodeService)
	// x := s.eng.Group("/v1/user")
	// x.POST("register", service.Register)
	// x.POST("login", service.Login)

	//用户组
	// userGroup := s.eng.Group("/v1/user")
	// userGroup.Use(middware.Access())
	// userGroup.POST("/modify", service.Modify)
	// userGroup.POST("/delete", service.Delete)
	// //这个API只能给用管理员使用,以后考虑权限控制
	// userGroup.POST("/list", service.List)
	// userGroup.POST("/get", service.Get)
	// userGroup.POST("/space/info", service.SpaceInfo)
	// userGroup.POST("/space/expand", service.SpaceExpand)
	// userGroup.POST("/upgrade", service.UpGrade)
	// userGroup.POST("/test", service.Test)

}
