package service

import (
	"bytes"
	"coredx/client"
	"coredx/config"
	"coredx/core/api/v1/app/request"
	"coredx/core/api/v1/app/response"
	"coredx/db/dal/model"
	"coredx/global"
	"coredx/log"
	"coredx/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/kardianos/osext"
	"github.com/wumansgy/goEncrypt/aes"
	"github.com/wumansgy/goEncrypt/rsa"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Upload
// @Tags 	  用户数据上传
// @Summary   用户数据上传
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ServiceUploadRequest true  "用户数据上传"
// @Success   200   {object}  response.Response{data=response.ServiceUploadResponse}  "用户数据上传"
// @Router    /v1/app/service/data/store [post]
func Upload(c *gin.Context) {
	//var form request.ServiceUploadRequest
	contentType := c.Request.Header.Get("Content-Type")
	log.Debug(contentType)
	var req request.ServiceUploadRequest
	pwd, err := osext.ExecutableFolder()
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	// 根据 Content-Type 处理数据(表单数据)
	if strings.Contains(contentType, "multipart/form-data") {
		err := c.ShouldBind(&req)
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		openFile, err := file.Open()
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		fileSize := file.Size
		if fileSize > global.Global.ThirtyMB {
			log.Debug("文件大于30M")
			//todo 立即进行响应,任务丢到协程和channel,进行回调,io.ReadAll在数据量大的时候,内存会不够用,待优化
			fileContent, err := io.ReadAll(openFile)
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			var appInfo *model.AppInfo
			//获取AppName 然后进行
			appName, exists := c.Get("AppName")
			if exists {
				appNameString, _ := appName.(string)
				appInfo, err = model.AppInfoQuery.FindByAppName(appNameString)
				if err != nil {
					response.FailedResponse(c, nil, "应用名称不存在")
					return
				}
			} else {
				response.FailedResponse(c, nil, "内部错误")
				return
			}
			//查询下DID 后期存放到中间件
			userID, exist := model.UserInfoQuery.ExistByDID(req.DID)
			if !exist {
				response.FailedResponse(c, nil, "did 在数据库不存在")
				return
			}
			meta := global.UploadMeta{
				AppID:       appInfo.ID,
				UserID:      userID,
				DataType:    req.FileType,
				Content:     fileContent,
				FileUrl:     file.Filename,
				FileSize:    fileSize,
				NodeID:      config.GetNode().NodePrivateMd5,
				FileName:    file.Filename,
				FileMd5:     utils.MD5(fileContent),
				CallbackUrl: req.CallbackUrl,
				StartTime:   time.Now(),
				EndTime:     time.Now(),
				TaskID:      uuid.NewString(),
			}
			global.Global.UploadCh <- meta
			response.SuccessResponse(c, meta.TaskID)
			return
		} else {
			log.Debug("数据小于30M")
			//数据大小小于30M

			dst := filepath.Join(pwd, "uploads", file.Filename)
			if err = c.SaveUploadedFile(file, dst); err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			fileContent, err := io.ReadAll(openFile)
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			//todo 数据库写入文件大小小于30M的数据日志
			//查询下DID 后期存放到中间件
			userID, exist := model.UserInfoQuery.ExistByDID(req.DID)
			if !exist {
				response.FailedResponse(c, nil, "did 在数据库不存在")
				return
			}
			err = model.DataInfoQuery.Create(&model.DataInfo{
				UseID:          int64(userID),
				DataID:         utils.MD5(fileContent),
				DataName:       file.Filename,
				DataDiectoryID: 1,
				DataType:       req.FileType,
				IsOpen:         global.Open,
				DataSize:       file.Size,
				DataPos:        dst,
				NodeID:         config.GetNode().NodePrivateMd5,
			})
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			response.SuccessResponse(c, "success")
		}
	}
	//todo 处理json数据请求
	if strings.Contains(contentType, "application/json") {
		err := c.ShouldBindJSON(&req)
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		//appName := req.AppName
		did := req.DID
		password := req.Password
		fileUrl := req.FileUrl
		fileName := req.FileName
		//fileMd5 := req.FileMD5
		fileRemark := req.FileRemark
		//fileDirectory := req.FileDirectory
		fileType := req.FileType
		//callBackUrl := req.CallbackUrl
		//校验did
		err = model.UserInfoQuery.VerifyDIDPassword(did, utils.Bcrypt(password))
		if err != nil {
			response.FailedResponse(c, nil, "did 在数据库不存在")
			return
		}
		userID, exist := model.UserInfoQuery.ExistByDID(req.DID)
		if !exist {
			response.FailedResponse(c, nil, "did 在数据库不存在")
			return
		}
		//下载文件(fileUrl)
		resp, err := client.Client.HttpC.R().Get(fileUrl)
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}

		fileHash := utils.MD5(resp.Body())
		if fileName == "" {
			fileName = fileHash
		}

		filer, err := os.Create(filepath.Join(pwd, "uploads", fileName))
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		io.Copy(filer, bytes.NewReader(resp.Body()))
		err = model.DataInfoQuery.Create(&model.DataInfo{
			UseID:          int64(userID),
			DataID:         fileHash,
			DataName:       fileName,
			DataDiectoryID: 1,
			DataType:       fileType,
			IsOpen:         global.Open,
			DataSize:       int64(len(resp.Body())),
			DataPos:        filepath.Join(pwd, "uploads", fileName),
			NodeID:         config.GetNode().NodePrivateMd5,
			Remark:         fileRemark,
		})
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		response.SuccessResponse(c, nil)
	}
}

