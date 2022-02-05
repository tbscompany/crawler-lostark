package models

import (
	"crawler-lostark/pkg/database"
	"crawler-lostark/pkg/exception"
)

type NewsRepositoryDb struct {
	client database.DbConnection
}

func NewNewsRepository() NewsRepositoryDb {
	return NewsRepositoryDb{
		client: database.GetDbConnection(),
	}
}

func (r NewsRepositoryDb) GetAllNews() ([]*News, *exception.AppError) {
	news := make([]*News, 0)
	err := r.client.DB.Find(&news).Error
	return news, exception.NewUnexpectedError("Unable to get all news", err)
}

func (r NewsRepositoryDb) GetNewsByTitle(title string) (*News, *exception.AppError) {
	c := new(News)
	err := r.client.DB.Debug().Where("title = ?", title).First(&c).Error
	return c, exception.NewUnexpectedError("Unable to get news by title", err)
}

func (r NewsRepositoryDb) CountByTitle(title string) (int64, *exception.AppError) {
	var count int64
	err := r.client.DB.Model(&News{}).Debug().Where("title = ?", title).Count(&count).Error
	return count, exception.NewUnexpectedError("Unable to count news by title", err)
}

func (r NewsRepositoryDb) CreateNews(news *News) *exception.AppError {
	err := r.client.DB.Create(news).Error
	return exception.NewUnexpectedError("Unable to create news", err)
}
