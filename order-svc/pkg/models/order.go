package models

type Order struct {
	Id     int64 `json:"id" gorm:"primaryKey"`
	Price  int64 `json:"price"`
	BookId int64 `json:"book_id"`
	UserId int64 `json:"user_id"`
}