// Subscribe
// @Tags 	  业务应用订阅平台数据目录
// @Summary   业务应用订阅平台数据目录
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ServiceDirectorySubscribeRequest true  "业务应用订阅平台数据目录"
// @Success   200   {object}  response.Response{data=response.ServiceUploadResponse}  "业务应用订阅平台数据目录"
// @Router    /v1/app/service/directory/subscribe [post]
func Subscribe(c *gin.Context) {
	var req request.ServiceDirectorySubscribeRequest
	//todo 验证AppName
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//通过AppName获取业务应用名称
	appName, _ := c.Get("AppName")
	appNameString, _ := appName.(string)
	app, err := model.AppInfoQuery.FindByAppName(appNameString)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//查询数据库是否存在目录ID 如果不存在就返回错误信息
	exist := model.DataDirectQuery.Exist(req.RootDir)
	if !exist {
		response.FailedResponse(c, nil, "没有该目录ID")
		return
	}
	//优先支持parent_id为0的目录,也就是顶级目录
	if err = model.AppDirectQuery.Create(&model.AppDirect{
		AppID:    int64(app.ID),
		DirectID: int64(req.RootDir),
		Status:   strconv.Itoa(1),
	}); err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, nil)
}

// SubscribeService
// @Tags 	  业务应用查看订阅的服务信息
// @Summary   业务应用查看订阅的服务信息
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Success   200   {object}  response.Response{data=[]response.ServiceSubscribeServiceResponse}  "业务应用查看订阅的服务信息"
// @Router    /v1/app/service/subscribe [get]
func SubscribeService(c *gin.Context) {
	//
}

// BuyService
// @Tags 	  业务应用购买服务
// @Summary   业务应用购买服务
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ServiceBuyRequest true  "业务应用购买服务"
// @Success   200   {object}  response.Response{data=[]response.ServiceSubscribeServiceResponse}  "业务应用购买服务"
// @Router    /v1/app/service/buy [post]
func BuyService(c *gin.Context) {
	//todo 后续支持批量购买接口
	var req request.ServiceBuyRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//通过AppName获取业务应用名称
	appName, _ := c.Get("AppName")
	appNameString, _ := appName.(string)
	log.Debug(req)
	//根据服务ID查看服务ID是否存在
	if !model.ServiceInfoQuery.ExistServiceID(req.ServiceID) {
		response.FailedResponse(c, nil, "数据库不存在该服务ID")
		return
	}
	//绑定业务应用ID
	app, err := model.AppInfoQuery.FindByAppName(appNameString)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	expTime := utils.MillSecondTime(req.ServiceExp)
	if err = model.AppAllowQuery.Create(&model.AppAllow{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		AppID:        int64(app.ID),
		ServiceID:    int32(req.ServiceID),
		ServiceAllow: time.Now(),
		ServiceExp:   expTime,
		UseType:      "3",
		Status:       "3",
	}); err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, nil)
}

