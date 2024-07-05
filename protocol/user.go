package protocol

import "time"

type RegisterRequest struct {
	//真实姓名
	Username     string `json:"user_name"`
	NickName     string `json:"nick_name"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	PhoneType    string `json:"phone_type"`
	Email        string `json:"email"`
	IDCard       string `json:"id_card"`
	BankType     string `json:"bank_type"`
	BankNum      string `json:"bank_num"`
	UserSourceID string `json:"user_source_id"`
}
type LoginRequest struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

type ModifyRequest struct {
	PhoneType string `json:"phone_type"`
	PhoneNum  string `json:"phone_num"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IdCard    string `json:"id_card"`
	BankType  string `json:"bank_type"`
	BankNum   string `json:"bank_num"`
	NickName  string `json:"nick_name"`
	// UserName  string `json:"user_name"`
}

type ListRequest struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

// SingleUser Response
type SingleUser struct {
	NickName   string    `json:"nick_name"`
	PhoneType  string    `json:"phone_type"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"register_time"`
}

type GetRequest struct {
	FuzzyUsername string `json:"fuzzy_username"`
}

// SpaceInfo Response
type SpaceInfo struct {
	Total          int64   `json:"total"`
	UseSpace       float64 `json:"use_space"`
	AvailableSpace float64 `json:"available_space"`
	SpaceStatus    string  `json:"space_status"`
}
type SpaceExpandRequest struct {
	ExpandSize int64 `json:"expand_size"`
}
