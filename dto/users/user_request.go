package usersdto

import "time"

type CreateUserRequest struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	Email     string    `json:"email" form:"email" validate:"required"`
	Password  string    `json:"password" form:"password" validate:"required"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)" form:"phone" validate:"required"`
	Address   string    `json:"address" gorm:"type: varchar(255)" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Phone     string    `json:"phone" form:"phone"`
	Address   string    `json:"address" form:"address"`
	CreatedAt time.Time `json:"created_at"`
}
