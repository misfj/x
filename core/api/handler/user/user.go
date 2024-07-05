package user

import (
	"coredx/config"
	"coredx/db"
	v1 "coredx/db/model/v1"
	"coredx/log"
	"coredx/protocol"
	"coredx/utils"
	"time"

	"github.com/wumansgy/goEncrypt/rsa"
)

func Register(req *protocol.RegisterRequest) error {
	pwd := req.Password
	bcrypt := utils.Bcrypt(pwd)
	///todo 设置用户默认空间
	kp, err := rsa.GenerateRsaKeyBase64(2048)
	if err != nil {
		return err
	}
	nickName := req.NickName
	public := kp.PublicKey
	private := kp.PrivateKey
	publicMD5 := utils.MD5([]byte(public))
	privateMd5 := utils.MD5([]byte(private))
	//只有在客户选择sm2才需要存储密钥
	storeKey := ""
	algorithm := "rsa"
	timeStamp := time.Now().Unix()
	//检验是否合法
	err = utils.IslegalError(public, private)
	if err != nil {
		return err
	}
	userID, err := v1.AuthUserCreate(db.GDB, nickName,
		req.Username, bcrypt, req.Phone, req.PhoneType, req.Email, req.IDCard,
		req.BankType, req.BankNum, req.UserSourceID)
	if err != nil {
		log.Debug(1)
		return err
	}
	// 添加用户凭证
	_, err = v1.AuthUserTokenCreate(db.GDB, userID, publicMD5, time.Now().Add(time.Hour*time.Duration(config.JwtConfig().Exp)*24), "正常")
	if err != nil {
		log.Debug(2)
		return err
	}
	//用户证书添加d
	if err = v1.UserCaCreate(db.GDB, userID,
		public, private, publicMD5, privateMd5, storeKey, algorithm, timeStamp); err != nil {
		log.Debug(3)
		return err
	}
	//设置用户可用空间
	if err = v1.AuthUserSpaceCreate(db.GDB, userID,
		config.StoreConfig().DefaultSize, 0, float64(config.StoreConfig().DefaultSize)); err != nil {
		log.Debug(4)
		return err
	}
	return err

}
func Login(req *protocol.LoginRequest) (string, error) {

	userId, err := v1.AuthUserVerifyByNickNameAndPassword(db.GDB, req.NickName, req.Password)
	if err != nil {
		return "", err
	}
	accessKey, err := v1.AuthUserTokenFindUserAccessByUserId(db.GDB, userId)
	if err != nil {
		return "", err
	}
	return accessKey, nil
}

func Modify(req *protocol.ModifyRequest, accessKey string) error {
	//更新
	userID, err := v1.AuthUserTokenGetUserIdByAccessKey(db.GDB, accessKey)
	if err != nil {
		return err
	}
	phoneType := req.PhoneType
	phoneNum := req.PhoneNum
	email := req.Email
	password := req.Password
	idCard := req.IdCard
	bankType := req.BankType
	bankNum := req.BankNum
	nickName := req.NickName
	err = v1.AuthUserModifyByUserId(db.GDB, phoneType, phoneNum, email, password, idCard, bankType, bankNum, nickName, userID)
	return err
}

func Delete(accessKey string) error {
	//根据accessKey获取user_id
	userId, err := v1.AuthUserTokenGetUserIdByAccessKey(db.GDB, accessKey)
	if err != nil {
		return err
	}

	//设计到更新字段的表有auth_user,auth_user_token,user_ca,auth_user_space将is_delete的字段修改为1
	err = v1.AuthUserDeleteByUserId(db.GDB, userId)
	if err != nil {
		return err
	}
	err = v1.AuthUserTokenDeleteByUserId(db.GDB, userId)
	if err != nil {
		return err
	}
	err = v1.UserCaDeleteByUserIds(db.GDB, userId)
	if err != nil {
		return err
	}
	return v1.AuthUserSpaceDeleteByUserId(db.GDB, userId)
}

func List(pageSize int, pageNum int) ([]*protocol.SingleUser, error) {
	return v1.AuthUserListByPageSize(db.GDB, pageNum, pageSize)
}
func Get(fuzzyName string, batchSize int) ([]*protocol.SingleUser, error) {
	return v1.AuthUserFindMatchByFuzzyName(db.GDB, fuzzyName, batchSize)
}

func SpaceInfo(accessKey string) (*protocol.SpaceInfo, error) {
	userId, err := v1.AuthUserTokenGetUserIdByAccessKey(db.GDB, accessKey)
	if err != nil {
		return nil, err
	}
	return v1.AuthUserSpaceGetInfoByUserId(db.GDB, userId)
}
func SpaceExpand(accessKey string, expandSize int) (*protocol.SpaceInfo, error) {
	userId, err := v1.AuthUserTokenGetUserIdByAccessKey(db.GDB, accessKey)
	if err != nil {
		return nil, err
	}
	return v1.AuthUserSpaceExpandByUserId(db.GDB, userId, expandSize)
}
func Upgrade(accessKey string) error {
	//did 生成规则(私钥的Sha256)
	v1.AuthUserUpgradeByUserId()
}
