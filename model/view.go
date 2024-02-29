package model

import (
	"gorm.io/gorm"
	"time"
)

type UserView struct {
	ID          string         `json:"id"`
	Role        string         `json:"role"`
	Email       string         `json:"email"`
	Username    string         `json:"username"`
	NoHp        string         `json:"noHp"`
	Fullname    string         `json:"fullname"`
	Passwd      string         `json:"-"`
	PassVersion int            `json:"passVersion"`
	IsActive    bool           `json:"isActive"`
	LastLoginDt *time.Time     `json:"lastLoginDt"`
	PhotoID     string         `json:"photoId"`
	PhotoUrl    string         `json:"photoUrl"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`
}

func (UserView) TableName() string {
	return VIEW_USER
}

type KtpView struct {
	ID               string         `json:"id"`
	Nik              string         `json:"nik"`
	Nama             string         `json:"nama"`
	TempatLahir      string         `json:"tempatLahir"`
	TanggalLahir     time.Time      `json:"tanggalLahir"`
	Alamat           string         `json:"alamat"`
	Rtrw             string         `json:"rtrw"`
	KelurahanDesa    string         `json:"kelurahanDesa"`
	Kecamatan        string         `json:"kecamatan"`
	KabupatenKota    string         `json:"kabupatenKota"`
	Provinsi         string         `json:"provinsi"`
	Pekerjaan        string         `json:"pekerjaan"`
	StatusPerkawinan string         `json:"statusPerkawinan"`
	Kewarganegaraan  string         `json:"kewarganegaraan"`
	BerlakuHingga    *time.Time     `json:"berlakuHingga"`
	PhotoId          string         `json:"photoId"`
	CreateBy         string         `json:"createBy"`
	CreateDt         time.Time      `json:"createDt"`
	UpdateBy         string         `json:"updateBy"`
	UpdateDt         time.Time      `json:"updateDt"`
	DeleteDt         gorm.DeletedAt `json:"deleteDt"`
	CreateName       string         `json:"createName"`
	UpdateName       string         `json:"updateName"`
}

func (KtpView) TableName() string {
	return VIEW_KTP
}
