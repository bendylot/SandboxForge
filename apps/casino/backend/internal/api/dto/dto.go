package dto

type RegisterReq struct {
	Login    string `json:"login"    validate:"required,alphanum,min=3,max=24"`
	Email    string `json:"email"    validate:"required,email,max=254"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
