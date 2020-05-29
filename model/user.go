package model

type UserData struct{
	UserId string `json:"userid"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	Token string `json:"token"`
	LatestLogin string `json:"latestlogin"`
}

type UserDatas []UserData