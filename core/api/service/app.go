package service

import (
	"coredx/config"
	"coredx/core/api/middleware"
	"coredx/core/api/v1/app/request"
	_ "coredx/core/api/v1/app/request"
	"coredx/core/api/v1/app/response"
	_ "coredx/core/api/v1/app/response"
	"coredx/db/dal/model"
	"coredx/log"
	"coredx/utils"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"github.com/wumansgy/goEncrypt/rsa"
)

const (
	RSAAlgo      = "rsa"
	Sm2Algo      = "sm2"
	NormalStatus = "1"
	ForbidStatus = "2"
	StopStatus   = "3"
)

// AppTest
// @Tags      API测试
// @Summary   API测试
// @accept    application/json
// @Produce   application/json
// @Param data body request.AppTestRequest true  "请求参数"
// @Success   200   {object}  response.Response{data=string}  "测试接口"
// @Router    /test [post]
func AppTest(c *gin.Context) {
	log.Debug("come----------------------")
	var req *request.AppTestRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Code: 200,
		Data: "success",
		Msg:  "success",
	})
}

// AppRegister
// @Tags      业务应用注册
// @Summary   业务应用注册
// @accept    application/json
// @Produce   application/json
// @Param data body request.AppRegisterRequest true  "请求参数"
// @Success   200   {object}  response.Response{data=string}  "业务应用注册接口"
// @Router    /v1/app/register [post]
func AppRegister(c *gin.Context) {
	var req request.AppRegisterRequest
	// var pair rsa.RsaKey
	// var err error
	//err := c.ShouldBindBodyWith(&req, binding.JSON)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	_, err = model.AppInfoQuery.FindByAppName(req.AppName)
	if err != nil {
		if req.Algo == "rsa" {
			pair, err := rsa.GenerateRsaKeyBase64(2048)
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			if err := utils.IslegalRsaError(pair.PublicKey, pair.PrivateKey); err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			if err = model.AppInfoQuery.Create(&model.AppInfo{
				AppName:       req.AppName,
				AppUser:       req.AppUser,
				AppPassword:   utils.Bcrypt(req.AppPassword),
				AppPhone:      req.AppPhone,
				AppType:       req.AppType,
				AppPublic:     pair.PublicKey,
				AppPublicMd5:  utils.MD5([]byte(pair.PublicKey)),
				AppPrivate:    pair.PrivateKey,
				AppPrivateMd5: utils.MD5([]byte(pair.PrivateKey)),
				Algo:          RSAAlgo,
				StoreKey:      "",
				StartTime:     time.Now(),
				Status:        NormalStatus,
			}); err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			c.Set("AppName", req.AppName)
			response.SuccessResponse(c, "业务应用注册成功")
			return
		}
		if req.Algo == "sm2" {
			sm2pair, err := sm2.GenerateKey(rand.Reader)
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			if err = utils.Sm2IsLegalError(sm2pair, &sm2pair.PublicKey); err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			sm2PublicBase64, err := x509.WritePublicKeyToPem(&sm2pair.PublicKey)
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			sm2PrivateBase64, err := x509.WritePrivateKeyToPem(sm2pair, []byte(req.StoreKey))
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			if err = model.AppInfoQuery.Create(&model.AppInfo{
				AppName:       req.AppName,
				AppUser:       req.AppUser,
				AppPassword:   utils.Bcrypt(req.AppPassword),
				AppPhone:      req.AppPhone,
				AppType:       req.AppType,
				AppPublic:     string(sm2PublicBase64),
				AppPublicMd5:  utils.MD5(sm2PublicBase64),
				AppPrivate:    string(sm2PrivateBase64),
				AppPrivateMd5: utils.MD5(sm2PrivateBase64),
				Algo:          RSAAlgo,
				StoreKey:      "",
				StartTime:     time.Now(),
				Status:        NormalStatus,
			}); err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			//ctx 必须传递这个AppName,写日志需要
			c.Set("AppName", req.AppName)
			response.SuccessResponse(c, "业务应用注册成功")
			return
		}
		response.FailedResponse(c, nil, "平台不支持其他算法")
		return
	} else {
		c.Set("AppName", req.AppName)
		response.SuccessResponse(c, "业务应用已经存在")
	}

}

