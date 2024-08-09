package service

import (
	"coredx/core/api/v1/app/request"
	"coredx/core/api/v1/app/response"
	"coredx/db/dal/model"
	"github.com/gin-gonic/gin"
)

// DirectRegister 放一些平台内置API接口
// AppUserRegister
// @Tags  平台创建数据目录
// @Summary   平台创建数据目录
// @accept    application/json
// @Produce   application/json
// @Param data body request.InternalDirectoryRequest true  "平台创建数据目录"
// @Success   200   {object}  response.Response{data=string}  "平台创建数据目录"
// @Router    /v1/app/directory/register [post]
func DirectRegister(c *gin.Context) {
	var req request.InternalDirectoryRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	if err = model.DataDirectQuery.Create(&model.DataDirect{

		DirectID:   int64(req.DirectID),
		ParentID:   int64(req.ParentID),
		DirectName: req.DirectoryName,
		Tag:        req.Tag,
		Remark:     req.Remark,
	}); err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, nil)
}
