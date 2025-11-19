package api

type RegisterReq struct {
	Login    string `json:"login" binding:"required,min=2,max=24"`
	Password string `json:"password" binding:"required,min=6,max=24"`
}

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
