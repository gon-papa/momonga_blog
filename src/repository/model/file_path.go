package model

import "time"

type FilePath struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name	  string `gorm:"size:255" json:"name"`
	FilePath  string `gorm:"size:1024" json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}