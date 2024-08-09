package global

import (
	"bytes"
	"coredx/client"
	"coredx/db/dal/model"
	"coredx/log"
	"coredx/utils"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/wumansgy/goEncrypt/aes"
	"io"
	"os"
	"time"
)

const (
	Open   = "1"
	NoOpen = "2"
)

var Global *global

// Global 全局的Channel支持业务应用进行回调
type global struct {
	ThirtyMB      int64
	UploadCh      chan UploadMeta
	UploadErrCh   chan *UploadErrMeta
	CryptoCh      chan *CryptoMeta
	CryptoChError chan *CryptoErrMeta
	UdpCh         chan []byte
	CCh           chan int
}
type CryptoMeta struct {
	Content           []byte
	FileName          string `json:"file_name"`
	FileSize          int64  `json:"file_size"`
	FileMd5           string `json:"file_md5"`
	FileKey           string `json:"file_key"`
	FileAlgo          string `json:"file_algo"`
	FileEncryptedHash string `json:"file_encrypted_hash"`
	FileUrl           string `json:"file_url"`
	NodeID            string `json:"node_id"`
	TaskID            string `json:"task_id"`
	CallbackUrl       string `json:"callback_url"`
}
type CryptoErrMeta struct {
	FileName          string `json:"file_name"`
	FileSize          int64  `json:"file_size"`
	FileMd5           string `json:"file_md5"`
	FileKey           string `json:"file_key"`
	FileAlgo          string `json:"file_algo"`
	FileEncryptedHash string `json:"file_encrypted_hash"`
	FileUrl           string `json:"file_url"`
	NodeID            string `json:"node_id"`
	TaskID            string `json:"task_id"`
	CallbackUrl       string `json:"callback_url"`
	ErrMsg            string `json:"err_msg"`
}
type UploadMeta struct {
	UserID      uint      `json:"user_id"`
	AppID       uint      `json:"app_id"`
	Content     []byte    `json:"content"`
	DataType    string    `json:"data_type"`
	FileUrl     string    `json:"file_url"`
	FileSize    int64     `json:"file_size"`
	NodeID      string    `json:"node_id"`
	FileName    string    `json:"file_name"`
	FileMd5     string    `json:"file_md5"`
	CallbackUrl string    `json:"callback_url"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TaskID      string    `json:"task_id"`
	Remark      string    `json:"remark"`
}
type UploadErrMeta struct {
	FileUrl     string `json:"file_url"`
	FileSize    int64  `json:"file_size"`
	NodeID      string `json:"node_id"`
	FileName    string `json:"file_name"`
	FileMd5     string `json:"file_md5"`
	CallbackUrl string `json:"callback_url"`
	CostMill    int64  `json:"cost_mill"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	TaskID      string `json:"task_id"`
	ErrMsg      string `json:"err_msg"`
}

func InitGlobal(num int) {
	Global = &global{}
	//初始化UDP Channel
	Global.UdpCh = make(chan []byte, 4096)
	Global.UploadCh = make(chan UploadMeta, num)
	Global.UploadErrCh = make(chan *UploadErrMeta, num)
	Global.CryptoCh = make(chan *CryptoMeta, num)
	Global.CryptoChError = make(chan *CryptoErrMeta, num)
	Global.ThirtyMB = 30 * 1024 * 1024
	Global.CCh = make(chan int, 20)
	//Global = &global{
	//	UploadCh:    make(chan *UploadMeta, num),
	//	UploadErrCh: make(chan *UploadErrMeta, num),
	//	ThirtyMB:    30 * 1024 * 1024,
	//}
}

