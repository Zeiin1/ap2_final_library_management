package models

type StockDecreaseLog struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	OrderId   int64 `json:"order_id"`
	BookRefer int64 `json:"book_id"`
}