// ServiceList
// @Tags 	  业务应用获取服务列表
// @Summary   业务应用获取服务列表
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ServiceListRequest true  "业务应用购买服务"
// @Success   200   {object}  response.Response{data=[]response.ServiceListResponse}  "业务应用购买服务"
// @Router    /v1/app/service/list [post]
func ServiceList(c *gin.Context) {
	//暂时不需要分页,就几条数据
	list, err := model.ServiceInfoQuery.List()
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	log.Debug(list)
	res := make([]response.ServiceListResponse, 0, 3)

	err = copier.Copy(&res, &list)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, res)
}

// DataList
// @Tags 	  业务应用获取数据列表
// @Summary   业务应用获取数据列表
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ServiceDataListRequest true  "业务应用获取数据列表"
// @Success   200   {object}  response.Response{data=[]response.SingleDataResponse}  "业务应用获取数据列表"
// @Router    /v1/app/service/data/list [post]

func DataList(c *gin.Context) {
	var req request.ServiceDataListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//通过AppName获取业务应用名称
	//appName, _ := c.Get("AppName")
	//appNameString, _ := appName.(string)
	//根据id获取用户信息
	list, err := model.DataInfoQuery.List(req.PageNum, req.PageSize, req.Conditions)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, list)
}

// Water
// @Tags 	  业务应用访问水印服务添加水印
// @Summary   业务应用访问水印服务添加水印
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ServiceWaterAddRequest true  "业务应用访问水印服务添加水印"
// @Success   200   {object}  response.Response{data=string}  "业务应用访问水印服务添加水印"
// @Router    /v1/app/service/water [post]
func Water(c *gin.Context) {
	var req request.ServiceWaterAddRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	remoteUdpAddress := ""
	defaultIp := ""
	args := strings.Split(config.Config.Watertwo.Args, " ")
	splitIport := strings.Split(args[0], ":")
	//fmt.Println(len(splitIport))
	//fmt.Println(splitIport[0])
	//fmt.Println(splitIport[1])
	if splitIport[0] == "" {
		//补全UDP地址
		remoteUdpAddress = fmt.Sprintf("%s%s", defaultIp, args[0])
	} else {
		remoteUdpAddress = args[0]
	}
	data, _ := json.Marshal(req)
	global.SendUDPMessage(remoteUdpAddress, data)
	resp := <-global.Global.UdpCh
	response.SuccessResponse(c, resp)
}

