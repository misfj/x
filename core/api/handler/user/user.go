package user

import (
	"coredx/db"
	v1 "coredx/db/model/v1"
	"coredx/protocol"
	"coredx/utils"
)

func Register(req *protocol.RegisterRequest) error {
	pwd := req.Password
	bcrypt := utils.Bcrypt(pwd)
	///todo 设置用户默认空间
	return v1.AuthUserCreate(db.GDB, req.NickName,
		req.Username, bcrypt, req.Phone, req.PhoneType, req.Email, req.IDCard,
		req.BankType, req.BankNum, req.UserSourceID)
}
