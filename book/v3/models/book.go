package models

type BookSet struct {
	Total int64   `json:"total"`
	Items []*Book `json:"items"`
}

type Book struct {
	Id uint `json:"id" gorm:"primaryKey;column:id"`

	BookSpec
}
type BookSpec struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price" validate:"required"`
	IsSale bool    `json:"is_sale" gorm:"column:is_sale"`
}

func (b *Book) TableName() string {
	return "books"
}