// CG
// @Tags 	  业务应用进行授权确权
// @Summary   业务应用进行授权确权
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.CgRequest true  "业务应用进行授权确权"
// @Success   200   {object}  response.Response{data=string}  "业务应用进行授权确权"
// @Router    /v1/app/service/cg [post]
func CG(c *gin.Context) {
	fmt.Println(1111111111)
	var req request.CgRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	if req.CgDataMD5 != "" {
		//进行比对
		if !strings.EqualFold(utils.MD5([]byte(req.CgData)), req.CgDataMD5) {
			response.FailedResponse(c, nil, "md5 不匹配")
			return
		}
		ownerUserID, e1 := model.UserInfoQuery.ExistByDID(req.OwnerDID)
		peerUserID, e2 := model.UserInfoQuery.ExistByDID(req.PeerDID)
		if e1 == true && e2 == true {
			//进行确权
			if strings.EqualFold(req.Cg, "c") {
				fmt.Println(1111111111)
				ownerCa, err := model.UserCaQuery.FindByUserID(ownerUserID)
				if err != nil {
					response.FailedResponse(c, nil, "did not exist")
					return
				}
				if strings.EqualFold(ownerCa.Algo, "rsa") {
					//todo 可能待授权确权的数据太大,需要进行分块处理,然后进行合并
					signStr, err := rsa.RsaSignBase64([]byte(req.CgData), ownerCa.Private)
					if err != nil {
						response.FailedResponse(c, nil, "sm2 is supporting")
						return
					}
					signStrMd5 := utils.MD5([]byte(signStr))
					response.SuccessResponse(c, map[string]string{
						"sign":    signStr,
						"signMd5": signStrMd5,
					})
					return
				} else {
					//todo 支持sm2
					response.FailedResponse(c, nil, "sm2 is supporting")
					return
				}
			}
			//进行授权
			if strings.EqualFold(req.Cg, "g") {
				peerCa, err := model.UserCaQuery.FindByUserID(peerUserID)
				if err != nil {
					response.FailedResponse(c, nil, "did not exist")
					return
				}
				if strings.EqualFold(peerCa.Algo, "rsa") {
					//todo 可能待授权确权的数据太大,需要进行分块处理,然后进行合并
					signStr, err := rsa.RsaSignBase64([]byte(req.CgData), peerCa.Private)
					if err != nil {
						response.FailedResponse(c, nil, "sm2 is supporting")
						return
					}
					signStrMd5 := utils.MD5([]byte(signStr))
					response.SuccessResponse(c, map[string]string{
						"sign":    signStr,
						"signMd5": signStrMd5,
					})
					return
				} else {
					//todo 支持sm2
					response.FailedResponse(c, nil, "sm2 is supporting")
					return
				}
			}
			response.FailedResponse(c, nil, "cg is supporting")
			return
			//进行授权确权
		} else {
			response.FailedResponse(c, nil, "did not exist")
			return
		}
	} else {
		//查询数据是否存在双方di
		ownerUserID, e1 := model.UserInfoQuery.ExistByDID(req.OwnerDID)
		peerUserID, e2 := model.UserInfoQuery.ExistByDID(req.PeerDID)
		if e1 == true && e2 == true {
			//进行确权
			if strings.EqualFold(req.Cg, "c") {
				fmt.Println(1111111111)
				ownerCa, err := model.UserCaQuery.FindByUserID(ownerUserID)
				if err != nil {
					response.FailedResponse(c, nil, "did not exist")
					return
				}
				if strings.EqualFold(ownerCa.Algo, "rsa") {
					//todo 可能待授权确权的数据太大,需要进行分块处理,然后进行合并
					signStr, err := rsa.RsaSignBase64([]byte(req.CgData), ownerCa.Private)
					if err != nil {
						response.FailedResponse(c, nil, "sm2 is supporting")
						return
					}
					signStrMd5 := utils.MD5([]byte(signStr))
					response.SuccessResponse(c, map[string]string{
						"sign":    signStr,
						"signMd5": signStrMd5,
					})
					return
				} else {
					//todo 支持sm2
					response.FailedResponse(c, nil, "sm2 is supporting")
					return
				}
			}
			//进行授权
			if strings.EqualFold(req.Cg, "g") {
				peerCa, err := model.UserCaQuery.FindByUserID(peerUserID)
				if err != nil {
					response.FailedResponse(c, nil, "did not exist")
					return
				}
				if strings.EqualFold(peerCa.Algo, "rsa") {
					//todo 可能待授权确权的数据太大,需要进行分块处理,然后进行合并
					signStr, err := rsa.RsaSignBase64([]byte(req.CgData), peerCa.Private)
					if err != nil {
						response.FailedResponse(c, nil, "sm2 is supporting")
						return
					}
					signStrMd5 := utils.MD5([]byte(signStr))
					response.SuccessResponse(c, map[string]string{
						"sign":    signStr,
						"signMd5": signStrMd5,
					})
					return
				} else {
					//todo 支持sm2
					response.FailedResponse(c, nil, "sm2 is supporting")
					return
				}
			}
			response.FailedResponse(c, nil, "cg is supporting")
			return
			//进行授权确权
		} else {
			response.FailedResponse(c, nil, "did not exist")
			return
		}
	}
}

