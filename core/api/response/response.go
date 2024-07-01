package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(code int, msg string, data string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

const (
	Code         = "code"
	Msg          = "msg"
	Data         = "data"
	ErrProtocol  = "协议错误"
	ErrParameter = "参数错误"
	Success      = "success"
)
