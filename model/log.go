package model

import (
	"github.com/jihanlugas/pandora/utils"
	"gorm.io/gorm"
	"time"
)

func (m *Log) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.ID == "" {
		m.ID = utils.GetUniqueID()
	}

	m.CreateDt = now
	return
}
