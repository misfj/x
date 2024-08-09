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
func AsyncUploadResponse(ctx *gin.Context, taskID string) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: taskID,
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
	PrivateMd5     string    `json:"private_md5"`
	StartTime      time.Time `json:"start_time"`
	NickName       string    `json:"nick_name"`
	UserLevel      string    `json:"user_level"`
	UserPhone      string    `json:"user_phone"`
	UserSourceType string    `json:"user_source_type"`
	UserSource     string    `json:"user_source"`
	UserType       string    `json:"user_type"`
	Did            string    `json:"did"`
}

type ServiceUploadResponse struct {
	NodeID   string `json:"node_id"` //节点ID
	FileID   string `json:"file_id"` //文件唯一标识,文件的md5
	FileName string `json:"file_name"`
}

type ServiceSubscribeServiceResponse struct {
	ServiceName      string `json:"service_name"`
	ServiceVersion   string `json:"service_version"`
	ServiceType      string `json:"service_type"`
	VisitUrl         string `json:"visit_url"`
	NodeID           string `json:"node_id"`
	ServiceAllowTime string `json:"service_allow_time"`
	ServiceExpTime   string `json:"service_exp_time"`
	UseType          string `json:"service_use_type"`
	Status           string `json:"status"`
}

type ServiceListResponse struct {
	ServiceName    string `json:"service_name"`
	ServiceVersion string `json:"service_version"`
	ServiceType    string `json:"service_type"`
	Status         string `json:"service_status"`
	VisitUrl       string `json:"visit_url"`
	Remark         string `json:"remark"`
	NodeID         string `json:"node_id"`
}

type SingleDataResponse struct {
	NickName      string `json:"nick_name"`
	DID           string `json:"did"`
	DataHash      string `json:"data_hash"`
	DataName      string `json:"data_name"`
	DataDirectory string `json:"data_directory"`
	DataType      string `json:"data_type"`
	DataSize      string ` json:"data_size"`
	DataUrl       string `json:"data_url"`
	NodeID        string `json:"node_id"`
	Remark        string `json:"remark"`
}

type SpaceInfoResponse struct {
	SpaceTotal     string `json:"space_total"`
	SpaceUse       string `json:"space_use"`
	SpaceAvailable string `json:"space_available"`
	Status         string `json:"status"`
}
