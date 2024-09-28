package model

import "time"


type Users struct {
	ID           uint      `gorm:"primaryKey" json:"id"`             
	UUID         string    `gorm:"size:36" json:"uuid"`              
	RefreshToken *string   `gorm:"size:255" json:"refresh_token"`    
	TokenExpiry  *time.Time `json:"token_expiry"`           
	Active	   bool      `json:"active"`         
	UserID       string    `gorm:"size:255" json:"user_id"`          
	Password     string    `gorm:"size:255" json:"password"`         
	CreatedAt    time.Time  `json:"created_at"`                      
	UpdatedAt    time.Time  `json:"updated_at"`    
}