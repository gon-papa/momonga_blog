package model

import "time"


type Blog struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UUID      string `gorm:"size:36" json:"uuid"`
	Year	  int    `gorm:"size:4" json:"year"`
	Month	  int    `gorm:"size:2" json:"month"`
	Day		  int    `gorm:"size:2" json:"day"`
	Title	 string `gorm:"size:255" json:"title"`
	Body	  string `gorm:"type:text" json:"body"`
	IsShow	 bool   `gorm:"default:true" json:"is_show"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"delete_at"`

	Tags []Tag `gorm:"many2many:blog_tags" json:"tags"`
}