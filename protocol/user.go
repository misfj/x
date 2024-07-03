package protocol

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
