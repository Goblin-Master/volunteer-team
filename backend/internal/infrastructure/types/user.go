package types

type LoginReq struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	LoginType string `json:"login_type"`
}

type LoginResp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
