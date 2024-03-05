package log

import (
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
)

type Usecase interface {
	Create(req model.Log) error
	SchedulleDelete() error
}

type usecaseLog struct {
	repo Repository
}

func (u usecaseLog) Create(req model.Log) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	err = u.repo.Create(tx, req)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseLog) SchedulleDelete() error {
	return nil
}

func NewUsecase(repo Repository) Usecase {
	return usecaseLog{
		repo: repo,
	}
}
