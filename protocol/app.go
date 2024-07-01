package protocol

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
