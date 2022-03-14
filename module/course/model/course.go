package coursemodel

import "time"

type Course struct {
	ID          int        `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description" gorm:"column:description"`
	ImageUrl    string     `json:"imageUrl" gorm:"column:imageUrl"`
	Price       float64    `json:"price" gorm:"column:price"`
	Status      string     `json:"status" gorm:"column:status"`
	UserId      int        `json:"userId" gorm:"column:user_id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
