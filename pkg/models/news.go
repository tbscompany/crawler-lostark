package models

import (
	"crawler-lostark/pkg/exception"
	"time"

	"gorm.io/gorm"
)

type News struct {
	Title string `gorm:"primary_key" json:"title"`
	Date  string `json:"date"`
	URL   string `json:"url"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type NewsRepository interface {
	GetAllNews() ([]*News, *exception.AppError)
	GetNewsByTitle(title string) (*News, *exception.AppError)
	CountByTitle(title string) (int64, *exception.AppError)
	CreateNews(news *News) *exception.AppError
}
