package types

type LoginReq struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Code      string `json:"code"`
	LoginType string `json:"login_type"`
}

type LoginResp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type GetCodeReq struct {
	Email string `json:"email"`
}
type GetCodeResp struct {
	Code string `json:"code"`
}
