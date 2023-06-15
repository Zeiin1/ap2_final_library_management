package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	SurName  string `json:"sur_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
