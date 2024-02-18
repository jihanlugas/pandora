package user

import (
	"fmt"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"github.com/jihanlugas/pandora/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.User, error)
	GetByUsername(conn *gorm.DB, username string) (model.User, error)
	GetByEmail(conn *gorm.DB, email string) (model.User, error)
	GetByNoHp(conn *gorm.DB, noHp string) (model.User, error)
	GetViewById(conn *gorm.DB, id string) (model.UserView, error)
	GetViewByUsername(conn *gorm.DB, username string) (model.UserView, error)
	GetViewByEmail(conn *gorm.DB, email string) (model.UserView, error)
	GetViewByNoHp(conn *gorm.DB, noHp string) (model.UserView, error)
	Create(conn *gorm.DB, data model.User) error
	Update(conn *gorm.DB, data model.User) error
	Delete(conn *gorm.DB, data model.User) error
	Page(conn *gorm.DB, req *request.PageUser) ([]model.UserView, int64, error)
}

type repository struct {
}

func (r repository) Page(conn *gorm.DB, req *request.PageUser) ([]model.UserView, int64, error) {
	var err error
	var data []model.UserView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(email) LIKE LOWER(?)", "%"+req.Email+"%").
		Where("LOWER(username) LIKE LOWER(?)", "%"+req.Username+"%").
		Where("LOWER(no_hp) LIKE LOWER(?)", "%"+utils.FormatPhoneTo62(req.NoHp)+"%")

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "fullname", "asc"))
	}
	err = query.Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (r repository) GetById(conn *gorm.DB, id string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetByEmail(conn *gorm.DB, email string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("email = ? ", email).First(&data).Error
	return data, err
}

func (r repository) GetByUsername(conn *gorm.DB, username string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("username = ? ", username).First(&data).Error
	return data, err
}

func (r repository) GetByNoHp(conn *gorm.DB, noHp string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("no_hp = ? ", utils.FormatPhoneTo62(noHp)).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewByEmail(conn *gorm.DB, email string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("email = ? ", email).First(&data).Error
	return data, err
}

func (r repository) GetViewByUsername(conn *gorm.DB, username string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("username = ? ", username).First(&data).Error
	return data, err
}

func (r repository) GetViewByNoHp(conn *gorm.DB, noHp string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("no_hp = ? ", utils.FormatPhoneTo62(noHp)).First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.User) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.User) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.User) error {
	return conn.Delete(&data).Error
}

func NewRepository() Repository {
	return repository{}
}
