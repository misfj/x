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
		log.Debug("-----------------------进入logger-------------------")
		// 请求前
		t := time.Now()
		reqPath := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()
		rawQuery := c.Request.URL.RawQuery
		// JSON和FORM表单打印请求的Body, 其他内容类型，比如文件上传不打印
		var requestBody string
		var requestBody2 string
		contentType := c.GetHeader("Content-Type")
		if contentType != "" &&
			(strings.HasPrefix(contentType, "application/json") ||
				strings.HasPrefix(contentType, "application/x-www-form-urlencoded")) {
			requestBody, err := io.ReadAll(c.Request.Body)
			if err != nil {
				requestBody = []byte{}
			}
			log.Debugf("请求体:%s", requestBody)
			requestBody2 = string(requestBody)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}
		log.Info(fmt.Sprintf("host:%s %s %s start", ip, method, reqPath), "query", rawQuery, "body", requestBody)
		log.Debugf("请求体:%s", requestBody2)
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		//将需要的数据放入ctx
		c.Set("req_url", reqPath)
		c.Set("method", method)
		c.Set("remote_ip", ip)
		c.Set("req_content", requestBody2)
		//newCtx := context.WithValue(c, "", "")
		c.Writer = writer
		c.Next()
		log.Debug("-----------------------离开logger-------------------")
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
	log.Debugf("进入path:%s", path)
	switch path {
	case "/v1/app/register":
		//newCtx 有AppName
		if appName, ok := newCtx.Value("AppName").(string); ok {
			app, err := model.AppInfoQuery.FindByAppName(appName)
			if err != nil {
				newCtx.Set("res_err", err.Error())
				newCtx.Set("app_id", 0)
				newCtx.Set("service_type", "error")
			} else {
				newCtx.Set("res_err", "")
				newCtx.Set("app_id", app.ID)
				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
				newCtx.Set("service_type", "平台注册业务应用服务")
			}
			log.Debugf("取出名称:%s", appName)
		}
	case "/v1/app/account/bind":
		if appName, ok := newCtx.Value("AppName").(string); ok {
			app, err := model.AppInfoQuery.FindByAppName(appName)
			if err != nil {
				newCtx.Set("res_err", err.Error())
				newCtx.Set("app_id", 0)
				newCtx.Set("service_type", "error")
			} else {
				newCtx.Set("res_err", "")
				newCtx.Set("app_id", app.ID)
				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
				newCtx.Set("service_type", "业务应用用户账号存储")
			}
			log.Debugf("取出名称:%s", appName)
		}
	case "/v1/app/user/register":
		if appName, ok := newCtx.Value("AppName").(string); ok {
			app, err := model.AppInfoQuery.FindByAppName(appName)
			if err != nil {
				newCtx.Set("res_err", err.Error())
				newCtx.Set("app_id", 0)
				newCtx.Set("service_type", "error")
			} else {
				log.Debugf("根据应用名称查询的应用信息:%v", app)
				newCtx.Set("res_err", "")
				newCtx.Set("app_id", app.ID)
				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
				newCtx.Set("service_type", "业务应用代理用户进行主权账户注册")
			}
			log.Debugf("取出名称:%s", appName)
		}
	case "/v1/app/info":
		if appName, ok := newCtx.Value("AppName").(string); ok {
			app, err := model.AppInfoQuery.FindByAppName(appName)
			if err != nil {
				newCtx.Set("res_err", err.Error())
				newCtx.Set("app_id", 0)
				newCtx.Set("service_type", "error")
			} else {
				log.Debugf("根据应用名称查询的应用信息:%v", app)
				newCtx.Set("res_err", "")
				newCtx.Set("app_id", app.ID)
				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
				newCtx.Set("service_type", "业务应用代理用户进行主权账户注册")
			}
			log.Debugf("取出名称:%s", appName)
		}
	}
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
