package middleware

import (
	"coredx/db/dal/model"
	"coredx/global"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	testUrl               = "/v1/app/test"
	bindUrl               = "/v1/app/account/bind"
	registerUrl           = "/v1/app/register"
	userRegisterUrl       = "/v1/app/user/register"
	appInfoUrl            = "/v1/app/info"
	userCaUrl             = "/v1/app/user/ca"
	dataStoreUrl          = "/v1/app/service/data/store"
	directoryRegisterUrl  = "/v1/app/directory/register"
	directorySubscribeUrl = "/v1/app/service/directory/subscribe"
	serviceListUrl        = "/v1/app/service/directory/subscribe"
	serviceBuyUrl         = "/v1/app/service/buy"
)

func in(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func filter(path string, ctx *gin.Context) error {

	if in(global.RouteStrings, path) {
		if strings.EqualFold(path, testUrl) {
			ctx.Set("AppName", "test")
			//newCtx 有AppName
			if _, ok := ctx.Value("AppName").(string); ok {
				ctx.Set("res_err", "")
				ctx.Set("app_id", uint(999))
				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
				ctx.Set("service_type", path)
				return nil
			}
		}
		if strings.EqualFold(path, directoryRegisterUrl) {
			ctx.Set("app_id", uint(9999))
			ctx.Set("res_err", "")
			ctx.Set("service_type", path)
			return nil
		}
		if appName, ok := ctx.Value("AppName").(string); ok {
			app, err := model.AppInfoQuery.FindByAppName(appName)
			if err != nil {
				ctx.Set("res_err", err.Error())
				ctx.Set("app_id", 0)
				ctx.Set("service_type", "error")
				return nil
			} else {
				ctx.Set("res_err", "")
				ctx.Set("app_id", app.ID)
				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
				ctx.Set("service_type", path)
				return nil
			}
		}
	}
	//log.Debugf("回调path:%s", path)
	return errors.New("no log")
}

//for _, v := range global.ApiRoutesInfo {
//	if v.Path == path {
//		if v.Path == testUrl {
//			ctx.Set("AppName", "test")
//			//newCtx 有AppName
//			if _, ok := ctx.Value("AppName").(string); ok {
//				ctx.Set("res_err", "")
//				ctx.Set("app_id", uint(999))
//				//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
//				ctx.Set("service_type", path)
//				return nil
//			}
//		} else if v.Path == directoryRegisterUrl {
//			ctx.Set("app_id", uint(9999))
//			ctx.Set("res_err", "")
//			ctx.Set("service_type", path)
//			return nil
//		} else {
//			if appName, ok := ctx.Value("AppName").(string); ok {
//				app, err := model.AppInfoQuery.FindByAppName(appName)
//				if err != nil {
//					ctx.Set("res_err", err.Error())
//					ctx.Set("app_id", 0)
//					ctx.Set("service_type", "error")
//					return nil
//				} else {
//					ctx.Set("res_err", "")
//					ctx.Set("app_id", app.ID)
//					//在日志表有一个服务类别的字段,但是在该url记录成注册接口即可(其他接口正常填智能,数权等服务)
//					ctx.Set("service_type", path)
//					return nil
//				}
//			}
//		}
//		continue
//	} else {
//		log.Debugf("回调path:%s", path)
//		return errors.New("no log")
//	}
//}
//	return nil
//}
