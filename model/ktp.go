package model

import (
	"github.com/jihanlugas/pandora/utils"
	"gorm.io/gorm"
	"time"
)

func (m *Ktp) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.ID == "" {
		m.ID = utils.GetUniqueID()
	}

	m.CreateDt = now
	m.UpdateDt = now
	return
}

func (m *Ktp) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.UpdateDt = now
	return
}

func (m *Ktp) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Save(m).Error
}
