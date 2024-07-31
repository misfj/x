package middleware

import (
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

// func Access() gin.HandlerFunc {
// 	return func(gctx *gin.Context) {
// 		AccessKeyHeader := gctx.Request.Header.Get("Authorization")
// 		if AccessKeyHeader == "" {
// 			response(gctx, http.StatusBadRequest, "token not exist", "")
// 			gctx.Abort()
// 			return
// 		}
// 		checkToken := strings.Split(AccessKeyHeader, " ")
// 		if len(checkToken) == 0 {
// 			response(gctx, http.StatusBadRequest, "token not exist", "")
// 			gctx.Abort()
// 			return
// 		}
// 		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
// 			response(gctx, http.StatusBadRequest, "token not exist", "")
// 			gctx.Abort()
// 			return
// 		}
// 		log.Info(checkToken[1])
// 		//数据库检验key是否过期
// 		err := v1.AuthUserTokenIsExpireByAccessKey(db.GDB, checkToken[1])
// 		if err != nil {
// 			response(gctx, http.StatusBadRequest, err.Error(), "")
// 			gctx.Abort()
// 			return
// 		}
// 		gctx.Set("AccessKey", checkToken[1])
// 		gctx.Next()
// 	}
// }