// AI
// @Tags 	  AI服务正在光速开发中😀
// @Summary   AI服务正在光速开发中
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.CgRequest true  "AI服务正在光速开发中"
// @Success   200   {object}  response.Response{data=string}  "AI服务正在光速开发中"
// @Router    /v1/app/service/ai [post]
func AI(c *gin.Context) {

}

// Pay
// @Tags 	  支付服务正在光速开发中😀
// @Summary   支付服务正在光速开发中
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.CgRequest true  "支付服务正在光速开发中"
// @Success   200   {object}  response.Response{data=string}  "支付服务正在光速开发中"
// @Router    /v1/app/service/pay [post]
func Pay(c *gin.Context) {

}

// Crypto
// @Tags 	  文件加解密服务
// @Summary   文件加解密服务
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.CryptoRequest true  "文件加解密服务"
// @Success   200   {object}  response.Response{data=string}  "文件加解密服务"
// @Router    /v1/app/service/crypto [post]
func Crypto(c *gin.Context) {
	//
	contentType := c.Request.Header.Get("Content-Type")
	fmt.Println(strings.Contains(contentType, "multipart/form-data"))
	var req request.CryptoRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//form 表单
	if strings.Contains(contentType, "multipart/form-data") {
		//todo 转存第三方存储服务器
		// 保存文件到本地的文件夹下
		file, err := c.FormFile("file")
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		openFile, err := file.Open()
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		fileName := file.Filename
		fileSize := file.Size
		fileCody, err := io.ReadAll(openFile)
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		bufMd5 := utils.MD5(fileCody)

		execfolder, err := osext.ExecutableFolder()
		if err != nil {
			response.FailedResponse(c, nil, fmt.Sprintf("%s: 获取可执行路径失败", err.Error()))
			return
		}

		dst := filepath.Join(execfolder, global.UploadsDir, fileName)
		//保存源文件
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			response.FailedResponse(c, nil, fmt.Sprintf("%s: 保存文件失败", err.Error()))
			return
		}
		log.Debug("--------------------------------")
		//todo 为了数据的完整可恢复,我们不能只存储加密之后的文件,我们还要存储源数据文件。
		//加密数据一旦损坏,就会找不回来
		//此时我们将待加密的数据源文件存放到encrypt-files
		_, err = os.Stat(filepath.Join(execfolder, global.EncryptFileDir))
		if err != nil {

			err = os.Mkdir(filepath.Join(execfolder, global.EncryptFileDir), 0666)
			if err != nil {
				response.FailedResponse(c, nil, fmt.Sprintf("%s: 创建encrypt-files文件夹失败", err.Error()))
				return
			}
		}
		//代表加密,从缓存
		if req.Action == 1 {
			//form-data的方式我们先不查数据库
			//todo 考虑到数据太大,走回调
			if fileSize > global.Global.ThirtyMB {
				taskID := uuid.NewString()
				//将信息丢入通道
				global.Global.CryptoCh <- &global.CryptoMeta{
					Content:     fileCody,
					FileName:    fileName,
					FileSize:    fileSize,
					FileMd5:     bufMd5,
					FileKey:     global.AES128Default,
					FileAlgo:    global.AES128,
					CallbackUrl: req.CallbackUrl,
					//通道处理过程中进行补全字段
					FileEncryptedHash: "",
					FileUrl:           filepath.Join(execfolder, global.EncryptFileDir, fileName),
					NodeID:            config.GetNode().NodePrivateMd5,
				}
				response.SuccessResponse(c, taskID)
				return
			}
			//小于30M
			if fileSize <= global.Global.ThirtyMB {

				filer, err := os.Create(filepath.Join(execfolder, global.EncryptFileDir, fileName))
				if err != nil {
					log.Error(err)
					response.FailedResponse(c, nil, err.Error())
					return
				}
				//io.ReadAll内存会爆炸,如果
				fileContent, err := io.ReadAll(openFile)
				if err != nil {
					log.Error(err)
					response.FailedResponse(c, nil, err.Error())
					return
				}
				//默认使用AES128进行加密,因为对内存和时间的消耗的都是最低的
				cipherText, err := aes.AesCbcEncrypt(fileContent, []byte(global.AES128Default), []byte(global.AES128Default))
				if err != nil {
					log.Error(err)
					response.FailedResponse(c, nil, err.Error())
					return
				}
				_, err = io.Copy(filer, bytes.NewReader(cipherText))
				if err != nil {
					log.Error(err)
					response.FailedResponse(c, nil, err.Error())
					return
				}
				//数据库写入记录
				err = model.DataEncQuery.Create(&model.DataEnc{
					FileName:          file.Filename,
					FileSize:          fileSize,
					FileMd5:           bufMd5,
					FileKey:           global.AES128Default,
					FileAlgo:          global.AES128,
					FileEncryptedHash: utils.MD5(cipherText),
					FileURL:           filepath.Join(execfolder, global.EncryptFileDir, fileName),
					NodeID:            config.GetNode().NodePrivateMd5,
				})
				if err != nil {
					log.Error(err)
					response.FailedResponse(c, nil, err.Error())
					return
				}
				response.SuccessResponse(c, nil)
				//var de *model.DataEnc
				//if de, err = model.DataEncQuery.ExistRawByDataAid(req.DataAID, ""); err != nil {
				//	//todo 先去data_info这张表里面去找一下
				//	dataInfo, err := model.DataInfoQuery.FindByDataID(req.DataAID)
				//	if err != nil {
				//		response.FailedResponse(c, nil, fmt.Sprintf("%s:没找到原始数据,请上传一份原始数据", err.Error()))
				//		return
				//	}
				//	//todo 通过dataInfo 获取数据访问链接,目前存在本地uploads文件下,以后切换到其他文件存储服务器
				//	visitUrl := dataInfo.DataPos
				//	fileContent, err := os.ReadFile(visitUrl)
				//	if err != nil {
				//		response.FailedResponse(c, nil, fmt.Sprintf("%s: 读取文件失败", err.Error()))
				//		return
				//	}
				//if de.FileKey != "" {
				//	//数据未加密
				//}
				//if strings.EqualFold(req.Algo, "aes") {
				//
				//}
			}
		}
		//代表解密
		if req.Action == 2 {
			log.Debug(len(fileCody))
			readFile, _ := os.ReadFile(filepath.Join(execfolder, global.EncryptFileDir, file.Filename))
			decrypt, err := aes.AesCbcDecrypt(readFile, []byte(global.AES128Default), []byte(global.AES128Default))
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			c.Writer.Write(decrypt)
		}

		//存在数据

	}
	//body json数据
	if strings.Contains(contentType, "application/json") {
		log.Debug(req)
	}
	response.FailedResponse(c, nil, "其他数据格式正在光速开发中😀")
	//验证AppName
	//去数据库查询是否存在dataAID
	//找出dataAID的数据属性
	//appName := req.AppName

	log.Debug(req)
}

