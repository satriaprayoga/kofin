package auth

type Login struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	Username string   `json:"userName"`
	Role     []string `json:"authority"`
}

type LoginResponse struct {
	UserInfo UserInfo `json:"user"`
	Token    string   `json:"token"`
}
