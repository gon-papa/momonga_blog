package model

import "time"


type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      string    `gorm:"size:36" json:"uuid"`
	Name      string    `gorm:"size:255" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Blogs	 []Blog    `gorm:"many2many:blog_tags" json:"blogs"`
}