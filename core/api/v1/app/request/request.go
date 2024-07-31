package request

type AppLoginRequest struct {
	AppName     string `json:"app_name"`
	AppAccount  string `json:"app_account"`
	AppPassword string `json:"app_password"`
}
