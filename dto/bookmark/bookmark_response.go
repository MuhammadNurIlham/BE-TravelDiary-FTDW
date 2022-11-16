package bookmarkdto

import "time"

type BookmarkResponse struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int       `json:"user_id" form:"user_id" gorm:"int"`
	JourneyID int       `json:"journey_id" form:"journey_id" gorm:"int"`
	CreatedAt time.Time `json:"created_at"`
}
