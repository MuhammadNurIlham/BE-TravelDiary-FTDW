package journeydto

import "time"

type JourneyRequest struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type:text" form:"description"`
	Image       string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID      int       `json:"user_id" form:"user_id" gorm:"type: int"`
	Books       string    `json:"books"`
	CreatedAt   time.Time `json:"created_at"`
}

type JourneyUpdateRequest struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type:text" form:"description"`
	Image       string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID      int       `json:"user_id" form:"user_id" gorm:"type: int"`
	Books       string    `json:"books"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
