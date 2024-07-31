package middleware

import (
	"bytes"
	"context"
	"coredx/db"
	"coredx/db/dal/query"
	"coredx/log"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ReqLogger 记录每次请求的请求信息和响应信息
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		//AppID := 0
		//ServiceType := ""
		//ReqURL := ""
		//Method := ""
		//RemoteIP := ""
		//ReqContent := ""
		//ResMsg := ""
		//ResSize := ""
		//ResErr := ""
		// 请求前
		t := time.Now()
		reqPath := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()
		//rawQuery := c.Request.URL.RawQuery
		// JSON和FORM表单打印请求的Body, 其他内容类型，比如文件上传不打印
		contentType := c.GetHeader("Content-Type")
		if contentType != "" &&
			(strings.HasPrefix(contentType, "application/json") ||
				strings.HasPrefix(contentType, "application/x-www-form-urlencoded")) {
			log.Debugf("request type:%s\n", contentType)
			requestBody, err := io.ReadAll(c.Request.Body)
			if err != nil {
				log.Debugf("read body error:%s\n", err.Error())
				requestBody = []byte{}
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}
		log.Debugf("request path:%s\n", reqPath)
		//
		log.Debugf("请求路径:%s\n", reqPath)
		var reqMsg []byte
		var reqMsgErr error
		reqMsg, reqMsgErr = io.ReadAll(c.Request.Body)
		if reqMsgErr != nil {
			reqMsg = []byte(reqMsgErr.Error())
		}
		reqLog := fmt.Sprintf("host:%s %s %s start", ip, method, reqPath)
		log.Debugf("reqLog :%s, body:%s", reqLog, string(reqMsg))
		reqMap := make(map[string]interface{})
		_ = json.Unmarshal(reqMsg, &reqMap)
		app_name := ""
		ok := false
		log.Debugf("map:%v\n", reqMap)
		if app_name, ok = reqMap["app_name"].(string); !ok {
			log.Error("no app_name field")
			c.Abort()
		}
		if strings.EqualFold(reqPath, "/v1/app/login") {
			//根据应用名称查询数据库
			if app_name != "" {
				appInfo, err := query.Use(db.GDB).AppInfo.WithContext(context.Background()).Where(query.AppInfo.AppName.Eq(app_name)).Debug().First()
				if err != nil {
					log.Error(err)
					c.Abort()
				}
				log.Debugf("appInfo:%v error:%v\n", appInfo, err)
				log.Debug(appInfo.DeleteAt == time.Time{})
				//appInfo, err := query.Use(db.GDB).AppInfo.WithContext(context.Background()).
				if err != nil {
					log.Error(err)
				}
			}

		} else {
			fmt.Println(1111)
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		c.Next()
		// 请求后
		latency := time.Since(t).Microseconds()
		log.Debugf(fmt.Sprintf("host:%s %s %s end", ip, method, reqPath), "cost/us", latency)
		log.Debugf("Status Code:%d\n", c.Writer.Status())
		log.Debugf("Status Code:%d\n", c.Writer.Size())
		bd, _ := io.ReadAll(writer.body)
		log.Debugf("Response Body:%s\n", string(bd))

		//query.Use(db.GDB).AppLog.Create(&model.AppLog{
		//	CreateAt:    time.Now(),
		//	UpdateAt:    time.Now(),
		//	DeleteAt:    nil,
		//	AppID:       0,
		//	ServiceType: "",
		//	ReqURL:      "",
		//	Method:      "",
		//	RemoteIP:    "",
		//	ReqContent:  "",
		//	ResMsg:      "",
		//	ResSize:     0,
		//	ResErr:      "",
		//})
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
