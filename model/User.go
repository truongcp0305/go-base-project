package model

type User struct {
	UserName    string `json:"userName" form:"userName"`
	Password    string `json:"password" form:"password" `
	DisplayName string `json:"displayName" form:"displayName"`
	Token       string `json:"token"`
}
