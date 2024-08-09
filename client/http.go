package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var Client *client

type client struct {
	HttpC *resty.Client
	DrApi DRApi
}
type DRApi struct {
	AuthenticationUrl   string
	DelDataUrl          string
	DrAccountInfoUrl    string
	DrAccountSyncUrl    string
	DrUserCreateUrl     string
	DrUserModifyUrl     string
	EncryptUrl          string
	GetDrUserInfoUrl    string
	SelectUploadDataUrl string
	UploadDataUrl       string
}

func Init() {
	Client = &client{HttpC: resty.New()}
	Client.HttpC.SetDebug(true).SetHeader("User-Agent", "nwyl-yyds").SetTimeout(time.Second * 30)
	// var dr DRApi
	// dr.AuthenticationUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().Authentication)
	// dr.DelDataUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().DelData)
	// dr.DrAccountInfoUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().DrAccountInfo)
	// dr.DrAccountSyncUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().DrAccountSync)
	// dr.DrUserCreateUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().DrUserCreate)
	// dr.DrUserModifyUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().DrUserModify)
	// dr.EncryptUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().Encrypt)
	// dr.GetDrUserInfoUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().GetDrUserInfo)
	// dr.SelectUploadDataUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().SelectUploadData)
	// dr.UploadDataUrl = fmt.Sprintf("%s/%s", config.DRConfig().EndPoint, config.DRConfig().UploadData)
	// Client.DrApi = dr
}