//log.Debug(req)
//var de *model.DataEnc
//if de, err = model.DataEncQuery.ExistRawByDataAid(req.DataAID, ""); err != nil {
//	//todo 先去data_info这张表里面去找一下
//	dataInfo, err := model.DataInfoQuery.FindByDataID(req.DataAID)
//	if err != nil {
//		response.FailedResponse(c, nil, fmt.Sprintf("%s:没找到原始数据,请上传一份原始数据", err.Error()))
//		return
//	}
//	//todo 通过dataInfo 获取数据访问链接,目前存在本地uploads文件下,以后切换到其他文件存储服务器
//	visitUrl := dataInfo.DataPos
//	fileContent, err := os.ReadFile(visitUrl)
//	if err != nil {
//		response.FailedResponse(c, nil, fmt.Sprintf("%s: 读取文件失败", err.Error()))
//		return
//	}

// Store
// @Tags 	  存储数据文件光速开发中😀
// @Summary   存储数据文件
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.StoreRequest true  "存储数据文件"
// @Success   200   {object}  response.Response{data=string}  "存储数据文件"
// @Router    /v1/app/service/data/store [post]
func Store(c *gin.Context) {
	var req request.StoreRequest
	err := c.BindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, "参数解析失败")
		return
	}
	//contentType := c
}

