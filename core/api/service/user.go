package service

import (
	"coredx/core/api/handler/user"
	"coredx/protocol"
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(ctx *gin.Context, code int, msg string, data string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
func Register(gctx *gin.Context) {
	var req *protocol.RegisterRequest
	err := gctx.ShouldBindJSON(&req)
	if err != nil {
		response(gctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	err = req.Validate()
	if err != nil {
		response(gctx, http.StatusNotFound, err.Error(), "")
		return
	}
	err = user.Register(req)
	if err != nil {
		response(gctx, http.StatusBadRequest, err.Error(), "")
		return
	}
	response(gctx, http.StatusOK, "success", "")
}
