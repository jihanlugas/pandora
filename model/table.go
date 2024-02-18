package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Role        string         `gorm:"not null" json:"role"`
	Email       string         `gorm:"not null" json:"email"`
	Username    string         `gorm:"not null" json:"username"`
	NoHp        string         `gorm:"not null" json:"noHp"`
	Fullname    string         `gorm:"not null" json:"fullname"`
	Passwd      string         `gorm:"not null" json:"-"`
	PassVersion int            `gorm:"not null" json:"passVersion"`
	IsActive    bool           `gorm:"not null" json:"isActive"`
	PhotoID     string         `gorm:"not null" json:"photoId"`
	LastLoginDt *time.Time     `gorm:"null" json:"lastLoginDt"`
	CreateBy    string         `gorm:"not null" json:"createBy"`
	CreateDt    time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy    string         `gorm:"not null" json:"updateBy"`
	UpdateDt    time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy    string         `gorm:"not null" json:"deleteBy"`
	DeleteDt    gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Ktp struct {
	ID               string         `gorm:"primaryKey" json:"id"`
	Nik              string         `gorm:"not null" json:"nik"`
	Nama             string         `gorm:"not null" json:"nama"`
	TempatLahir      string         `gorm:"not null" json:"tempatLahir"`
	TanggalLahir     time.Time      `gorm:"not null" json:"tanggalLahir"`
	Alamat           string         `gorm:"not null" json:"alamat"`
	Rtrw             string         `gorm:"not null" json:"rtrw"`
	KelurahanDesa    string         `gorm:"not null" json:"kelurahanDesa"`
	Kecamatan        string         `gorm:"not null" json:"kecamatan"`
	KabupatenKota    string         `gorm:"not null" json:"kabupatenKota"`
	Provinsi         string         `gorm:"not null" json:"provinsi"`
	Pekerjaan        string         `gorm:"not null" json:"pekerjaan"`
	StatusPerkawinan string         `gorm:"not null" json:"statusPerkawinan"`
	Kewarganegaraan  string         `gorm:"not null" json:"kewarganegaraan"`
	BerlakuHingga    *time.Time     `gorm:"null" json:"berlakuHingga"`
	PhotoId          string         `gorm:"not null" json:"photoId"`
	CreateBy         string         `gorm:"not null" json:"createBy"`
	CreateDt         time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy         string         `gorm:"not null" json:"updateBy"`
	UpdateDt         time.Time      `gorm:"not null" json:"updateDt"`
	DeleteDt         gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}
