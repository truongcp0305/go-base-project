package model

type User struct {
	UserName    string `json:"userName" form:"userName" bson:"userName" example:"truong"`
	Password    string `json:"password" form:"password" bson:"password" example:"1234"`
	DisplayName string `json:"displayName" form:"displayName" bson:"displayName" example:"456"`
	Token       string `json:"token" bson:"token" example:""`
}
