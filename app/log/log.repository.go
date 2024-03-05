package log

import (
	"github.com/jihanlugas/pandora/model"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetDataBefore(conn *gorm.DB, time time.Time) ([]model.Log, error)
	Create(conn *gorm.DB, data model.Log) error
	Delete(conn *gorm.DB, data model.Log) error
	Deletes(conn *gorm.DB, data []model.Log) error
}

type repository struct {
}

func (r repository) GetDataBefore(conn *gorm.DB, time time.Time) ([]model.Log, error) {
	var err error
	var data []model.Log

	err = conn.Where("create_dt < ?", time).Find(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Log) error {
	return conn.Create(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Log) error {
	return conn.Delete(&data).Error
}

func (r repository) Deletes(conn *gorm.DB, data []model.Log) error {
	return conn.Delete(&data).Error
}

func NewRepository() Repository {
	return repository{}
}
