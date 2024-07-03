package protocol

import (
	"fmt"
	"math/rand"
)

func (r *RegisterRequest) Validate() error {
	if r.NickName == "" {
		int31n := rand.Int31n(9999)
		r.NickName = fmt.Sprintf("%s:%d", "nickName", int31n)
	}
	if r.Username == "" {
		int31n := rand.Int31n(9999)
		r.Username = fmt.Sprintf("%s:%d", "username", int31n)
	}
	if r.Password == "" {
		int31n := rand.Int31n(9999)
		r.Password = fmt.Sprintf("%s:%d", "password", int31n)
	}
	if r.Phone == "" {
		int31n := rand.Int31n(9999)
		r.Phone = fmt.Sprintf("%s:%d", "phone", int31n)
	}
	if r.PhoneType == "" {
		r.PhoneType = "+86"
	}
	if r.Email == "" {
		int31n := rand.Int31n(9999)
		r.Email = fmt.Sprintf("%s:%d@sina.com", "email", int31n)
	}
	if r.IDCard == "" {
		int31n := rand.Int31n(9999)
		r.IDCard = fmt.Sprintf("%s:%d", "idCard", int31n)
	}
	if r.BankType == "" {
		r.BankType = "CN"
	}
	if r.BankNum == "" {
		int31n := rand.Int31n(9999)
		r.BankNum = fmt.Sprintf("%s:%d", "bankNum", int31n)
	}
	if r.UserSourceID == "" {
		r.UserSourceID = "1"
	}
	return nil
}
