package service

import (
	"coredx/core/api/handler/user"
	"coredx/protocol"
	"net/http"

	"github.com/gin-gonic/gin"
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

func Login(gctx *gin.Context) {
	var req protocol.LoginRequest
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
	accessKey, err := user.Login(&req)
	if err != nil {
		response(gctx, http.StatusInternalServerError, err.Error(), "")
		return
	}
	response(gctx, http.StatusOK, "success", accessKey)
}
func Test(gctx *gin.Context) {
	response(gctx, http.StatusOK, "success", "")
}

func Modify(gctx *gin.Context) {
	var req protocol.ModifyRequest
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
	accessKeyInterface, exist := gctx.Get("AccessKey")
	if !exist {
		response(gctx, http.StatusInternalServerError, "internal error", "")
		return
	}
	accessKey, _ := accessKeyInterface.(string)
	err = user.Modify(&req, accessKey)
	if err != nil {
		response(gctx, http.StatusNotFound, err.Error(), "")
		return
	}
	response(gctx, http.StatusOK, "success", "")
	return
}
func Delete(gctx *gin.Context) {
	accessKey, exist := gctx.Get("AccessKey")
	if !exist {
		response(gctx, http.StatusInternalServerError, "internal error", "")
		return
	}
	err := user.Delete(accessKey.(string))
	if err != nil {
		response(gctx, http.StatusInternalServerError, err.Error(), "")
	}
	response(gctx, http.StatusOK, "success", "")
}

func List(gctx *gin.Context) {

}