func (g *global) Start() {
	for {
		select {
		case v := <-g.UploadCh:
			log.Debug("Channel Recv Info")
			start := time.Now().UnixMilli()
			res, err := client.Client.HttpC.R().SetBody(map[string]interface{}{
				"file_url":     v.FileUrl,
				"file_size":    v.FileSize,
				"node_id":      v.NodeID,
				"file_name":    v.FileName,
				"file_md5":     v.FileMd5,
				"callback_url": v.CallbackUrl,
				"start_time":   v.StartTime,
				"end_time":     v.EndTime,
				"task_id":      v.TaskID,
			}).Post(v.CallbackUrl)
			end := time.Now().UnixMilli()
			errMsg := ""
			if err != nil {
				errMsg = err.Error()
			}
			filer, err := os.Create(v.FileName)
			if err != nil {
				log.Error("create file failed, err:", err)
				continue
			}
			io.Copy(filer, bytes.NewReader(v.Content))
			v.Content = nil
			reqContent, _ := json.Marshal(v)
			//记录日志app_log
			err = model.AppLogQuery.Create(&model.AppLog{
				AppID:       int64(v.AppID),
				ServiceType: "上传回调",
				ReqURL:      v.CallbackUrl,
				Method:      "POST",
				RemoteIP:    v.CallbackUrl,
				ReqContent:  string(reqContent),
				ResMsg:      string(res.Body()),
				ResSize:     int32(len(res.Body())),
				ResErr:      errMsg,
				Cost:        end - start,
			})
			if err != nil {
				//todo 将错误日志写入数据库(回调日志)
				log.Error("http post to callback url failed, err:%v\n", err)
				uploadErrMeta := &UploadErrMeta{}
				err = copier.Copy(uploadErrMeta, v)
				if err != nil {
					log.Error("copy data failed, err:", err)
					continue
				}
				uploadErrMeta.ErrMsg = err.Error()
				g.UploadErrCh <- uploadErrMeta
			}
			//记录数据信息到数据库data_info
			err = model.DataInfoQuery.Create(&model.DataInfo{
				UseID:          int64(v.UserID),
				DataID:         v.FileMd5,
				DataName:       v.FileName,
				DataDiectoryID: 1,
				DataType:       v.DataType,
				IsOpen:         Open,
				DataSize:       v.FileSize,
				DataPos:        v.FileUrl,
				NodeID:         v.NodeID,
			})
			if err != nil {
				log.Error(err)
			}
			//log.Info("http post to callback url success, response status code:", res.Body())
		case errV := <-g.UploadErrCh:
			//todo 将返回日志写入数据库(是否成功)
			log.Error("http post to callback url error:%s", errV.ErrMsg)
		case vx := <-g.CCh:
			log.Infof("http post to callback url:%d", vx)
		case chCrypto := <-g.CryptoCh:
			//保存文件
			cipherText, err := aes.AesCbcEncrypt(chCrypto.Content, []byte(AES128Default), []byte(AES128Default))
			if err != nil {
				log.Error("create file failed, err:", err)
				break
			}
			filer, err := os.Create(chCrypto.FileUrl)
			if err != nil {
				log.Error("create file failed, err:", err)
				break
			}
			io.Copy(filer, bytes.NewReader(cipherText))
			encMd5 := utils.MD5(cipherText)
			err = model.DataEncQuery.Create(&model.DataEnc{
				FileEncryptedHash: encMd5,
				FileName:          chCrypto.FileName,
				FileSize:          chCrypto.FileSize,
				FileMd5:           chCrypto.FileMd5,
				FileKey:           chCrypto.FileKey,
				FileAlgo:          chCrypto.FileAlgo,
				//FileEncryptedHash: chCrypto.FileEncryptedHash,
				FileURL: chCrypto.FileUrl,
				NodeID:  chCrypto.NodeID,
			})
			if err != nil {
				g.CryptoChError <- &CryptoErrMeta{
					FileName:          chCrypto.FileName,
					FileSize:          chCrypto.FileSize,
					FileMd5:           chCrypto.FileMd5,
					FileKey:           chCrypto.FileKey,
					FileAlgo:          chCrypto.FileAlgo,
					FileEncryptedHash: chCrypto.FileEncryptedHash,
					FileUrl:           chCrypto.FileUrl,
					NodeID:            chCrypto.NodeID,
					ErrMsg:            err.Error(),
				}
			}
			//todo 添加回调逻辑
			res, err := client.Client.HttpC.R().SetBody(map[string]interface{}{
				"file_url":     chCrypto.FileUrl,
				"file_size":    chCrypto.FileSize,
				"node_id":      chCrypto.NodeID,
				"file_name":    chCrypto.FileName,
				"file_md5":     encMd5,
				"callback_url": chCrypto.CallbackUrl,
				//"start_time":   chCrypto.StartTime,
				//"end_time":     v.EndTime,
				"task_id": chCrypto.TaskID,
			}).Post(chCrypto.CallbackUrl)
			if err != nil {
				log.Error("http post to callback url failed, err:", err)
				break
			}
			log.Debugf("加密回调响应:%s", string(res.Body()))
		case chCryptoErr := <-g.CryptoChError:
			log.Errorf("将加密数据信息写入数据库失败:%v", chCryptoErr)
		default:

		}
	}
}