// SpaceInfo
// @Tags 	  用户空间状态信息😀
// @Summary   用户空间状态信息
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.StoreRequest true  "用户空间状态信息"
// @Success   200   {object}  response.Response{data=response.SpaceInfoResponse}  "用户空间状态信息"
// @Router    /v1/app/service/space/info [post]
func SpaceInfo(c *gin.Context) {
	var req request.SpaceInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//todo  验证AppName 的信息
	//根据did 查询数据库信息是否存在记录
	userID, b := model.UserInfoQuery.ExistByDID(req.DID)
	if !b {
		response.FailedResponse(c, nil, "did不存在")
		return
	}
	userInfo, err := model.UserInfoQuery.FindByID(userID)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//进行密码验证
	//bcrypt := utils.Bcrypt(req.Password)
	err = model.UserInfoQuery.VerifyDIDPassword(req.DID, req.Password)
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, "密码错误")
		return
	}
	pdc, err := model.UserPdcQuery.FindByUserID(int(userInfo.ID))
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	log.Debugf("用户pdc:%v\n", pdc)
	res := response.SpaceInfoResponse{}
	err = copier.Copy(&res, pdc)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, res)
}

// Expand
// @Tags 	  业务应用代理用户空间进行扩容
// @Summary   业务应用代理用户空间进行扩容
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ExpandRequest true  "业务应用代理用户空间进行扩容"
// @Success   200   {object}  response.Response{data=string}  "业务应用代理用户空间进行扩容"
// @Router    /v1/app/service/expand [post]
func Expand(c *gin.Context) {
	var req request.ExpandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	userID, b := model.UserInfoQuery.ExistByDID(req.DID)
	if !b {
		response.FailedResponse(c, nil, "did不存在")
		return
	}
	//todo 验证应用名称是否合法
	_, err := model.UserPdcQuery.ExpandCapacity(int64(userID), float64(req.ExpandSpaceSize))
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, "扩容成功")
}

// Update
// @Tags 	  业务应用代理用户进行数据更新
// @Summary   业务应用代理用户进行数据更新
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.UpdateRequest true  "业务应用代理用户进行数据更新"
// @Success   200   {object}  response.Response{data=string}  "业务应用代理用户进行数据更新"
// @Router    /v1/app/data/update [post]
func Update(c *gin.Context) {
	var req request.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//检查DID 是否存在
	_, b := model.UserInfoQuery.ExistByDID(req.DID)
	if !b {
		response.FailedResponse(c, nil, "did不存在")
		return
	}
	if err := model.DataInfoQuery.CloneDb().Model(&model.DataInfo{}).
		Where("data_id = ?", req.DataID).Updates(map[string]interface{}{
		"data_name": req.DataName,
		"data_type": req.DataType,
		"is_open":   req.IsOpen,
		"data_pos":  req.DataPos,
		"node_id":   req.NodeID,
		"remark":    req.Remark,
	}).Error; err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, "数据更新成功")
}

