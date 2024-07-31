package service

import (
	_ "coredx/core/api/v1/app/request"
	"coredx/core/api/v1/app/response"
	_ "coredx/core/api/v1/app/response"
	"github.com/gin-gonic/gin"
	"time"
)

// AppLogin
// @Tags      业务应用登录
// @Summary   业务应用登录
// @accept    application/json
// @Produce   application/json
// @Param data body request.AppLoginRequest true  "请求参数"
// @Success   200   {object}  response.Response{data=string}  "业务应用登录"
// @Router    /v1/app/login [post]
func AppLogin(c *gin.Context) {
	time.Sleep(3 * time.Second)
	c.JSON(200, response.Response{
		Code: 200,
		Data: "ok",
		Msg:  "ok",
	})
}