// AppUserRegister
// @Tags  业务应用代理用户注册
// @Summary   业务应用代理用户注册
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.AppUserRegisterRequest true  "业务应用代理用户注册"
// @Success   200   {object}  response.Response{data=string}  "业务应用代理用户注册"
// @Router    /v1/app/user/register [post]
func AppUserRegister(c *gin.Context) {
	//开启业务应用名称与Token的校验规则,在数据库新增表token与业务应用信息绑定表
	var req request.AppUserRegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	appName, _ := c.Get("AppName")
	appNameString, _ := appName.(string)
	_, err = model.AppInfoQuery.FindByAppName(appNameString)
	if err != nil {
		response.FailedResponse(c, nil, "应用名称不存在")
		return
	}
	userName := req.UserName
	nickName := req.NickName
	password := req.Password
	userLevel := "0" //最低用户级别
	userPhone := req.UserPhone
	//记录用户存储来源
	userSourceType := "3" //1表示超级节点 2表示平台注册 3业务应用
	//userSource 用业务引用名称:节点ID进行表示 记录用户存储来源详细信息
	userSource := fmt.Sprintf("%s:%s", appNameString, config.GetNode().NodePrivateMd5)
	userType := "1" //1表示个人 2表示企业 3表示政府部门
	//did 有用户的证书生成,私钥的MD5
	pair, err := rsa.GenerateRsaKeyBase64(2048)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	err = utils.IslegalRsaError(pair.PublicKey, pair.PrivateKey)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	did := utils.MD5([]byte(pair.PrivateKey))

	if err = model.UserInfoQuery.Create(&model.UserInfo{
		UserName:       userName,
		NickName:       nickName,
		Password:       utils.Bcrypt(password),
		UserLevel:      userLevel,
		UserPhone:      userPhone,
		UserSourceType: userSourceType,
		UserSource:     userSource,
		UserType:       userType,
		Did:            did,
	}); err != nil {
		log.Debug(len(userSource))
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//查询用户ID
	userID, err := model.UserInfoQuery.FindUserIDByNickName(nickName)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	log.Debugf("用户ID:%d", userID)
	//创建证书记录
	err = model.UserCaQuery.Create(&model.UserCa{
		UserID:     int64(userID),
		Public:     pair.PublicKey,
		PublicMd5:  utils.MD5([]byte(pair.PublicKey)),
		Private:    pair.PrivateKey,
		PrivateMd5: utils.MD5([]byte(pair.PrivateKey)),
		Algo:       "rsa",
		StoreKey:   "",
		TimeStamp:  time.Now().Unix(),
	})
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//创建用户pdc 容器空间
	if err = model.UserPdcQuery.CloneDb().Where("user_id = ?", userID).First(&model.UserPdc{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//创建记录
			if err = model.UserPdcQuery.CloneDb().Create(&model.UserPdc{
				UserID: int64(userID),
				//默认500M
				SpaceTotal:     1024 * 1024 * 500,
				SpaceUse:       0,
				SpaceAvailable: 0,
				Status:         "1",
			}).Error; err != nil {
				log.Error(err.Error())
				response.FailedResponse(c, nil, err.Error())
				return
			}
		} else {
			log.Error(err.Error())
			response.FailedResponse(c, nil, err.Error())
			return
		}
	}
	model.UserPdcQuery.CloneDb().Where("")
	response.SuccessResponse(c, did)

}

// AppLogin
// @Tags      业务应用登录
// @Summary   业务应用登录
// @accept    application/json
// @Produce   application/json
// @Param data body request.AppLoginRequest true  "请求参数"
// @Success   200   {object}  response.Response{data=string}  "业务应用登录"
// @Router    /v1/app/login [post]
func AppLogin(c *gin.Context) {
	var req request.AppLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//查询数据库是否存在该业务应用的token,存在就返回,不存在就创建
	bcryptPassword := utils.Bcrypt(req.AppPassword)
	//验证密码
	if model.AppInfoQuery.VerifyUserPassword(req.AppName, bcryptPassword) {
		//查询业务应用信息
		app, err := model.AppInfoQuery.FindByAppName(req.AppName)
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		//查询appID
		token, err := model.AppTokenQuery.FindByAppID(int(app.ID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//创建token 插入数据库 然后返回
				exp := time.Now().Add(time.Duration(config.GetJwt().Exp) * time.Hour * 24)
				cusClaim := middleware.BuildClaims(exp, req.AppName)
				//签发token
				token, err = middleware.GenToken(cusClaim, config.GetJwt().Secret)
				if err != nil {
					response.FailedResponse(c, nil, err.Error())
					return
				}
				//插入数据库
				if err = model.AppTokenQuery.Create(&model.AppToken{
					AppID:    int64(app.ID),
					Token:    token,
					TokenExp: exp,
				}); err != nil {
					response.FailedResponse(c, nil, err.Error())
					return
				}
				response.SuccessResponse(c, token)
				return
			} else {
				response.FailedResponse(c, nil, err.Error())
				return
			}
		}
		response.SuccessResponse(c, token)
	} else {
		response.FailedResponse(c, nil, "应用名称或者密码错误")
		return
	}
}

// AppAccountBind
// @Tags      业务应用用户账号存储
// @Summary   业务应用用户账号存储
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.AppAccountBindRequest true  "请求参数"
// @Success   200   {object}  response.Response{data=string}  "业务应用用户账号存储"
// @Router    /v1/app/account/bind [post]
func AppAccountBind(c *gin.Context) {
	var req request.AppAccountBindRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//
	userID, exist := model.UserInfoQuery.ExistByDID(req.DID)
	if exist == false {
		response.FailedResponse(c, nil, "数据库不存在该DID")
		return
	}
	jsonAccounts, err := json.Marshal(req.AppCounts)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	if err = model.UserAccountQuery.Create(&model.UserAccount{

		UserID:  int64(userID),
		Account: string(jsonAccounts),
		Approve: "",
	}); err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, nil)
}

