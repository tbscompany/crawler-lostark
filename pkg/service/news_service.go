package service

import (
	"crawler-lostark/pkg/exception"
	"crawler-lostark/pkg/models"
)

type NewsService interface {
	GetAllNews() ([]*models.News, *exception.AppError)
	GetNewsByTitle(title string) (*models.News, *exception.AppError)
	CountByTitle(title string) (int64, *exception.AppError)
	CreateNews(news *models.News) *exception.AppError
}

type DefaultNewsService struct {
	repository models.NewsRepository
}

func NewNewsService(repository models.NewsRepository) DefaultNewsService {
	return DefaultNewsService{repository: repository}
}

func (s DefaultNewsService) GetAllNews() ([]*models.News, *exception.AppError) {
	return s.repository.GetAllNews()
}

func (s DefaultNewsService) CountByTitle(title string) (int64, *exception.AppError) {
	return s.repository.CountByTitle(title)
}

func (s DefaultNewsService) GetNewsByTitle(title string) (*models.News, *exception.AppError) {
	return s.repository.GetNewsByTitle(title)
}

func (s DefaultNewsService) CreateNews(news *models.News) *exception.AppError {
	return s.repository.CreateNews(news)
}
