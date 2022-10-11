package request

// SignUp 注册请求参数
type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// Login 登录请求参数
type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
