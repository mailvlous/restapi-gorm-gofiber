package entity

import (
    "time"
    "gorm.io/gorm"
)

type Books struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Author string `json:"author"`
	Cover string `json:"cover"`
	CreatedAt time.Time `json:"created_at",gorm:"autoUpdateTime:nano"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at; gorm:"index,column:deleted_at"`
}