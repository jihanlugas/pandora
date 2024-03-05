package scheduler

import (
	"fmt"
	"github.com/jihanlugas/pandora/app/log"
	"github.com/jihanlugas/pandora/config"
	"github.com/jihanlugas/pandora/db"
	"time"
)

func DeleteLog() {
	fmt.Println("schedulle delete log: run")
	var err error
	logRepo := log.NewRepository()

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	now := time.Now()
	timeTodelete := now.Add(time.Duration(-config.ToDeleteLogHour) * time.Hour)
	data, err := logRepo.GetDataBefore(tx, timeTodelete)
	if err != nil {
		fmt.Printf("Error when get data log : %s\n", err.Error())
		return
	}

	if len(data) > 0 {
		fmt.Printf("schedulle delete log: %d to delete\n", len(data))
		err = logRepo.Deletes(tx, data)
		if err != nil {
			fmt.Printf("Error when delete data log : %s\n", err.Error())
			return
		}
	}

	err = tx.Commit().Error
	if err != nil {
		fmt.Printf("Error when commit transaction delete log: %s\n", err.Error())
		return
	}

	fmt.Println("schedulle delete log: end")
}
