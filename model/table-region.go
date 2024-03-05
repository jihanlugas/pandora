package model

type Province struct {
	ID           string `gorm:"primaryKey"`
	ProvinceName string `gorm:"not null"`
}

func (u *Province) TableName() string {
	return "region.provinces"
}

type Regency struct {
	ID          string `gorm:"primaryKey"`
	ProvinceID  string `gorm:"not null"`
	RegencyName string `gorm:"not null"`
}

func (u *Regency) TableName() string {
	return "region.regencies"
}

type District struct {
	ID           string `gorm:"primaryKey"`
	ProvinceID   string `gorm:"not null"`
	RegencyID    string `gorm:"not null"`
	DistrictName string `gorm:"not null"`
}

func (u *District) TableName() string {
	return "region.districts"
}

type Village struct {
	ID          string `gorm:"primaryKey"`
	ProvinceID  string `gorm:"not null"`
	RegencyID   string `gorm:"not null"`
	DistrictID  string `gorm:"not null"`
	VillageName string `gorm:"not null"`
}

func (u *Village) TableName() string {
	return "region.villages"
}
