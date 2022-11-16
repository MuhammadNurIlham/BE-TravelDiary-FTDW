package usersdto

import "time"

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name" validate:"required"`
	Email     string    `json:"email" form:"email" validate:"required"`
	Password  string    `json:"password" form:"password" validate:"required"`
	Phone     string    `gorm:"type: varchar(255)" form:"phone" json:"phone" validate:"required"`
	Address   string    `json:"address" gorm:"type: varchar(255)" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
