package dtos

type UserLogin struct {
	Username string `json:"username" description:"账号"`
	Password string `json:"password" description:"密码"`
	Code     string `json:"code" description:"验证码"`
	LoginIp  string `json:"login_ip"`
}
