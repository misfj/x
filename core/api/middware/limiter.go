package middware

import (
	"coredx/core/api/response"
	"coredx/log"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"strings"
)

func Limit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := strings.Split(ctx.Request.RemoteAddr, ":")
		ip := addr[0]
		//先查黑名单
		var count interface{}
		var b bool
		if blackList.find(ip) {
			response.Response(http.StatusBadRequest, "当前ip属于非法ip,10min后自动解除限制", "", ctx)
			ctx.Abort()
			return
		}
		count, b = LimitCache.Get(ip)
		//存在ip
		if b {
			expiration, t, b2 := LimitCache.GetWithExpiration(ip)
			log.Debug(expiration, t, b2)
			c := count.(int)
			//达到限流次数
			if c >= 20 {
				response.Response(http.StatusBadRequest, "请求过于频繁,10min后自动解除限制", "", ctx)
				blackList.add(ip)
				ctx.Abort()
				return
			}
			err := LimitCache.Increment(ip, 1)
			if err != nil {
				response.Response(http.StatusBadRequest, err.Error(), "", ctx)
				ctx.Abort()
				return
			}
			ctx.Next()
		} else {
			//不存在这条ip
			LimitCache.Set(ip, 0, cache.DefaultExpiration)
			err := LimitCache.Increment(ip, 1)
			if err != nil {
				response.Response(http.StatusBadRequest, err.Error(), "", ctx)
				ctx.Abort()
				return
			}
			ctx.Next()
		}
	}
}
