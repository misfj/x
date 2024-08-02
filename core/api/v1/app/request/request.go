package request

type AppLoginRequest struct {
	AppName     string `json:"app_name"`
	AppAccount  string `json:"app_account"`
	AppPassword string `json:"app_password"`
}

type AppRegisterRequest struct {
	AppName     string `json:"app_name"`
	AppUser     string `json:"app_user"`
	AppPassword string `json:"app_password"`
	AppPhone    string `json:"app_phone"`
	AppType     string `json:"app_type"`
	Algo        string `json:"algo"`
	StoreKey    string `json:"store_key"`
}

type AppUserRegisterRequest struct {
	//加上AppName 防止篡token与业务应用,在Service层进行校验
	AppName        string `json:"app_name"`
	UserName       string `json:"user_name"`
	NickName       string `json:"nick_name"`
	UserLevel      string `json:"user_level"`
	UserPhone      string `json:"user_phone"`
	UserSourceType string `json:"user_source_type"`
	UserSource     string `json:"user_source"`
	UserType       string `json:"user_type"`
}
type AppTestRequest struct {
	Name string `json:"name"`
}

type singleCount struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type AppAccountBindRequest struct {
	AppName   string        `json:"app_name"`
	DID       string        `json:"did"`
	AppCounts []singleCount `json:"app_counts"`
}

type AppInfoRequest struct {
	AppName  string `json:"app_name"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

type AppStatRequest struct {
	AppName string `json:"app_name"`
}
type AppUserCARequest struct {
	AppName string `json:"app_name"`
	DID     string `json:"did"`
}
