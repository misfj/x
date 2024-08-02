package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}
func FailedResponse(ctx *gin.Context, data interface{}, errMsg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  errMsg,
		Data: data,
	})
}

type SingleAppInfo struct {
	AppName       string `json:"app_name"`
	AppUser       string `json:"app_user"`
	AppPhone      string `json:"app_phone"`
	AppType       string `json:"app_type"`
	AppPublic     string `json:"app_public"`
	AppPublicMd5  string `json:"app_public_md5"`
	AppPrivateMd5 string `json:"app_private_md5"`
	AppAlgo       string `json:"app_algo"`
	StoreKey      string `json:"store_key"`
}
type SingleCA struct {
	UserName  string `json:"user_name"`
	Public    string `json:"public"`
	PublicMd5 string `json:"public_md5"`
	//Private    string `json:"private"`
	PrivateMd5 string    `json:"private_md5"`
	StartTime  time.Time `json:"start_time"`
}
