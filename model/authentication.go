package model

type Init struct {
	UserID   string `json:"userId"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	NoHp     string `json:"noHp"`
	Fullname string `json:"fullname"`
	PhotoUrl string `json:"photoUrl"`
}
