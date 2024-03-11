package model

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	ID        string    `gorm:"primaryKey"`
	ClientIP  string    `gorm:"not null"`
	Method    string    `gorm:"not null"`
	Path      string    `gorm:"not null"`
	Code      int       `gorm:"not null"`
	Loginuser string    `gorm:"not null"`
	Request   string    `gorm:"not null"`
	Response  string    `gorm:"not null"`
	CreateDt  time.Time `gorm:"not null"`
}

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
	DeleteDt    gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Ktp struct {
	ID               string         `gorm:"primaryKey" json:"id"`
	Nik              string         `gorm:"not null" json:"nik"`
	Nama             string         `gorm:"not null" json:"nama"`
	TempatLahir      string         `gorm:"not null" json:"tempatLahir"`
	TanggalLahir     time.Time      `gorm:"not null" json:"tanggalLahir"`
	JenisKelamin     string         `gorm:"not null" json:"jenisKelamin"`
	ProvinceID       string         `gorm:"not null" json:"provinceId"`
	RegencyID        string         `gorm:"not null" json:"regencyId"`
	DistrictID       string         `gorm:"not null" json:"districtId"`
	VillageID        string         `gorm:"not null" json:"villageId"`
	Alamat           string         `gorm:"not null" json:"alamat"`
	Rtrw             string         `gorm:"not null" json:"rtrw"`
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
