package middleware

import (
	"bytes"
	"coredx/db/dal/model"
	"coredx/log"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 记录每次请求的请求信息和响应信息

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Debug("-----------------------进入logger-------------------")
		// 请求前
		t := time.Now()
		reqPath := c.Request.URL.Path
		//log.Debugf("请求路径:%s", reqPath)
		method := c.Request.Method
		ip := c.ClientIP()
		rawQuery := c.Request.URL.RawQuery
		// JSON和FORM表单打印请求的Body, 其他内容类型，比如文件上传不打印
		var requestBody string
		var requestBody2 string
		contentType := c.GetHeader("Content-Type")
		//log.Debugf("打印API头部:%s", contentType)
		if contentType != "" &&
			(strings.HasPrefix(contentType, "application/json") ||
				strings.HasPrefix(contentType, "application/x-www-form-urlencoded")) {
			requestBody, err := io.ReadAll(c.Request.Body)
			if err != nil {
				requestBody = []byte{}
			}
			requestBody2 = string(requestBody)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}
		log.Info(fmt.Sprintf("host:%s %s %s start", ip, method, reqPath), "query", rawQuery, "body", requestBody)
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		//将需要的数据放入ctx
		c.Set("req_url", reqPath)
		c.Set("method", method)
		c.Set("remote_ip", ip)
		c.Set("req_content", requestBody2)
		c.Writer = writer
		c.Next()
		_, exist := c.Get("jwt_error")
		if exist {
			c.Abort()
			return
		}
		c.Set("res_msg", writer.body.String())
		c.Set("res_size", writer.Size())
		latency := time.Since(t).Microseconds()
		c.Set("cost", latency)
		matchLogger(reqPath, c)
		// 请求后
		log.Info(fmt.Sprintf("host:%s %s %s end", ip, method, reqPath), "cost/us", latency)
	}
}

func matchLogger(path string, newCtx *gin.Context) (err error) {

	err = filter(path, newCtx)
	if err != nil {
		return
	}

	//为数据添加上下文
	appID, _ := newCtx.Get("app_id")
	serviceType, _ := newCtx.Get("service_type")
	reqUrl, _ := newCtx.Get("req_url")
	method, _ := newCtx.Get("method")
	remoteIP, _ := newCtx.Get("remote_ip")
	reqContent, _ := newCtx.Get("req_content")
	resMsg, _ := newCtx.Get("res_msg")
	resSize, _ := newCtx.Get("res_size")
	resErr, _ := newCtx.Get("res_err")
	cost, _ := newCtx.Get("cost")
	appIDuint64 := appID.(uint)
	serviceTypeString := serviceType.(string)
	reqUrlString := reqUrl.(string)
	methodString := method.(string)
	remoteIPString := remoteIP.(string)
	reqContentString := reqContent.(string)
	resMsgString := resMsg.(string)
	resSizeInt32 := resSize.(int)
	resErrString := resErr.(string)
	costInt64, _ := cost.(int64)
	err = model.AppLogQuery.Create(&model.AppLog{
		AppID:       int64(appIDuint64),
		ServiceType: serviceTypeString,
		ReqURL:      reqUrlString,
		Method:      methodString,
		RemoteIP:    remoteIPString,
		ReqContent:  reqContentString,
		ResMsg:      resMsgString,
		ResSize:     int32(resSizeInt32),
		ResErr:      resErrString,
		Cost:        costInt64,
	})
	if err != nil {
		log.Error(err)
	}
	return nil
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
