package handler

import (
	"bytes"
	"context"
	"coredx/client"
	"coredx/core/api/middware"
	"coredx/core/api/response"
	"coredx/db"
	"coredx/db/model"
	"coredx/log"
	"coredx/store/minio"
	"coredx/utils"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	upload = "upload"
)

func AppLogin(exp int, iss, sub, username, email string) (code int, msg string, data string) {
	//
	if username == "internal" && email == "internal" {
		token, err := middware.GenerateToken(exp, iss, sub, username)
		if err != nil {
			log.Error(err)
			return http.StatusInternalServerError, err.Error(), ""
		}
		return http.StatusOK, response.Success, token
	}
	//查询数据库
	_, err := model.SysUserFindByUsername(db.GDB, username)
	if err != nil {
		return http.StatusInternalServerError, err.Error(), ""
	}
	token, err := middware.GenerateToken(exp, iss, sub, username)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, err.Error(), ""
	}
	return http.StatusOK, response.Success, token
}

func DRUserRegister(useName string, password string) (code int, msg string, data string) {
	res, err := client.Client.HttpC.R().SetQueryParams(map[string]string{
		"userName": useName,
		"password": password,
	}).Post(client.Client.DrApi.DrUserCreateUrl)
	if err != nil {
		return http.StatusInternalServerError, err.Error(), ""
	}
	body := res.Body()
	type x struct {
		Data struct {
			Id    int    `json:"id"`
			Token string `json:"token"`
		} `json:"data"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Total   string `json:"total"`
	}
	var xx x
	err = json.Unmarshal(body, &xx)
	if err != nil {
		return http.StatusInternalServerError, err.Error(), ""
	}
	//todo 将数权ID写入数据库
	return http.StatusOK, response.Success, xx.Data.Token
}

// Upload 支持form表单和API上传用户数据
func Upload(did string, filename string, x []byte) (code int, msg string, data string) {
	useGB, err := model.GetTotalUseGB(db.GDB, did)
	log.Debug(useGB)
	md5 := utils.MD5(x)
	//去数据查询每个did用户的容量是否小于1G
	//todo 5G以下的数据返回是etag是md5
	//用did作为同名
	etag, err := minio.Put(context.Background(), upload, filename, bytes.NewReader(x), int64(len(x)))
	if err != nil {
		return http.StatusInternalServerError, err.Error(), ""
	}
	if strings.EqualFold(md5, etag) {
		return http.StatusOK, response.Success, etag
	}
	//删除问价
	return http.StatusOK, response.Success, etag
}