// Delete
// @Tags 	  业务应用代理用户进行数据删除
// @Summary   业务应用代理用户进行数据删除
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.DeleteRequest true  "业务应用代理用户进行数据删除"
// @Success   200   {object}  response.Response{data=string}  "业务应用代理用户进行数据删除"
// @Router    /v1/app/service/delete [post]
func Delete(c *gin.Context) {
	var req request.DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//todo 验证AppName
	//检查DID是否存在
	userID, b := model.UserInfoQuery.ExistByDID(req.DID)
	if !b {
		response.FailedResponse(c, nil, "did不存在")
		return
	}
	if err := model.DataInfoQuery.CloneDb().Model(&model.DataInfo{}).
		Where("data_id = ? and use_id = ? ", req.DataID, userID).Delete(&model.DataInfo{}).Error; err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, "数据删除成功")
}

// Share
// @Tags 	  数据共享
// @Summary   数据共享
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.ShareRequest true  "数据共享"
// @Success   200   {object}  response.Response{data=string}  "数据共享"
// @Router    /v1/app/data/share [post]
func Share(c *gin.Context) {
	var req request.ShareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//todo 找节点信息,找数据信息,目前就在平台找吧
	//根据did 获取用户信息,根据用户信息比对数据信息
	userID, b := model.UserInfoQuery.ExistByDID(req.DIDSeller)
	if !b {
		response.FailedResponse(c, nil, "seller did不存在")
		return
	}
	var dataInfo model.DataInfo
	if err := model.DataInfoQuery.CloneDb().Model(&model.DataInfo{}).Where("use_id = ? and data_id = ?", userID, req.DataID).
		First(&dataInfo).Error; err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//在对方的数据里面添加一条数据
	userID, b = model.UserInfoQuery.ExistByDID(req.DIDBuyer)
	if !b {
		response.FailedResponse(c, nil, "buyer did不存在")
		return
	}
	if err := model.DataInfoQuery.CloneDb().Model(&model.DataInfo{}).Create(&model.DataInfo{
		UseID:          int64(userID),
		DataID:         dataInfo.DataID,
		DataName:       dataInfo.DataName,
		DataDiectoryID: dataInfo.DataDiectoryID,
		DataType:       dataInfo.DataType,
		IsOpen:         dataInfo.IsOpen,
		DataSize:       dataInfo.DataSize,
		DataPos:        dataInfo.DataPos,
		NodeID:         dataInfo.NodeID,
		Remark:         dataInfo.Remark,
	}).Error; err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//todo 添加回调
	response.SuccessResponse(c, "数据共享成功")
	//model.DataInfoQuery.CloneDb().Model(&model.DataInfo{}).Where("")
}

// Rules
// @Tags 	  设置数据安全策略
// @Summary   设置数据安全策略
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.RulesRequest true  "设置数据安全策略"
// @Success   200   {object}  response.Response{data=string}  "设置数据安全策略"
// @Router    /v1/app/service/rules [post]
func Rules(c *gin.Context) {
	var req request.RulesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//todo 检查AppName是否存在
	//获取用户ID
	userID, b := model.UserInfoQuery.ExistByDID(req.DID)
	if !b {
		response.FailedResponse(c, nil, "did 在平台没查询到")
		return
	}
	//查询是否存在dataID数据
	if err := model.DataInfoQuery.CloneDb().Model(&model.DataInfo{}).Where("user_id = ? and data_id = ?", userID, req.DataID).
		First(&model.DataInfo{}).Error; err != nil {
		response.FailedResponse(c, nil, "dataID 在平台没查询到")
		return
	}
	//设置数据安全规则
	if err := model.DataSafeQuery.CloneDb().Model(&model.DataSafe{}).
		Where("data_id = ? and user_id = ?", req.DataID, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//创建记录
			model.DataSafeQuery.CloneDb().Model(&model.DataSafe{}).Create(&model.DataSafe{
				DataID:    req.DataID,
				SafeLevel: req.SafeLevel,
				SafeInfo:  req.SafeInfo,
				UserID:    int64(userID),
			})
		} else {
			response.FailedResponse(c, nil, err.Error())
			return
		}
	}
}

func Notice(c *gin.Context) {

}
