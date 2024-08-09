package service

import (
	"bytes"
	"coredx/config"
	"coredx/core/api/v1/app/request"
	"coredx/core/api/v1/app/response"
	"coredx/db/dal/model"
	"coredx/global"
	"coredx/log"
	"coredx/utils"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/osext"
	"github.com/wumansgy/goEncrypt/aes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//todo 添加raw body json接口

func CryptoXX(c *gin.Context) {
	execDir, err := osext.ExecutableFolder()
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	contentType := c.Request.Header.Get("Content-Type")
	var req request.CryptoRequest
	err = c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		response.FailedResponse(c, nil, err.Error())
		return
	}
	fileBody := make([]byte, 0, 2*1024*1024)
	var filerr error
	if strings.Contains(contentType, "multipart/form-data") {
		file, err := c.FormFile("file")
		if err != nil {
			log.Error(err)
			response.FailedResponse(c, nil, err.Error())
			return
		}
		fileReader, err := file.Open()
		if err != nil {
			log.Error(err)
			response.FailedResponse(c, nil, err.Error())
			return
		}

		//warning:io.ReadAll
		fileBody, filerr = io.ReadAll(fileReader)
		if filerr != nil {
			log.Error(err)
			response.FailedResponse(c, nil, err.Error())
			return
		}
		//代表加密
		if req.Action == 1 {
			//判断存放文件的目录是否存在,不存在就创建
			_, err = os.Stat(filepath.Join(execDir, global.UploadsDir))
			if err != nil {
				if os.IsNotExist(err) {
					//创建文件夹
					err = os.Mkdir(filepath.Join(execDir, global.UploadsDir), 0666)
					if err != nil {
						log.Error(err)
						response.FailedResponse(c, nil, err.Error())
						return
					}
				} else {
					response.FailedResponse(c, nil, err.Error())
					return
				}
			}
			_, err = os.Stat(filepath.Join(execDir, global.EncryptFileDir))
			if err != nil {
				if os.IsNotExist(err) {
					//创建文件夹
					err = os.Mkdir(filepath.Join(execDir, global.EncryptFileDir), 0666)
					if err != nil {
						log.Error(err)
						response.FailedResponse(c, nil, err.Error())
						return
					}
				} else {
					response.FailedResponse(c, nil, err.Error())
					return
				}
			}
			//todo 文件数据太大,进行异步回调
			cipherText, err := aes.AesCbcEncrypt(fileBody, []byte(global.AES128Default), []byte(global.AES128Default))
			if err != nil {
				log.Error(err)
				response.FailedResponse(c, nil, err.Error())
				return
			}
			fileName := file.Filename

			//保存源文件到uploads目录下面
			filer, err := os.Create(filepath.Join(execDir, global.UploadsDir, fileName))
			defer filer.Close()
			writeN, err := io.Copy(filer, bytes.NewReader(fileBody))
			if err != nil {
				log.Error(err)
				log.Debug(writeN, len(fileBody))
				response.FailedResponse(c, nil, err.Error())
				return
			}
			if writeN != int64(len(fileBody)) {
				log.Debug(writeN, len(fileBody))
				response.FailedResponse(c, nil, "文件拷贝错误")
				return
			}
			fileBodyMd5 := utils.MD5(fileBody)
			//保存加密文件到encrypt-files目录下面
			efiler, err := os.Create(filepath.Join(execDir, global.EncryptFileDir, fileName))
			if err != nil {
				log.Error(err)
				response.FailedResponse(c, nil, err.Error())
				return
			}
			defer efiler.Close()
			writeN, err = io.Copy(efiler, bytes.NewReader(cipherText))
			if err != nil {
				log.Error(err)
				response.FailedResponse(c, nil, err.Error())
				return
			}
			if writeN != int64(len(cipherText)) {
				log.Error(err)
				log.Debug(writeN, len(fileBody))
				response.FailedResponse(c, nil, "文件拷贝错误")
				return
			}
			cipherTextMD5 := utils.MD5(cipherText)
			//信息写入数据库
			if err = model.DataEncQuery.Create(&model.DataEnc{
				FileName:          file.Filename,
				FileSize:          file.Size,
				FileMd5:           fileBodyMd5,
				FileKey:           global.AES128Default,
				FileAlgo:          global.AES128,
				FileEncryptedHash: cipherTextMD5,
				FileURL:           filepath.Join(execDir, global.EncryptFileDir, fileName),
				NodeID:            config.GetNode().NodePrivateMd5,
			}); err != nil {
				log.Debug(writeN, len(fileBody))
				response.FailedResponse(c, nil, err.Error())
				return
			}
			response.SuccessResponse(c, nil)
		}
		if req.Action == 2 {
			decrypt, err := aes.AesCbcDecrypt(fileBody, []byte(global.AES128Default), []byte(global.AES128Default))
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			c.Writer.Write(decrypt)
		}
	}
	if strings.Contains(contentType, "application/json") {
		//验证AppName
		fileMd5 := req.DataAID
		//先去数据库查找data_info(1级数据)
		//todo 记得验证数据名称和dataID的验证
		dataInfo, err := model.DataInfoQuery.FindByDataID(fileMd5)
		if err != nil {
			response.FailedResponse(c, nil, err.Error())
			return
		}
		fileReadContent, err := os.ReadFile(dataInfo.DataPos)
		if err != nil {
			log.Error(err)
			response.FailedResponse(c, nil, err.Error())
			return
		}
		fileReadContentMD5 := utils.MD5(fileReadContent)
		if req.Action == 1 {
			cipherText, err := aes.AesCbcEncrypt(fileReadContent, []byte(global.AES128Default), []byte(global.AES128Default))
			if err != nil {
				log.Error(err)
				response.FailedResponse(c, nil, err.Error())
				return
			}
			//保存在本地文件夹
			if err := model.DataEncQuery.Create(&model.DataEnc{
				FileName:          dataInfo.DataName,
				FileSize:          dataInfo.DataSize,
				FileMd5:           fileReadContentMD5,
				FileKey:           global.AES128Default,
				FileAlgo:          global.AES128,
				FileEncryptedHash: utils.MD5(cipherText),
				FileURL:           filepath.Join(execDir, global.EncryptFileDir, dataInfo.DataName),
				NodeID:            config.GetNode().NodePrivateMd5,
			}); err != nil {
				log.Error(err)
				response.FailedResponse(c, nil, err.Error())
				return
			}
			response.SuccessResponse(c, nil)
		}
		if req.Action == 2 {
			//从数据库data_enc 查找文件,然后读取进行解密,返回文件流
			aid, err := model.DataEncQuery.ExistRawByDataAid(req.DataAID, "")
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			//读取文件
			fileContent, err := os.ReadFile(aid.FileURL)
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			symmetricKey := aid.FileKey
			//进行解密,根据具体算法解密
			plainText, err := aes.AesCbcDecrypt(fileContent, []byte(symmetricKey), []byte(symmetricKey))
			if err != nil {
				response.FailedResponse(c, nil, err.Error())
				return
			}
			c.Writer.Write(plainText)
		}

	}
}