// AppInfo
// @Tags      业务应用列表
// @Summary   业务应用列表
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.AppInfoRequest true  "请求参数"
// @Success   200   {object}  response.Response{data=[]response.SingleAppInfo}  "业务应用用户账号存储"
// @Router    /v1/app/info [post]
func AppInfo(c *gin.Context) {
	var req request.AppInfoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Debug(1)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	list, err := model.AppInfoQuery.List(req.PageNum, req.PageSize)
	if err != nil {
		log.Debug(1)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	data := make([]response.SingleAppInfo, 0, 1)
	err = copier.Copy(&data, &list)
	if err != nil {
		log.Debug(1)
		log.Debug(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, data)
}

// AppStat
// @Tags 	  查询业务应用信息
// @Summary   查询业务应用信息
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.AppStatRequest true  "查询业务应用信息"
// @Success   200   {object}  response.Response{data=[]response.SingleAppInfo}  "查询业务应用信息"
// @Router    /v1/app/stat [post]
func AppStat(c *gin.Context) {
	var req request.AppStatRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//if
}

// AppUserCA
// @Tags 	  查询用户证书
// @Summary   查询用户证书
// @accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param Bear header string true "token"
// @Param data body request.AppUserCARequest true  "查询用户证书"
// @Success   200   {object}  response.Response{data=response.SingleCA}  "查询用户证书"
// @Router    /v1/app/user/ca [post]
func AppUserCA(c *gin.Context) {
	var req request.AppUserCARequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	userID, exist := model.UserInfoQuery.ExistByDID(req.DID)
	if exist == false {
		response.FailedResponse(c, nil, "数据库不存在该DID")
		return
	}
	uca, err := model.UserCaQuery.FindByUserID(userID)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	//根据用户ID获取用户信息
	uinfo, err := model.UserInfoQuery.FindByID(uint(uca.UserID))
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	log.Debugf("用户证书:%v\n", uca)
	dta := response.SingleCA{}
	err = copier.Copy(&dta, uca)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	err = copier.Copy(&dta, uinfo)
	if err != nil {
		response.FailedResponse(c, nil, err.Error())
		return
	}
	response.SuccessResponse(c, dta)
}
