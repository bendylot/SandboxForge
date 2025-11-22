package api

type RegisterReq struct {
	Login    string `json:"login" binding:"required,min=4,max=24"`
	Email    string `json:"email" binding:"required,min=4,max=24"`
	Password string `json:"password" binding:"required,min=4,max=24"`
}

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
